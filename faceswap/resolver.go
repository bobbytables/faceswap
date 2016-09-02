package faceswap

import (
	"errors"
	"regexp"
)

// ErrNoInterface is an interface to indicate that a search path doesn't
// include a lookupable Interface type
var ErrNoInterface = errors.New("no interface found in search string")

var resolveRegex = regexp.MustCompile(`^"?([^"]+)"?\.([\w\d\-\_]+)$`)

// Resolve handles a search string and turns into a searchable interface
// separated from the package name.
func Resolve(search string) (*Interface, error) {
	matches := resolveRegex.FindAllStringSubmatch(search, 2)
	if matches == nil {
		return nil, ErrNoInterface
	}

	return InterfaceFromPackage(matches[0][1], matches[0][2])
}
