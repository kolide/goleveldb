package iterator_test

import (
	"testing"

	"github.com/kolide/goleveldb/leveldb/testutil"
)

func TestIterator(t *testing.T) {
	testutil.RunSuite(t, "Iterator Suite")
}
