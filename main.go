package main

import (
	"fmt"
	"go/types"

	"github.com/bobbytables/faceswap/faceswap"

	"golang.org/x/tools/go/loader"
)

var pkgPath = "github.com/bobbytables/gangway/events"

func main() {
	var conf loader.Config
	conf.Import(pkgPath)
	prog, err := conf.Load()
	if err != nil {
		panic(err)
	}

	pkgInfo := prog.Package(pkgPath)
	obj := pkgInfo.Pkg.Scope().Lookup("Listener")
	t := obj.Type()
	fmt.Printf("%+v\n", t.Underlying())
	iface, ok := t.Underlying().(*types.Interface)
	if !ok {
		panic("Type searched was not an interface")
	}

	fi := faceswap.NewInterface("Listener", iface)
	fmt.Println(fi.Methods()[0].Name)

	fmt.Printf("PkgInfo: %+v\n", obj)
}
