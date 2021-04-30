package branch

import (
	"testing"

	"hsc.philips.com.cn/go-mutation-test/test"
)

func TestMutatorIf(t *testing.T) {
	test.Mutator(
		t,
		MutatorIf,
		"../../testdata/branch/mutateif.go",
		2,
	)
}
