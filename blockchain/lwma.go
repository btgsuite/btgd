// https://github.com/BTCGPU/BTCGPU/blob/c919e0774806601f8b192378d078f63f7804b721/src/pow.cpp#L74

package blockchain

import (
	"fmt"
	"math/big"
	"time"

	"github.com/btgsuite/btgd/chaincfg"
	"github.com/btgsuite/btgd/wire"
)

// CalcNextBits calculates the expected bits value from values
// in last blocks based on LWMA.
func CalcNextBits(height uint32, curTimestamp time.Time, previousBlocks []wire.BlockHeader, lwmaConfig *chaincfg.LwmaConfig) (uint32, error) {
	if int32(len(previousBlocks)) <= lwmaConfig.AveragingWindow {
		return 0, AssertError(fmt.Sprintf("LWMA need the last %d blocks to determine the next target", lwmaConfig.AveragingWindow+1))
	}

	prevBlocks := make(map[uint32]wire.BlockHeader)
	for _, b := range previousBlocks {
		prevBlocks[b.Height] = b
	}

	for i := height - uint32(lwmaConfig.AveragingWindow) - 1; i < height; i++ {
		if _, ok := prevBlocks[i]; !ok {
			return 0, AssertError(fmt.Sprintf("Block with height %d is missing, cannot calculate next target", i))
		}
	}

	// loss of precision when converting target to bits, comparing target to target (from bits) will result in different uint256
	nextTarget := getLwmaTarget(height, curTimestamp, prevBlocks, lwmaConfig)
	bits := targetToBits(*nextTarget)
	return bits, nil
}

func getLwmaTarget(height uint32, curTimestamp time.Time, prevBlocks map[uint32]wire.BlockHeader, lwmaConfig *chaincfg.LwmaConfig) *big.Int {
	weight := lwmaConfig.AdjustWeight
	prev := prevBlocks[height-1]

	// Special testnet handling
	if lwmaConfig.Regtest {
		return bitsToTarget(prev.Bits)
	}

	if lwmaConfig.Testnet && curTimestamp.Unix() > prev.Timestamp.Unix()+int64(lwmaConfig.PowTargetSpacing*2) {
		return lwmaConfig.PowLimit
	}

	totalBig := new(big.Int)
	t := int64(0)
	j := int64(0)
	ts := int64(6 * lwmaConfig.PowTargetSpacing)
	dividerBig := new(big.Int).SetInt64(int64(weight * lwmaConfig.AveragingWindow * lwmaConfig.AveragingWindow))

	// Loop through N most recent blocks.  "< height", not "<="
	for i := height - uint32(lwmaConfig.AveragingWindow); i < height; i++ {
		cur := prevBlocks[i]
		prev := prevBlocks[i-1]

		solvetime := cur.Timestamp.Unix() - prev.Timestamp.Unix()
		if lwmaConfig.SolveTimeLimitation && solvetime > ts {
			solvetime = ts
		}

		j++
		t += solvetime * j
		targetBig := bitsToTarget(cur.Bits)
		totalBig.Add(totalBig, targetBig.Div(targetBig, dividerBig))
	}

	// Keep t reasonable in case strange solvetimes occurred.
	if t < int64(lwmaConfig.AveragingWindow*weight/lwmaConfig.MinDenominator) {
		t = int64(lwmaConfig.AveragingWindow * weight / lwmaConfig.MinDenominator)
	}

	totalBig.Mul(totalBig, new(big.Int).SetInt64(t))
	if totalBig.Cmp(lwmaConfig.PowLimit) >= 0 {
		totalBig = lwmaConfig.PowLimit
	}

	return totalBig
}

func bitsToTarget(bits uint32) *big.Int {
	bitsBig := new(big.Int).SetInt64(int64(bits))
	bitsBig.Rsh(bitsBig, 24)
	size := uint(bitsBig.Uint64())
	word := bits & 0x007fffff

	wordBig := new(big.Int).SetInt64(int64(word))
	if size <= 3 {
		return wordBig.Rsh(wordBig, 8*(3-size))
	}

	return wordBig.Lsh(wordBig, 8*(size-3))
}

func targetToBits(target big.Int) uint32 {
	nsize := int64((target.BitLen() + 7) / 8)
	cBig := new(big.Int).SetUint64(0)

	if nsize <= 3 {
		cBig = target.Lsh(&target, uint(8*(3-nsize)))
	} else {
		cBig = target.Rsh(&target, uint(8*(nsize-3)))
	}

	c := cBig.Int64()
	if c&0x00800000 != 0 {
		c >>= 8
		nsize++
	}

	c |= nsize << 24
	return uint32(c)
}
