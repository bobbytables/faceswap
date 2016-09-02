package render

import (
	"go/types"
	"testing"

	"github.com/bobbytables/faceswap/faceswap"
	"github.com/stretchr/testify/assert"
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
