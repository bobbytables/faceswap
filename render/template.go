package render

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/bobbytables/faceswap/faceswap"
)

// InterfaceTemplate represents an easy-to-template value of an actual
// Go interface.
type InterfaceTemplate struct {
	Name    string
	Methods []faceswap.Method
}

// RenderFuncs returns the default rendering funcs for go templates
var RenderFuncs = template.FuncMap{
	"varList": varList,
}

func varList(tuples []faceswap.Tuple) string {
	var args []string
	for _, t := range tuples {
		formatted := fmt.Sprintf("%s %s", t.Name, t.Type.String())
		args = append(args, strings.TrimSpace(formatted))
	}

	return strings.Join(args, ", ")
}
