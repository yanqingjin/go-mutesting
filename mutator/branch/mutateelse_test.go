package branch

import (
	"testing"

	"hsc.philips.com.cn/go-mutation-test/test"
)

func TestMutatorElse(t *testing.T) {
	test.Mutator(
		t,
		MutatorElse,
		"../../testdata/branch/mutateelse.go",
		1,
	)
}
