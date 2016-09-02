# Faceswap

Faceswap is a simple CLI tool that allows you to load a Go interface definition
and output whatever you like using Go templates. It is useful for creating
decorators, middleware for grpc services, and even mocks.

## Installation

To install:

```sh
go get github.com/bobbytables/faceswap
```

## Usage

Faceswap uses Go templates to render its output. You can create whichever
template you like and run Faceswap using it.

Lets make a file called `my_template.txt`:

```
Total Methods on {{ .Name }}: {{ len .Methods }}

Methods:
{{ range .Methods }}
Method: {{ .Name }}, Parameters: {{ varList .Parameters }}, Returns: {{ varList .Returns }}
{{ end }}
```

Now let's try using this file:

```
$ faceswap generate --interface io.ReadWriteCloser --template my_template.txt
Total Methods on ReadWriteCloser: 3

Methods:

Method: Close, Parameters: , Returns: error

Method: Read, Parameters: p []byte, Returns: n int, err error

Method: Write, Parameters: p []byte, Returns: n int, err error

```

You can see Faceswap loaded the internal Go interface `ReadWriteCloser` and
rendered our template file using the result of the interface.

The available variables in the Faceswap template are from the struct
`render.InterfaceTemplate`, [found here](render/template.go)

## Template Functions

Faceswap provides 2 functions that are useful for parameter / return lists.

* `varList .Parameters` - Returns a comma separated list of parameters and their types
* `varListNamed .Parameters "ctx" "in"` - Returns a comma separated list of parameters with the names replaced by the variadic list you provide.

This is nice if you're attempting to render the parameters as they are defined for things such as middlewares or mocks.

## Contributing

Fork and Pull Requests are more than welcome for Faceswap. You can also chat me on the #Gophers Slack channel, @bobbytables
