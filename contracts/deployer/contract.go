package deployer

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

type DeployCallbackFn = func(*backends.SimulatedBackend, *bind.TransactOpts) (common.Address, error)

// DeployContract deploy a smart contract to simulated chain to the target address in the StateDB
func DeployContract(deployCallback DeployCallbackFn) (code []byte, storage map[common.Hash]common.Hash, err error) {
	// Generate a new random account and a funded simulator
	prvKey, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(prvKey)
	auth.GasLimit = 80000000
	sim := backends.NewSimulatedBackend(core.GenesisAlloc{auth.From: {Balance: new(big.Int).Lsh(big.NewInt(1), 256-7)}}, auth.GasLimit)
	address, err := deployCallback(sim, auth)
	if err != nil {
		log.Error("Unable to deploy contract", "error", err)
		return nil, nil, err
	}
	sim.Commit()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	code, _ = sim.CodeAt(ctx, address, nil)
	storage = make(map[common.Hash]common.Hash)
	sim.ForEachStorageAt(ctx, address, nil, func(key, val common.Hash) bool {
		log.Debug("decoding contract storage", "address", address, "key", key, "val", val.Hex(), "storage value", storage[key].Hex())
		storage[key] = val
		return true
	})
	return code, storage, err
}
