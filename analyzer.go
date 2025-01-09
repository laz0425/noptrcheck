package noptrcheck

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var Analyzer = &analysis.Analyzer{
	Name:     "noptrcheck",
	Doc:      "checks for struct with pointer fields used as map keys or set elements",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			mapType, ok := n.(*ast.MapType)
			if !ok {
				return true
			}

			keyType := pass.TypesInfo.TypeOf(mapType.Key)
			if keyType == nil {
				return true
			}

			if _, isPointer := keyType.Underlying().(*types.Pointer); isPointer {
				pass.Reportf(mapType.Pos(),
					"a pointer type cannot be used as a map key")
				return true
			}

			if structType, ok := keyType.Underlying().(*types.Struct); ok {
				if hasPointerField(pass, structType, make(map[string]bool)) {
					structName := keyType.String()
					pass.Reportf(mapType.Pos(),
						"%s has pointer fields and cannot be used as a map key",
						structName)
				}
			}
			return true
		})
	}
	return nil, nil
}

func hasPointerField(pass *analysis.Pass, structType *types.Struct, visited map[string]bool) bool {
	structName := structType.String()

	if visited[structName] {
		return false
	}
	visited[structName] = true

	for i := 0; i < structType.NumFields(); i++ {
		field := structType.Field(i)
		fieldType := field.Type()

		if isPointerType(fieldType) {
			return true
		}

		if nestedStruct, ok := fieldType.Underlying().(*types.Struct); ok {
			if hasPointerField(pass, nestedStruct, visited) {
				return true
			}
		}
	}

	return false
}

func isPointerType(t types.Type) bool {
	switch tt := t.(type) {
	case *types.Pointer:
		return true
	case *types.Named:
		return isPointerType(tt.Underlying())
	default:
		return false
	}
}
