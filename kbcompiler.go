package main

import (
	"compiler/parser"
	"fmt"
	"flag"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a code path and a flag")
		return
	}

	compile := flag.Bool("compile", false, "Compile the code")
	liveness := flag.Bool("liveness", false, "Start liveness analysis")
	constants := flag.Bool("constants", false, "Start constant propagation analysis")

	flag.Parse()

	var path string

	if len(os.Args) < 4 {
		path = os.Args[2]
	}else{
		fmt.Println("No provided path")
		return
	}

	if !*compile && !*liveness && !*constants {
		fmt.Println("Please specify what the program should do. Use -help if needed")
		return
	}

	if *compile {
		// Send code to tokenizer
		parser.Parse(path, true)
	}

	if *liveness {
		panic ("Not implemented")
	}

	if *constants {
		panic ("Not implemented")
	}
}
