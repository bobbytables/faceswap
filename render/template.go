package template

import "github.com/bobbytables/faceswap/faceswap"

// InterfaceTemplate represents an easy-to-template value of an actual
// Go interface.
type InterfaceTemplate struct {
	Name    string
	Methods []faceswap.Method
}
