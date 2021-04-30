package branch

import (
	"go/ast"
	"go/types"

	"hsc.philips.com.cn/go-mutation-test/util"
	"hsc.philips.com.cn/go-mutation-test/mutator"
)

func init() {
	mutator.Register("branch/else", MutatorElse)
}

// MutatorElse implements a mutator for else branches.
func MutatorElse(pkg *types.Package, info *types.Info, node ast.Node) []mutator.Mutation {
	n, ok := node.(*ast.IfStmt)
	if !ok {
		return nil
	}
	// We ignore else ifs and nil blocks
	_, ok = n.Else.(*ast.IfStmt)
	if ok || n.Else == nil {
		return nil
	}

	old := n.Else

	return []mutator.Mutation{
		{
			Change: func() {
				n.Else = util.CreateNoopOfStatement(pkg, info, old)
			},
			Reset: func() {
				n.Else = old
			},
		},
	}
}
