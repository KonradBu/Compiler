package main

import (
	"compiler/parser"
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a code path and a flag")
		return
	}
	if len(os.Args) > 3 {
		fmt.Println("To many arguments")
		return
	}

	compile := flag.Bool("compile", false, "Compile the code")
	liveness := flag.Bool("liveness", false, "Start liveness analysis")
	constants := flag.Bool("constants", false, "Start constant propagation analysis")

	flag.Parse()

	if !*compile && !*liveness && !*constants {
		fmt.Println("Please specify what the program should do. Use -help if needed")
		return
	}

	var path string

	if *compile {
		if len(os.Args) != 3 {
			fmt.Println("No path provided")
			return
		}
		path = os.Args[0]
		fmt.Println("Parsing...")
		// Send code to tokenizer
		parser.Parse(path, true)
	}

	if *liveness {
		fmt.Println("Not implemented")
		return
	}

	if *constants {
		fmt.Println("Not implemented")
		return
	}
}
