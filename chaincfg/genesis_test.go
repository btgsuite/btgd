// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

// TestGenesisBlock tests the genesis block of the main network for validity by
// checking the encoded bytes and hashes.
func TestGenesisBlock(t *testing.T) {
	// Encode the genesis block to raw bytes.
	var buf bytes.Buffer
	err := MainNetParams.GenesisBlock.Serialize(&buf)
	if err != nil {
		t.Fatalf("TestGenesisBlock: %v", err)
	}

	// Ensure the encoded block matches the expected bytes.
	if !bytes.Equal(buf.Bytes(), genesisBlockBytes) {
		t.Fatalf("TestGenesisBlock: Genesis block does not appear valid - "+
			"got %v, want %v", spew.Sdump(buf.Bytes()),
			spew.Sdump(genesisBlockBytes))
	}

	// Check hash of the block against expected hash.
	hash := MainNetParams.GenesisBlock.BlockHash()
	if !MainNetParams.GenesisHash.IsEqual(&hash) {
		t.Fatalf("TestGenesisBlock: Genesis block hash does not "+
			"appear valid - got %v, want %v", spew.Sdump(hash),
			spew.Sdump(MainNetParams.GenesisHash))
	}
}

// TestRegTestGenesisBlock tests the genesis block of the regression test
// network for validity by checking the encoded bytes and hashes.
func TestRegTestGenesisBlock(t *testing.T) {
	// Encode the genesis block to raw bytes.
	var buf bytes.Buffer
	err := RegressionNetParams.GenesisBlock.Serialize(&buf)
	if err != nil {
		t.Fatalf("TestRegTestGenesisBlock: %v", err)
	}

	// Ensure the encoded block matches the expected bytes.
	if !bytes.Equal(buf.Bytes(), regTestGenesisBlockBytes) {
		t.Fatalf("TestRegTestGenesisBlock: Genesis block does not "+
			"appear valid - got %v, want %v",
			spew.Sdump(buf.Bytes()),
			spew.Sdump(regTestGenesisBlockBytes))
	}

	// Check hash of the block against expected hash.
	hash := RegressionNetParams.GenesisBlock.BlockHash()
	if !RegressionNetParams.GenesisHash.IsEqual(&hash) {
		t.Fatalf("TestRegTestGenesisBlock: Genesis block hash does "+
			"not appear valid - got %v, want %v", spew.Sdump(hash),
			spew.Sdump(RegressionNetParams.GenesisHash))
	}
}

// TestTestNet3GenesisBlock tests the genesis block of the test network (version
// 3) for validity by checking the encoded bytes and hashes.
func TestTestNet3GenesisBlock(t *testing.T) {
	// Encode the genesis block to raw bytes.
	var buf bytes.Buffer
	err := TestNet3Params.GenesisBlock.Serialize(&buf)
	if err != nil {
		t.Fatalf("TestTestNet3GenesisBlock: %v", err)
	}

	// Ensure the encoded block matches the expected bytes.
	if !bytes.Equal(buf.Bytes(), testNet3GenesisBlockBytes) {
		t.Fatalf("TestTestNet3GenesisBlock: Genesis block does not "+
			"appear valid - got %v, want %v",
			spew.Sdump(buf.Bytes()),
			spew.Sdump(testNet3GenesisBlockBytes))
	}

	// Check hash of the block against expected hash.
	hash := TestNet3Params.GenesisBlock.BlockHash()
	if !TestNet3Params.GenesisHash.IsEqual(&hash) {
		t.Fatalf("TestTestNet3GenesisBlock: Genesis block hash does "+
			"not appear valid - got %v, want %v", spew.Sdump(hash),
			spew.Sdump(TestNet3Params.GenesisHash))
	}
}

