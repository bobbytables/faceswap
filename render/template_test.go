package render

import (
	"go/types"
	"testing"

	"github.com/bobbytables/faceswap/faceswap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVarListFunc(t *testing.T) {
	testcases := []struct {
		scenario string
		expected string
		tuples   []faceswap.Tuple
	}{
		{
			scenario: "Unnamed variable doesn't include name in list",
			expected: "string",
			tuples:   []faceswap.Tuple{{Name: "", Type: types.Typ[types.String]}},
		},
		{
			scenario: "Named variables include name and type",
			expected: "s string",
			tuples:   []faceswap.Tuple{{Name: "s", Type: types.Typ[types.String]}},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.scenario, func(t *testing.T) {
			result := varList(tc.tuples)

			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestVarListNamedFunc(t *testing.T) {
	testcases := []struct {
		scenario  string
		expected  string
		tuples    []faceswap.Tuple
		namedVars []string
	}{
		{
			scenario:  "Replaces an empty name with the parameter at the position",
			expected:  "s string",
			tuples:    []faceswap.Tuple{{Name: "", Type: types.Typ[types.String]}},
			namedVars: []string{"s"},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.scenario, func(t *testing.T) {
			result, err := varListNamed(tc.tuples, tc.namedVars...)
			require.Nil(t, err)
			assert.Equal(t, tc.expected, result)
		})
	}
}
