package faceswap

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResolver(t *testing.T) {
	testcases := []struct {
		scenario      string
		path          string
		packagePath   string
		interfaceName string
	}{
		{
			scenario:      "Stdlib is loaded correctly",
			path:          "fmt.Stringer",
			packagePath:   "fmt",
			interfaceName: "Stringer",
		},
		{
			scenario:      "External libs are loaded correctly",
			path:          "github.com/bobbytables/faceswap/faceswap/dummy.FakeInterface",
			packagePath:   "github.com/bobbytables/faceswap/faceswap/dummy",
			interfaceName: "FakeInterface",
		},
		{
			scenario:      "External libs are loaded correctly when the package is quoted",
			path:          `"github.com/bobbytables/faceswap/faceswap/dummy".FakeInterface`,
			packagePath:   "github.com/bobbytables/faceswap/faceswap/dummy",
			interfaceName: "FakeInterface",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.scenario, func(t *testing.T) {
			i, err := Resolve(tc.path)

			require.Nil(t, err)
			assert.Equal(t, tc.packagePath, i.Package.Path())
			assert.Equal(t, tc.interfaceName, i.Name)
		})
	}
}
