package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/bobbytables/faceswap/faceswap"
	"github.com/bobbytables/faceswap/render"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

const (
	version string = "0.1.0"
)

var (
	interfaceSearch string
	templateFile    string
)

func main() {
	app := cli.NewApp()
	app.Name = "faceswap"
	app.Usage = "Takes Go interfaces and allows you to create files from them using Go templating."
	app.Version = version
	app.Authors = []cli.Author{
		{
			Name:  "Robert Ross",
			Email: "robert@creativequeries.com",
		},
	}

	app.Commands = []cli.Command{generateCmd()}
	app.Run(os.Args)
}

func generateCmd() cli.Command {
	return cli.Command{
		Name:      "generate",
		ShortName: "g",
		Usage: `The name of a package and interface defined within it to be loaded

		Using the fmt.Stringer interface would look like:

		"fmt".Stringer

		A custom interface lookup could be like:

		"github.com/bobbytables/faceswap/faceswap/dummy".FakeInterface`,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "interface",
				Usage:       "The interface lookup",
				Destination: &interfaceSearch,
			},
			cli.StringFlag{
				Name:        "template",
				Usage:       "The file defining the template to be used on generation",
				Destination: &templateFile,
			},
		},
		Action: generate,
	}
}

func generate(cli *cli.Context) error {
	i, err := faceswap.Resolve(interfaceSearch)
	if err != nil {
		return errors.Wrap(err, "could not resolve package")
	}

	it := render.InterfaceTemplate{
		Name:    i.Name,
		Methods: i.Methods(),
		Package: i.Package,
	}

	content, err := ioutil.ReadFile(templateFile)
	if err != nil {
		return err
	}

	t, err := template.New("custom").
		Funcs(render.RenderFuncs).
		Parse(string(content))

	if err != nil {
		return errors.Wrap(err, "could not parse template")
	}

	err = t.Execute(os.Stdout, it)
	if err != nil {
		return errors.Wrap(err, "could not execute template")
	}
	fmt.Println()
	return nil
}
