package initial

import "flag"

var verbose bool
var readFromStdin bool

func Init() {
	InitialFlag()
}

func InitialFlag() {
	flag.BoolVar(&verbose, "v", false, "print debug log")
	flag.BoolVar(&readFromStdin, "i", false, "read from stdin")
	flag.Parse()
}

func PathOrIn() bool {
	return readFromStdin
}
