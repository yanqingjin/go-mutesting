package branch

import (
	"testing"

	"hsc.philips.com.cn/go-mutation-test/test"
)

func TestMutatorCase(t *testing.T) {
	test.Mutator(
		t,
		MutatorCase,
		"../../testdata/branch/mutatecase.go",
		3,
	)
}
