package render

import (
	"errors"
	"fmt"
	"go/types"
	"html/template"
	"strings"

	"github.com/bobbytables/faceswap/faceswap"
)

// ErrVarCountMismatch is an error returned when varListNamed is passed a mismatch
// of tuples vs names
var ErrVarCountMismatch = errors.New("tuples and name count does not match")

// InterfaceTemplate represents an easy-to-template value of an actual
// Go interface.
type InterfaceTemplate struct {
	Name    string
	Package *types.Package
	Methods []faceswap.Method
}

// RenderFuncs returns the default rendering funcs for go templates
var RenderFuncs = template.FuncMap{
	"varList":      varList,
	"varListNamed": varListNamed,
}

// ShortQualifier handles pretty-printing types for varLists in templates.
// The default printing of types.Type includes the fully qualified package name.
// This returns only the exported package name to prevent type names being
// *github.com/bobbytables/faceswap/type.Type and instead make them *type.Type
// Checkout: https://godoc.org/go/types#Qualifier
func ShortQualifier(p *types.Package) string {
	return p.Name()
}

func varList(tuples []faceswap.Tuple) string {
	var args []string
	for _, t := range tuples {
		typ := types.TypeString(t.Type, ShortQualifier)
		formatted := fmt.Sprintf("%s %s", t.Name, typ)
		args = append(args, strings.TrimSpace(formatted))
	}

	return strings.Join(args, ", ")
}

func varListNamed(tuples []faceswap.Tuple, names ...string) (string, error) {
	if len(tuples) != len(names) {
		return "", ErrVarCountMismatch
	}

	var ts []faceswap.Tuple
	for i, tuple := range tuples {
		ts = append(ts, faceswap.Tuple{Name: names[i], Type: tuple.Type})
	}

	return varList(ts), nil
}