// TestSimNetGenesisBlock tests the genesis block of the simulation test network
// for validity by checking the encoded bytes and hashes.
func TestSimNetGenesisBlock(t *testing.T) {
	// Encode the genesis block to raw bytes.
	var buf bytes.Buffer
	err := SimNetParams.GenesisBlock.Serialize(&buf)
	if err != nil {
		t.Fatalf("TestSimNetGenesisBlock: %v", err)
	}

	// Ensure the encoded block matches the expected bytes.
	if !bytes.Equal(buf.Bytes(), simNetGenesisBlockBytes) {
		t.Fatalf("TestSimNetGenesisBlock: Genesis block does not "+
			"appear valid - got %v, want %v",
			spew.Sdump(buf.Bytes()),
			spew.Sdump(simNetGenesisBlockBytes))
	}

	// Check hash of the block against expected hash.
	hash := SimNetParams.GenesisBlock.BlockHash()
	if !SimNetParams.GenesisHash.IsEqual(&hash) {
		t.Fatalf("TestSimNetGenesisBlock: Genesis block hash does "+
			"not appear valid - got %v, want %v", spew.Sdump(hash),
			spew.Sdump(SimNetParams.GenesisHash))
	}
}

// genesisBlockBytes are the wire encoded bytes for the genesis block of the
// main network as of protocol version 60002.
var genesisBlockBytes = []byte{
	0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x3b, 0xa3, 0xed, 0xfd, /* |....;...| */
	0x7a, 0x7b, 0x12, 0xb2, 0x7a, 0xc7, 0x2c, 0x3e, /* |z{..z.,>| */
	0x67, 0x76, 0x8f, 0x61, 0x7f, 0xc8, 0x1b, 0xc3, /* |gv.a....| */
	0x88, 0x8a, 0x51, 0x32, 0x3a, 0x9f, 0xb8, 0xaa, /* |..Q2:...| */
	0x4b, 0x1e, 0x5e, 0x4a, 0x00, 0x00, 0x00, 0x00, /* |K.^J....| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x29, 0xab, 0x5f, 0x49, /* |....)._I| */
	0xff, 0xff, 0x00, 0x1d, 0x1d, 0xac, 0x2b, 0x7c, /* |......+|| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x00, /* |........| */
	0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0x4d, /* |.......M| */
	0x04, 0xff, 0xff, 0x00, 0x1d, 0x01, 0x04, 0x45, /* |.......E| */
	0x54, 0x68, 0x65, 0x20, 0x54, 0x69, 0x6d, 0x65, /* |The Time| */
	0x73, 0x20, 0x30, 0x33, 0x2f, 0x4a, 0x61, 0x6e, /* |s 03/Jan| */
	0x2f, 0x32, 0x30, 0x30, 0x39, 0x20, 0x43, 0x68, /* |/2009 Ch| */
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x6f, 0x72, /* |ancellor| */
	0x20, 0x6f, 0x6e, 0x20, 0x62, 0x72, 0x69, 0x6e, /* | on brin| */
	0x6b, 0x20, 0x6f, 0x66, 0x20, 0x73, 0x65, 0x63, /* |k of sec| */
	0x6f, 0x6e, 0x64, 0x20, 0x62, 0x61, 0x69, 0x6c, /* |ond bail| */
	0x6f, 0x75, 0x74, 0x20, 0x66, 0x6f, 0x72, 0x20, /* |out for | */
	0x62, 0x61, 0x6e, 0x6b, 0x73, 0xff, 0xff, 0xff, /* |banks...| */
	0xff, 0x01, 0x00, 0xf2, 0x05, 0x2a, 0x01, 0x00, /* |.....*..| */
	0x00, 0x00, 0x43, 0x41, 0x04, 0x67, 0x8a, 0xfd, /* |..CA.g..| */
	0xb0, 0xfe, 0x55, 0x48, 0x27, 0x19, 0x67, 0xf1, /* |..UH'.g.| */
	0xa6, 0x71, 0x30, 0xb7, 0x10, 0x5c, 0xd6, 0xa8, /* |.q0..\..| */
	0x28, 0xe0, 0x39, 0x09, 0xa6, 0x79, 0x62, 0xe0, /* |(.9..yb.| */
	0xea, 0x1f, 0x61, 0xde, 0xb6, 0x49, 0xf6, 0xbc, /* |..a..I..| */
	0x3f, 0x4c, 0xef, 0x38, 0xc4, 0xf3, 0x55, 0x04, /* |?L.8..U.| */
	0xe5, 0x1e, 0xc1, 0x12, 0xde, 0x5c, 0x38, 0x4d, /* |.....\8M| */
	0xf7, 0xba, 0x0b, 0x8d, 0x57, 0x8a, 0x4c, 0x70, /* |....W.Lp| */
	0x2b, 0x6b, 0xf1, 0x1d, 0x5f, 0xac, 0x00, 0x00, /* |+k.._...| */
	0x00, 0x00, /* |..| */
}

