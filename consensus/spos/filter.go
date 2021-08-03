package spos

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/rpc"
)

// LightBackend light backend
type LightBackend interface {
	HeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Header, error)
	HeaderByHash(ctx context.Context, blockHash common.Hash) (*types.Header, error)
	GetReceipts(ctx context.Context, blockHash common.Hash) (types.Receipts, error)
	GetLogs(ctx context.Context, blockHash common.Hash) ([][]*types.Log, error)
}

// filterBackend light backend implementation
type filterBackend struct {
	chain consensus.ChainReader
	db    ethdb.Reader
}

// BlockFilter single block log filter
type BlockFilter struct {
	backend LightBackend

	addresses []common.Address
	topics    [][]common.Hash

	block common.Hash // Block hash for filtering a single block
}

// NewBlockFilter creates a new filter which directly inspects the contents of
// a block to figure out whether it is interesting or not.
func NewBlockFilter(backend LightBackend, block common.Hash, addresses []common.Address, topics [][]common.Hash) *BlockFilter {
	// Create a generic filter and convert it into a block filter
	filter := newFilter(backend, addresses, topics)
	filter.block = block
	return filter
}

// newFilter creates a generic filter that can filter based on a block hash,
// The search criteria needs to be explicitly set.
func newFilter(backend LightBackend, addresses []common.Address, topics [][]common.Hash) *BlockFilter {
	return &BlockFilter{
		backend:   backend,
		addresses: addresses,
		topics:    topics,
	}
}

func (fb *filterBackend) HeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Header, error) {
	return fb.chain.GetHeaderByNumber(uint64(blockNr.Int64())), nil
}

func (fb *filterBackend) HeaderByHash(ctx context.Context, blockHash common.Hash) (*types.Header, error) {
	return fb.chain.GetHeaderByHash(blockHash), nil
}

func (fb *filterBackend) GetReceipts(ctx context.Context, blockHash common.Hash) (types.Receipts, error) {
	header := fb.chain.GetHeaderByHash(blockHash)
	if header == nil {
		return nil, nil
	}
	receipts := rawdb.ReadReceipts(fb.db, blockHash, header.Number.Uint64(), fb.chain.Config())
	return receipts, nil
}

func (fb *filterBackend) GetLogs(ctx context.Context, blockHash common.Hash) ([][]*types.Log, error) {
	receipts, _ := fb.GetReceipts(ctx, blockHash)
	logs := make([][]*types.Log, len(receipts))
	for i, receipt := range receipts {
		logs[i] = receipt.Logs
	}
	return logs, nil
}

// BlockLogs returns the logs matching the filter criteria within a single block.
func (f *BlockFilter) BlockLogs(ctx context.Context, header *types.Header) (logs []*types.Log, err error) {
	if bloomFilter(header.Bloom, f.addresses, f.topics) {
		found, err := f.checkMatches(ctx, header)
		if err != nil {
			return logs, err
		}
		logs = append(logs, found...)
	}
	return logs, nil
}

// checkMatches checks if the receipts belonging to the given header contain any log events that
// match the filter criteria. This function is called when the bloom filter signals a potential match.
func (f *BlockFilter) checkMatches(ctx context.Context, header *types.Header) (logs []*types.Log, err error) {
	// Get the logs of the block
	logsList, err := f.backend.GetLogs(ctx, header.Hash())
	if err != nil {
		return nil, err
	}
	var unfiltered []*types.Log
	for _, logs := range logsList {
		unfiltered = append(unfiltered, logs...)
	}
	logs = filterLogs(unfiltered, nil, nil, f.addresses, f.topics)
	if len(logs) > 0 {
		// We have matching logs, check if we need to resolve full logs via the light client
		if logs[0].TxHash == (common.Hash{}) {
			receipts, err := f.backend.GetReceipts(ctx, header.Hash())
			if err != nil {
				return nil, err
			}
			unfiltered = unfiltered[:0]
			for _, receipt := range receipts {
				unfiltered = append(unfiltered, receipt.Logs...)
			}
			logs = filterLogs(unfiltered, nil, nil, f.addresses, f.topics)
		}
		return logs, nil
	}
	return nil, nil
}

func includes(addresses []common.Address, a common.Address) bool {
	for _, addr := range addresses {
		if addr == a {
			return true
		}
	}

	return false
}

// filterLogs creates a slice of logs matching the given criteria.
func filterLogs(logs []*types.Log, fromBlock, toBlock *big.Int, addresses []common.Address, topics [][]common.Hash) []*types.Log {
	var ret []*types.Log
Logs:
	for _, log := range logs {
		if fromBlock != nil && fromBlock.Int64() >= 0 && fromBlock.Uint64() > log.BlockNumber {
			continue
		}
		if toBlock != nil && toBlock.Int64() >= 0 && toBlock.Uint64() < log.BlockNumber {
			continue
		}

		if len(addresses) > 0 && !includes(addresses, log.Address) {
			continue
		}
		// If the to filtered topics is greater than the amount of topics in logs, skip.
		if len(topics) > len(log.Topics) {
			continue Logs
		}
		for i, sub := range topics {
			match := len(sub) == 0 // empty rule set == wildcard
			for _, topic := range sub {
				if log.Topics[i] == topic {
					match = true
					break
				}
			}
			if !match {
				continue Logs
			}
		}
		ret = append(ret, log)
	}
	return ret
}

func bloomFilter(bloom types.Bloom, addresses []common.Address, topics [][]common.Hash) bool {
	if len(addresses) > 0 {
		var included bool
		for _, addr := range addresses {
			if types.BloomLookup(bloom, addr) {
				included = true
				break
			}
		}
		if !included {
			return false
		}
	}

	for _, sub := range topics {
		included := len(sub) == 0 // empty rule set == wildcard
		for _, topic := range sub {
			if types.BloomLookup(bloom, topic) {
				included = true
				break
			}
		}
		if !included {
			return false
		}
	}
	return true
}
