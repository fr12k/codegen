package main

import (
	"flag"

	"github.com/fr12k/codegen/pkg/gen"
)

type v struct {
	Name string `json:"name"`
}

func main() {
	v := &v{}

	jsonFile := flag.String("json", "", "json file")

	flag.Parse()

	codegen := &gen.CodeGenerator{
		JsonFile:    *jsonFile,
		CodeFile:    "example_generated.go",
		Struct:      v,
		PackageName: "home",
		ImportCode: `import (
	"github.com/fr12k/home"
)`,
		StructDefintion: "var responseHome =",
	}

	codegen.WriteStructFile()
}