// regTestGenesisBlockBytes are the wire encoded bytes for the genesis block of
// the regression test network as of protocol version 60002.
var regTestGenesisBlockBytes = []byte{
	0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x3b, 0xa3, 0xed, 0xfd, /* |....;...| */
	0x7a, 0x7b, 0x12, 0xb2, 0x7a, 0xc7, 0x2c, 0x3e, /* |z{..z.,>| */
	0x67, 0x76, 0x8f, 0x61, 0x7f, 0xc8, 0x1b, 0xc3, /* |gv.a....| */
	0x88, 0x8a, 0x51, 0x32, 0x3a, 0x9f, 0xb8, 0xaa, /* |..Q2:...| */
	0x4b, 0x1e, 0x5e, 0x4a, 0x00, 0x00, 0x00, 0x00, /* |K.^J....| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0xda, 0xe5, 0x49, 0x4d, /* |......IM| */
	0xff, 0xff, 0x7f, 0x20, 0x02, 0x00, 0x00, 0x00, /* |... ....| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x00, /* |........| */
	0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0x4d, /* |.......M| */
	0x04, 0xff, 0xff, 0x00, 0x1d, 0x01, 0x04, 0x45, /* |.......E| */
	0x54, 0x68, 0x65, 0x20, 0x54, 0x69, 0x6d, 0x65, /* |The Time| */
	0x73, 0x20, 0x30, 0x33, 0x2f, 0x4a, 0x61, 0x6e, /* |s 03/Jan| */
	0x2f, 0x32, 0x30, 0x30, 0x39, 0x20, 0x43, 0x68, /* |/2009 Ch| */
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x6f, 0x72, /* |ancellor| */
	0x20, 0x6f, 0x6e, 0x20, 0x62, 0x72, 0x69, 0x6e, /* | on brin| */
	0x6b, 0x20, 0x6f, 0x66, 0x20, 0x73, 0x65, 0x63, /* |k of sec| */
	0x6f, 0x6e, 0x64, 0x20, 0x62, 0x61, 0x69, 0x6c, /* |ond bail| */
	0x6f, 0x75, 0x74, 0x20, 0x66, 0x6f, 0x72, 0x20, /* |out for | */
	0x62, 0x61, 0x6e, 0x6b, 0x73, 0xff, 0xff, 0xff, /* |banks...| */
	0xff, 0x01, 0x00, 0xf2, 0x05, 0x2a, 0x01, 0x00, /* |.....*..| */
	0x00, 0x00, 0x43, 0x41, 0x04, 0x67, 0x8a, 0xfd, /* |..CA.g..| */
	0xb0, 0xfe, 0x55, 0x48, 0x27, 0x19, 0x67, 0xf1, /* |..UH'.g.| */
	0xa6, 0x71, 0x30, 0xb7, 0x10, 0x5c, 0xd6, 0xa8, /* |.q0..\..| */
	0x28, 0xe0, 0x39, 0x09, 0xa6, 0x79, 0x62, 0xe0, /* |(.9..yb.| */
	0xea, 0x1f, 0x61, 0xde, 0xb6, 0x49, 0xf6, 0xbc, /* |..a..I..| */
	0x3f, 0x4c, 0xef, 0x38, 0xc4, 0xf3, 0x55, 0x04, /* |?L.8..U.| */
	0xe5, 0x1e, 0xc1, 0x12, 0xde, 0x5c, 0x38, 0x4d, /* |.....\8M| */
	0xf7, 0xba, 0x0b, 0x8d, 0x57, 0x8a, 0x4c, 0x70, /* |....W.Lp| */
	0x2b, 0x6b, 0xf1, 0x1d, 0x5f, 0xac, 0x00, 0x00, /* |+k.._...| */
	0x00, 0x00, /* |..| */
}

