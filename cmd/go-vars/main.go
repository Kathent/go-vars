package main

import (
	"os"

	"github.com/Kathent/go-vars/initial"
)

var verbose bool

func main() {
	initial.InitialFlag()

	if len(os.Args) < 1 {
		panic("at leaset 1 arguments required")
	}
}
