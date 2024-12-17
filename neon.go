package main

import (
	"compiler/parser"
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println()
	if len(os.Args) < 2 {
		fmt.Println("Please provide a code path and a flag")
		fmt.Println()
		return
	}
	if len(os.Args) > 3 {
		fmt.Println("To many arguments")
		fmt.Println()
		return
	}

	compile := flag.Bool("compile", false, "Compile the code")
	liveness := flag.Bool("liveness", false, "Start liveness analysis")
	constants := flag.Bool("constants", false, "Start constant propagation analysis")

	flag.Parse()

	if !*compile && !*liveness && !*constants {
		fmt.Println("Please specify what the program should do. Use -help if needed")
		fmt.Println()
		return
	}

	var path string

	if *compile {
		if len(os.Args) != 3 {
			fmt.Println("No path provided")
			fmt.Println()
			return
		}
		path = os.Args[2]
		// Send code to tokenizer
		_, parsingSuccesful := parser.Parse(path, true)
		fmt.Println()

		if !parsingSuccesful{
			return
		}

		
	}

	if *liveness {
		fmt.Println("Not implemented")
		fmt.Println()
		return
	}

	if *constants {
		fmt.Println("Not implemented")
		fmt.Println()
		return
	}
}
