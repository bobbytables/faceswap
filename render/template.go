package render

import (
	"errors"
	"fmt"
	"strings"
	"text/template"

	"github.com/bobbytables/faceswap/faceswap"
)

// ErrVarCountMismatch is an error returned when varListNamed is passed a mismatch
// of tuples vs names
var ErrVarCountMismatch = errors.New("tuples and name count does not match")

// InterfaceTemplate represents an easy-to-template value of an actual
// Go interface.
type InterfaceTemplate struct {
	Name    string
	Methods []faceswap.Method
}

// RenderFuncs returns the default rendering funcs for go templates
var RenderFuncs = template.FuncMap{
	"varList":      varList,
	"varListNamed": varListNamed,
}

func varList(tuples []faceswap.Tuple) string {
	var args []string
	for _, t := range tuples {
		formatted := fmt.Sprintf("%s %s", t.Name, t.Type.String())
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
