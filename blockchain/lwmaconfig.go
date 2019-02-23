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

func createPow(value string) *big.Int {
	powLimit, _ := new(big.Int).SetString(value, 0)
	return powLimit
}

// GetMainnetLWMA returns Mainnet config for LWMA
func GetMainnetLWMA() LwmaConfig {
	return LwmaConfig{
		enableHeight:        536200,
		testnet:             false,
		regtest:             false,
		powTargetSpacing:    600,
		averagingWindow:     45,
		adjustWeight:        13772,
		minDenominator:      10,
		solveTimeLimitation: true,
		powLimit:            createPow("14134776517815698497336078495404605830980533548759267698564454644503805952"),
	}
}

// GetTestnetLWMA returns Testnet config for LWMA
func GetTestnetLWMA() LwmaConfig {
	return LwmaConfig{
		enableHeight:        14300,
		testnet:             true,
		regtest:             false,
		powTargetSpacing:    600,
		averagingWindow:     45,
		adjustWeight:        13772,
		minDenominator:      10,
		solveTimeLimitation: false,
		powLimit:            createPow("14134776518227074636666380005943348126619871175004951664972849610340958207"),
	}
}

// GetRegtestLWMA returns Regtest config for LWMA
func GetRegtestLWMA() LwmaConfig {
	return LwmaConfig{
		enableHeight:        0,
		testnet:             false,
		regtest:             true,
		powTargetSpacing:    600,
		averagingWindow:     45,
		adjustWeight:        13772,
		minDenominator:      10,
		solveTimeLimitation: false,
		powLimit:            createPow("57896044618658097711785492504343953926634992332820282019728792003956564819967"),
	}
}
