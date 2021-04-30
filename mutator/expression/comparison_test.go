package expression

import (
	"testing"

	"hsc.philips.com.cn/go-mutation-test/test"
)

func TestMutatorComparison(t *testing.T) {
	test.Mutator(
		t,
		MutatorComparison,
		"../../testdata/expression/comparison.go",
		4,
	)
}
