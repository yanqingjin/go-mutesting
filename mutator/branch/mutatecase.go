package branch

import (
	"go/ast"
	"go/types"

	"hsc.philips.com.cn/go-mutation-test/util"
	"hsc.philips.com.cn/go-mutation-test/mutator"
)

func init() {
	mutator.Register("branch/case", MutatorCase)
}

// MutatorCase implements a mutator for case clauses.
func MutatorCase(pkg *types.Package, info *types.Info, node ast.Node) []mutator.Mutation {
	n, ok := node.(*ast.CaseClause)
	if !ok {
		return nil
	}

	old := n.Body

	return []mutator.Mutation{
		{
			Change: func() {
				n.Body = []ast.Stmt{
					util.CreateNoopOfStatements(pkg, info, n.Body),
				}
			},
			Reset: func() {
				n.Body = old
			},
		},
	}
}
