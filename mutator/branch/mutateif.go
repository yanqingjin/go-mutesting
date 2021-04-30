package branch

import (
	"go/ast"
	"go/types"

	"hsc.philips.com.cn/go-mutation-test/util"
	"hsc.philips.com.cn/go-mutation-test/mutator"
)

func init() {
	mutator.Register("branch/if", MutatorIf)
}

// MutatorIf implements a mutator for if and else if branches.
func MutatorIf(pkg *types.Package, info *types.Info, node ast.Node) []mutator.Mutation {
	n, ok := node.(*ast.IfStmt)
	if !ok {
		return nil
	}

	old := n.Body.List

	return []mutator.Mutation{
		{
			Change: func() {
				n.Body.List = []ast.Stmt{
					util.CreateNoopOfStatement(pkg, info, n.Body),
				}
			},
			Reset: func() {
				n.Body.List = old
			},
		},
	}
}
