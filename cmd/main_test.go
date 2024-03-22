package main

import (
	"os"
	"testing"
)

func TestMainFunc(t *testing.T) {
	os.Args = []string{"cmd", "-json", "./testdata/t.json"}
	defer func() {
		os.Remove("./example_generated.go")
	}()
	main()
}
