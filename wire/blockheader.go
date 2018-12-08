// Copyright (c) 2013-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wire

import (
	"bytes"
	"encoding/binary"
	"io"
	"time"

	"github.com/btgsuite/btgd/chaincfg/chainhash"
)

// MaxSolutionSize is the max known Equihash solution size (1344 is for Equihash-200,9)
const MaxSolutionSize = 1344

// MaxBlockHeaderPayload is the maximum number of bytes a block header can be.
const MaxBlockHeaderPayload = 144 + MaxSolutionSize

// BlockHeader defines information about a block and is used in the bitcoin
// block (MsgBlock) and headers (MsgHeaders) messages.
type BlockHeader struct {
	// Version of the block.  This is not the same as the protocol version.
	Version int32

	// Hash of the previous block header in the block chain.
	PrevBlock chainhash.Hash

	// Merkle tree reference to hash of all transactions for the block.
	MerkleRoot chainhash.Hash

	// The block height
	Height uint32

	// Reversed bytes (always zero)
	Reserved [7]uint32

	// Time the block was created.  This is, unfortunately, encoded as a
	// uint32 on the wire and therefore is limited to 2106.
	Timestamp time.Time

	// Difficulty target for the block.
	Bits uint32

	// Nonce used to generate the block.
	Nonce [32]byte

	// Equihash solution
	Solution []byte
}

// BlockHeaderBytesFromBuffer returns a slice of the input buffer with the data after the block
// header truncated.
func BlockHeaderBytesFromBuffer(buffer []byte) []byte {
	r := bytes.NewReader(buffer)
	var h BlockHeader
	h.Deserialize(r)
	return buffer[:h.BlockHeaderLen()]
}

// BlockHeaderLen returns the number of bytes for the block header.
func (h *BlockHeader) BlockHeaderLen() int {
	nSol := len(h.Solution)
	return 140 + VarIntSerializeSize(uint64(nSol)) + nSol
}

// BlockHash computes the block identifier hash for the given block header.
func (h *BlockHeader) BlockHash() chainhash.Hash {
	return h.blockHashInternal(len(h.Solution) == 0)
}

func (h *BlockHeader) blockHashInternal(legacy bool) chainhash.Hash {
	// Encode the header and double sha256 everything prior to the number of
	// transactions.  Ignore the error returns since there is no way the
	// encode could fail except being out of memory which would cause a
	// run-time panic.
	buf := bytes.NewBuffer(make([]byte, 0, MaxBlockHeaderPayload))
	if legacy {
		_ = writeBlockHeaderLegacy(buf, 0, h)
	} else {
		_ = writeBlockHeader(buf, 0, h)
	}

	return chainhash.DoubleHashH(buf.Bytes())
}

// BtcDecode decodes r using the bitcoin protocol encoding into the receiver.
// This is part of the Message interface implementation.
// See Deserialize for decoding block headers stored to disk, such as in a
// database, as opposed to decoding block headers from the wire.
func (h *BlockHeader) BtcDecode(r io.Reader, pver uint32, enc MessageEncoding) error {
	return readBlockHeader(r, pver, h)
}

// BtcEncode encodes the receiver to w using the bitcoin protocol encoding.
// This is part of the Message interface implementation.
// See Serialize for encoding block headers to be stored to disk, such as in a
// database, as opposed to encoding block headers for the wire.
func (h *BlockHeader) BtcEncode(w io.Writer, pver uint32, enc MessageEncoding) error {
	return writeBlockHeader(w, pver, h)
}

// Deserialize decodes a block header from r into the receiver using a format
// that is suitable for long-term storage such as a database while respecting
// the Version field.
func (h *BlockHeader) Deserialize(r io.Reader) error {
	// At the current time, there is no difference between the wire encoding
	// at protocol version 0 and the stable long-term storage format.  As
	// a result, make use of readBlockHeader.
	return readBlockHeader(r, 0, h)
}

// Serialize encodes a block header from r into the receiver using a format
// that is suitable for long-term storage such as a database while respecting
// the Version field.
func (h *BlockHeader) Serialize(w io.Writer) error {
	// At the current time, there is no difference between the wire encoding
	// at protocol version 0 and the stable long-term storage format.  As
	// a result, make use of writeBlockHeader.
	return writeBlockHeader(w, 0, h)
}

