package blockchain

import (
	"math/big"
)

// LwmaConfig for Linear Weighted Moving Average
// of Bitcoing Gold
type LwmaConfig struct {
	enableHeight        int32    // Height at witch the lwma is enabled
	testnet             bool     // Indicate if testnet
	regtest             bool     // Indicate if regtest
	powTargetSpacing    int32    // Spacing of pow target
	averagingWindow     int32    // Average window
	adjustWeight        int32    // Adjust weight
	minDenominator      int32    // Min denominator
	solveTimeLimitation bool     // Solve time limitation
	powLimit            *big.Int // Pow limit
}
