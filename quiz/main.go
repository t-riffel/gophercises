package main

import (
	"flag"
	"os"
)

func main() {
	// quiz problems passed in through command-line
	args := os.Args[1:]
	namePtr := flag.String("file", "problems.csv", "filename")
	flag.Parse()

}
