package main

import (
	"errors"
	"flag"
	"log"
	"os"

	"github.com/nobishino/jackc"
)

var tokenize = flag.Bool("t", false, "only tokenization, do not parse")

func main() {
	flag.Parse()
	if err := exec(); err != nil {
		log.Printf("%+v\n", err)
		os.Exit(1)
	}
}

func exec() error {
	var a jackc.Analyzer
	if len(flag.Args()) != 1 {
		return errors.New("exactly 1 argument is required")
	}
	arg := flag.Arg(0) // path
	if *tokenize {
		return a.ExecTokenize(arg)
	}
	return nil
}
