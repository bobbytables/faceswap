package faceswap

import (
	"go/types"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInterfaceExposesMethods(t *testing.T) {
	iface, err := InterfaceFromPackage("github.com/bobbytables/faceswap/faceswap/dummy", "FakeInterface")
	require.Nil(t, err)

	availableMethods := []string{
		"A_Hello",
		"B_AnError",
		"C_MultiReturn",
		"D_MultiReturnCustomType",
		"E_CustomReturn",
	}
	sort.Strings(availableMethods)

	t.Run("Methods() returns the methods on the interface", func(t *testing.T) {
		assert.Equal(t, "FakeInterface", iface.Name)
		require.Len(t, iface.Methods(), len(availableMethods))
		assert.Equal(t, availableMethods[0], iface.Methods()[0].Name, "method name is correct")
	})

	t.Run("Methods() has the parameters for the method embedded", func(t *testing.T) {
		method := iface.Methods()[0]

		assert.Equal(t, availableMethods[0], method.Name)
		require.Len(t, method.Parameters, 1, "parameters on the method")
		assert.IsType(t, types.Typ[types.String], method.Parameters[0].Type)
	})

	t.Run("Methods() has custom types for parameters", func(t *testing.T) {
		method := iface.Methods()[4]

		assert.Equal(t, availableMethods[4], method.Name)
		require.Len(t, method.Parameters, 1, "parameters on the method")
		assert.Equal(t, "*github.com/bobbytables/faceswap/faceswap/dummy.CustomType", method.Parameters[0].Type.String())
	})

	t.Run("Methods() has custom types for returns", func(t *testing.T) {
		method := iface.Methods()[3]

		assert.Equal(t, availableMethods[3], method.Name)
		require.Len(t, method.Returns, 2, "returns on the method")
		assert.Equal(t, "*github.com/bobbytables/faceswap/faceswap/dummy.CustomType", method.Returns[0].Type.String())
		assert.Equal(t, "error", method.Returns[1].Type.String())
	})
}
