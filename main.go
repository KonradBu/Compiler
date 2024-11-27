package main

import (
	_ "compiler/automata"
	_ "compiler/lexer"
	"compiler/parser"

	//compiler/util"
	"flag"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("Please provide a code path and no other inputs")
	}
	path := os.Args[2]
	compile := flag.Bool("compile", false, "Compile the code")
	liveness := flag.Bool("liveness", false, "Start liveness analysis")
	constants := flag.Bool("constants", false, "Start constant propagation analysis")

	flag.Parse()

	if !*compile && !*liveness && !*constants {
		panic("Please specify what the program should do. Use -help if needed")
	}
	if *compile {
		// Send code to tokenizer
		parser.Parse(path)
	}

	if *liveness {
		//
	}

	if *constants {

	}
}