// NewBlockHeader returns a new BlockHeader using the provided version, previous
// block hash, merkle root hash, difficulty bits, and nonce used to generate the
// block with defaults for the remaining fields.
func NewBlockHeader(version int32, prevHash, merkleRootHash *chainhash.Hash,
	height uint32, bits uint32, nonce *[32]byte, solution []byte) *BlockHeader {

	// Limit the timestamp to one second precision since the protocol
	// doesn't support better.
	solutionCopy := make([]byte, len(solution))
	copy(solutionCopy, solution)
	return &BlockHeader{
		Version:    version,
		PrevBlock:  *prevHash,
		MerkleRoot: *merkleRootHash,
		Timestamp:  time.Unix(time.Now().Unix(), 0),
		Height:     height,
		Reserved:   [7]uint32{},
		Bits:       bits,
		Nonce:      *nonce,
		Solution:   solutionCopy,
	}
}

// NewLegacyBlockHeader returns a legacy Bitcoin block header.
func NewLegacyBlockHeader(version int32, prevHash, merkleRootHash *chainhash.Hash,
	bits uint32, nonce uint32) *BlockHeader {
	nounce256 := Uint256FromUint32(nonce)
	return NewBlockHeader(version, prevHash, merkleRootHash, 0, bits, &nounce256, []byte{})
}

// ReadBlockHeaderLegacy reads a legacy bitcoin block from r.
func ReadBlockHeaderLegacy(r io.Reader, pver uint32, bh *BlockHeader) error {
	var nonce uint32
	err := readElements(r, &bh.Version, &bh.PrevBlock, &bh.MerkleRoot,
		(*uint32Time)(&bh.Timestamp), &bh.Bits, &nonce)
	if err != nil {
		return err
	}
	bh.Nonce = Uint256FromUint32(nonce)
	bh.Solution = []byte{}
	return nil
}

// readBlockHeader reads a Bitcoin Gold bitcoin block header from r.  See Deserialize for
// decoding block headers stored to disk, such as in a database, as opposed to
// decoding from the wire.
func readBlockHeader(r io.Reader, pver uint32, bh *BlockHeader) error {
	if err := readElements(r, &bh.Version, &bh.PrevBlock, &bh.MerkleRoot, &bh.Height); err != nil {
		return err
	}
	for i := range bh.Reserved {
		if err := readElement(r, &bh.Reserved[i]); err != nil {
			return err
		}
	}
	if err := readElements(r, (*uint32Time)(&bh.Timestamp), &bh.Bits, &bh.Nonce); err != nil {
		return err
	}
	solution, err := ReadVarBytes(r, pver, MaxSolutionSize, "Solution")
	if err != nil {
		return err
	}
	bh.Solution = solution
	return nil
}

// writeBlockHeader and writeBlockHeaderLegacy writes a bitcoin block
// header to w.  See Serialize for encoding block headers to be stored
// to disk, such as in a database, as opposed to encoding for the wire.
func writeBlockHeaderLegacy(w io.Writer, pver uint32, bh *BlockHeader) error {
	sec := uint32(bh.Timestamp.Unix())
	nonceUint32 := binary.LittleEndian.Uint32(bh.Nonce[0:4])
	return writeElements(w, bh.Version, &bh.PrevBlock, &bh.MerkleRoot,
		sec, bh.Bits, nonceUint32)
}
func writeBlockHeader(w io.Writer, pver uint32, bh *BlockHeader) error {
	sec := uint32(bh.Timestamp.Unix())
	if err := writeElements(w, bh.Version, &bh.PrevBlock, &bh.MerkleRoot, bh.Height); err != nil {
		return err
	}
	for _, v := range bh.Reserved {
		if err := writeElement(w, v); err != nil {
			return err
		}
	}
	if err := writeElements(w, sec, bh.Bits, bh.Nonce); err != nil {
		return err
	}
	if err := WriteVarBytes(w, pver, bh.Solution); err != nil {
		return err
	}
	return nil
}
