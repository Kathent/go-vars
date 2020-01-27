package main

import (
	"fmt"
	"os"

	"github.com/Kathent/go-vars/tools/vars"

	"github.com/Kathent/go-vars/initial"
)

var verbose bool

func main() {
	initial.InitialFlag()

	var filePath string
	var src interface{}
	if initial.PathOrIn() {
		src = os.Stdin
	} else {
		if len(os.Args) < 2 {
			panic("at leaset 1 arguments required")
		}
		filePath = os.Args[1]
	}
	names, err := vars.ReplaceVarWithDef(filePath, src)
	if err != nil {
		panic(err)
	}

	for k, v := range names {
		fmt.Print(v)
		if k != len(names)-1 {
			fmt.Print(",")
		}
	}
}