// testNet3GenesisBlockBytes are the wire encoded bytes for the genesis block of
// the test network (version 3) as of protocol version 60002.
var testNet3GenesisBlockBytes = []byte{
	0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x3b, 0xa3, 0xed, 0xfd, /* |....;...| */
	0x7a, 0x7b, 0x12, 0xb2, 0x7a, 0xc7, 0x2c, 0x3e, /* |z{..z.,>| */
	0x67, 0x76, 0x8f, 0x61, 0x7f, 0xc8, 0x1b, 0xc3, /* |gv.a....| */
	0x88, 0x8a, 0x51, 0x32, 0x3a, 0x9f, 0xb8, 0xaa, /* |..Q2:...| */
	0x4b, 0x1e, 0x5e, 0x4a, 0x00, 0x00, 0x00, 0x00, /* |K.^J....| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x7c, 0x35, 0x5e, 0x5a, /* |....|5^Z| */
	0xff, 0xff, 0x00, 0x1d, 0x42, 0x51, 0xbd, 0x56, /* |....BQ.V| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x00, /* |........| */
	0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0x4d, /* |.......M| */
	0x04, 0xff, 0xff, 0x00, 0x1d, 0x01, 0x04, 0x45, /* |.......E| */
	0x54, 0x68, 0x65, 0x20, 0x54, 0x69, 0x6d, 0x65, /* |The Time| */
	0x73, 0x20, 0x30, 0x33, 0x2f, 0x4a, 0x61, 0x6e, /* |s 03/Jan| */
	0x2f, 0x32, 0x30, 0x30, 0x39, 0x20, 0x43, 0x68, /* |/2009 Ch| */
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x6f, 0x72, /* |ancellor| */
	0x20, 0x6f, 0x6e, 0x20, 0x62, 0x72, 0x69, 0x6e, /* | on brin| */
	0x6b, 0x20, 0x6f, 0x66, 0x20, 0x73, 0x65, 0x63, /* |k of sec| */
	0x6f, 0x6e, 0x64, 0x20, 0x62, 0x61, 0x69, 0x6c, /* |ond bail| */
	0x6f, 0x75, 0x74, 0x20, 0x66, 0x6f, 0x72, 0x20, /* |out for | */
	0x62, 0x61, 0x6e, 0x6b, 0x73, 0xff, 0xff, 0xff, /* |banks...| */
	0xff, 0x01, 0x00, 0xf2, 0x05, 0x2a, 0x01, 0x00, /* |.....*..| */
	0x00, 0x00, 0x43, 0x41, 0x04, 0x67, 0x8a, 0xfd, /* |..CA.g..| */
	0xb0, 0xfe, 0x55, 0x48, 0x27, 0x19, 0x67, 0xf1, /* |..UH'.g.| */
	0xa6, 0x71, 0x30, 0xb7, 0x10, 0x5c, 0xd6, 0xa8, /* |.q0..\..| */
	0x28, 0xe0, 0x39, 0x09, 0xa6, 0x79, 0x62, 0xe0, /* |(.9..yb.| */
	0xea, 0x1f, 0x61, 0xde, 0xb6, 0x49, 0xf6, 0xbc, /* |..a..I..| */
	0x3f, 0x4c, 0xef, 0x38, 0xc4, 0xf3, 0x55, 0x04, /* |?L.8..U.| */
	0xe5, 0x1e, 0xc1, 0x12, 0xde, 0x5c, 0x38, 0x4d, /* |.....\8M| */
	0xf7, 0xba, 0x0b, 0x8d, 0x57, 0x8a, 0x4c, 0x70, /* |....W.Lp| */
	0x2b, 0x6b, 0xf1, 0x1d, 0x5f, 0xac, 0x00, 0x00, /* |+k.._...| */
	0x00, 0x00, /* |..| */
}

