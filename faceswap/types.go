package faceswap

import (
	"go/types"

	"github.com/pkg/errors"
	"golang.org/x/tools/go/loader"
)

// ErrNotAnInterface is an error that is returned when the loaded type is
// not an interface
var ErrNotAnInterface = errors.New("not an interface")

// Interface contains the logic for easily iterating over methods defined
// on a *types.Interface.
// This allows for easy interaction inside of a Go template-esque file.
type Interface struct {
	iface *types.Interface
	Name  string
}

// Method represents a method on an interface in Go
type Method struct {
	Name       string
	Parameters []Tuple
	Returns    []Tuple
}

// Tuple represents something that can be used as a parameter or return value
type Tuple struct {
	Name string
	Type types.Type
}

// InterfaceFromPackage loads a package and looks up an interface inside of it
func InterfaceFromPackage(path, name string) (*Interface, error) {
	var conf loader.Config
	conf.Import(path)
	prog, err := conf.Load()
	if err != nil {
		return nil, errors.Wrap(err, "could not load package")
	}

	pkgInfo := prog.Package(path)
	obj := pkgInfo.Pkg.Scope().Lookup(name)
	if obj == nil {
		return nil, errors.New("could not find interface")
	}

	if !types.IsInterface(obj.Type()) {
		return nil, ErrNotAnInterface
	}

	return NewInterface(name, obj.Type().Underlying().(*types.Interface)), nil
}

// NewInterface returns an initialized Interface from an internal Go type.
func NewInterface(name string, i *types.Interface) *Interface {
	return &Interface{
		Name:  name,
		iface: i,
	}
}

// Methods returns all of the methods on the interface
func (i *Interface) Methods() []Method {
	var methods []Method

	for n := 0; n < i.iface.NumExplicitMethods(); n++ {
		f := i.iface.ExplicitMethod(n)
		methods = append(methods, methodFromFunc(f))
	}

	return methods
}

func methodFromFunc(f *types.Func) Method {
	return Method{
		Name:       f.Name(),
		Parameters: parametersFromFunc(f),
		Returns:    returnsFromFunc(f),
	}
}

func parametersFromFunc(f *types.Func) []Tuple {
	var params []Tuple

	s := f.Type().(*types.Signature)
	for i := 0; i < s.Params().Len(); i++ {
		v := s.Params().At(i)
		params = append(params, Tuple{Name: v.Name(), Type: v.Type()})
	}

	return params
}

func returnsFromFunc(f *types.Func) []Tuple {
	var returns []Tuple

	s := f.Type().(*types.Signature)
	for i := 0; i < s.Results().Len(); i++ {
		v := s.Results().At(i)
		returns = append(returns, Tuple{Name: v.Name(), Type: v.Type()})
	}

	return returns
}
