package chaincfg

import (
	"math/big"
)

// LwmaConfig for Linear Weighted Moving Average
// of Bitcoing Gold
type LwmaConfig struct {
	EnableHeight        int32    // Height at witch the lwma is enabled
	Testnet             bool     // Indicate if testnet
	Regtest             bool     // Indicate if regtest
	PowTargetSpacing    int32    // Spacing of pow target
	AveragingWindow     int32    // Average window
	AdjustWeight        int32    // Adjust weight
	MinDenominator      int32    // Min denominator
	SolveTimeLimitation bool     // Solve time limitation
	PowLimit            *big.Int // Pow limit
}