// simNetGenesisBlockBytes are the wire encoded bytes for the genesis block of
// the simulation test network as of protocol version 70002.
var simNetGenesisBlockBytes = []byte{
	0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x3b, 0xa3, 0xed, 0xfd, /* |....;...| */
	0x7a, 0x7b, 0x12, 0xb2, 0x7a, 0xc7, 0x2c, 0x3e, /* |z{..z.,>| */
	0x67, 0x76, 0x8f, 0x61, 0x7f, 0xc8, 0x1b, 0xc3, /* |gv.a....| */
	0x88, 0x8a, 0x51, 0x32, 0x3a, 0x9f, 0xb8, 0xaa, /* |..Q2:...| */
	0x4b, 0x1e, 0x5e, 0x4a, 0x00, 0x00, 0x00, 0x00, /* |K.^J....| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x45, 0x06, 0x86, 0x53, /* |....E..S| */
	0xff, 0xff, 0x7f, 0x20, 0x02, 0x00, 0x00, 0x00, /* |... ....| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x00, /* |........| */
	0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0x4d, /* |.......M| */
	0x04, 0xff, 0xff, 0x00, 0x1d, 0x01, 0x04, 0x45, /* |.......E| */
	0x54, 0x68, 0x65, 0x20, 0x54, 0x69, 0x6d, 0x65, /* |The Time| */
	0x73, 0x20, 0x30, 0x33, 0x2f, 0x4a, 0x61, 0x6e, /* |s 03/Jan| */
	0x2f, 0x32, 0x30, 0x30, 0x39, 0x20, 0x43, 0x68, /* |/2009 Ch| */
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x6f, 0x72, /* |ancellor| */
	0x20, 0x6f, 0x6e, 0x20, 0x62, 0x72, 0x69, 0x6e, /* | on brin| */
	0x6b, 0x20, 0x6f, 0x66, 0x20, 0x73, 0x65, 0x63, /* |k of sec| */
	0x6f, 0x6e, 0x64, 0x20, 0x62, 0x61, 0x69, 0x6c, /* |ond bail| */
	0x6f, 0x75, 0x74, 0x20, 0x66, 0x6f, 0x72, 0x20, /* |out for | */
	0x62, 0x61, 0x6e, 0x6b, 0x73, 0xff, 0xff, 0xff, /* |banks...| */
	0xff, 0x01, 0x00, 0xf2, 0x05, 0x2a, 0x01, 0x00, /* |.....*..| */
	0x00, 0x00, 0x43, 0x41, 0x04, 0x67, 0x8a, 0xfd, /* |..CA.g..| */
	0xb0, 0xfe, 0x55, 0x48, 0x27, 0x19, 0x67, 0xf1, /* |..UH'.g.| */
	0xa6, 0x71, 0x30, 0xb7, 0x10, 0x5c, 0xd6, 0xa8, /* |.q0..\..| */
	0x28, 0xe0, 0x39, 0x09, 0xa6, 0x79, 0x62, 0xe0, /* |(.9..yb.| */
	0xea, 0x1f, 0x61, 0xde, 0xb6, 0x49, 0xf6, 0xbc, /* |..a..I..| */
	0x3f, 0x4c, 0xef, 0x38, 0xc4, 0xf3, 0x55, 0x04, /* |?L.8..U.| */
	0xe5, 0x1e, 0xc1, 0x12, 0xde, 0x5c, 0x38, 0x4d, /* |.....\8M| */
	0xf7, 0xba, 0x0b, 0x8d, 0x57, 0x8a, 0x4c, 0x70, /* |....W.Lp| */
	0x2b, 0x6b, 0xf1, 0x1d, 0x5f, 0xac, 0x00, 0x00, /* |+k.._...| */
	0x00, 0x00, /* |..| */
}

func visibleChar(c byte) string {
	if 0x20 <= c && c <= 0x7e {
		return string(c)
	} else {
		return "."
	}
}

func dumpBinaryAsGo(d []byte) {
	columns := 8
	var ascii string
	for i := 0; i < len(d); i++ {
		if i%columns == 0 {
			fmt.Printf("\t")
			ascii = ""
		}
		fmt.Printf("0x%02x, ", d[i])
		ascii += visibleChar(d[i])
		if i+1 >= len(d) || ((i+1)%columns == 0 && i > 0) {
			fmt.Printf("/* |%s| */\n", ascii)
		}
	}
}
