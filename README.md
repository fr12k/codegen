# codegen

This is a code generator for golang to generate code for test boilerplate code.

## Usage

Define the following main program in your golang project.
```go
package main

import (
	"github.com/goflink/codegen/pkg/gen"
)

func main() {
	codegen := &gen.CodeGenerator{
		JsonFile:    "./example.json",
		CodeFile:    "example_generated.go",
		Struct:      exampleStruct,
		PackageName: "home",
		ImportCode: `import (
	"github.com/fr12k/home"
)`,
		StructDefintion: "var responseHome =",
	}

	codegen.WriteStructFile()
}
```

Then in the unit test just add the following go instruction.
```go
//go:generate go get github.com/fr12k/codegen
//go:generate go run main.go
//go:generate go fmt example_generated.go
//go:generate go mod tidy
```

