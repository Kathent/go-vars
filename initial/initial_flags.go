package initial

import "flag"

var verbose bool

func Init() {
	InitialFlag()
}

func InitialFlag() {
	flag.BoolVar(&verbose, "v", false, "print debug log")
	flag.Parse()
}
