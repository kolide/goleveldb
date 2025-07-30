// Copyright (c) 2012, Suryandaru Triandana <syndtr@gmail.com>
// All rights reserved.
//
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package leveldb

import (
	"bytes"
	"testing"
)

func decodeEncode(v *sessionRecord) (res bool, err error) {
	b := new(bytes.Buffer)
	err = v.encode(b)
	if err != nil {
		return
	}
	v2 := &sessionRecord{}
	err = v.decode(b)
	if err != nil {
		return
	}
	b2 := new(bytes.Buffer)
	err = v2.encode(b2)
	if err != nil {
		return
	}
	return bytes.Equal(b.Bytes(), b2.Bytes()), nil
}

func TestSessionRecord_EncodeDecode(t *testing.T) {
	big := int64(1) << 50
	v := &sessionRecord{}
	i := int64(0)
	test := func() {
		res, err := decodeEncode(v)
		if err != nil {
			t.Fatalf("error when testing encode/decode sessionRecord: %v", err)
		}
		if !res {
			t.Error("encode/decode test failed at iteration:", i)
		}
	}

	for ; i < 4; i++ {
		test()
		ik1, err1 := makeInternalKey(nil, []byte("foo"), uint64(big+500+1), keyTypeVal)
		if err1 != nil {
			panic(err1) // Test code - should not fail with valid parameters
		}
		ik2, err2 := makeInternalKey(nil, []byte("zoo"), uint64(big+600+1), keyTypeDel)
		if err2 != nil {
			panic(err2) // Test code - should not fail with valid parameters
		}
		v.addTable(3, big+300+i, big+400+i, ik1, ik2)
		v.delTable(4, big+700+i)
		ik3, err3 := makeInternalKey(nil, []byte("x"), uint64(big+900+1), keyTypeVal)
		if err3 != nil {
			panic(err3) // Test code - should not fail with valid parameters
		}
		v.addCompPtr(int(i), ik3)
	}

	v.setComparer("foo")
	v.setJournalNum(big + 100)
	v.setPrevJournalNum(big + 99)
	v.setNextFileNum(big + 200)
	v.setSeqNum(uint64(big + 1000))
	test()
}
