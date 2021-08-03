package params

import "github.com/ethereum/go-ethereum/common"

var (
	BurnAddress       = common.HexToAddress("0x62") // 62 is the hex byte value of b character
	ValidatorContract = common.HexToAddress("0x66") // 66 is rchain network id
	NetworkFeeAddress = common.HexToAddress("0x72") // 72 is the hex byte value of r character
)
