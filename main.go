package main

import (
	"Compilerbau/Tokenizer"
	"fmt"
)

func main() {
	path := "code/Simple1.java"
	fmt.Println(Tokenizer.Tokenize(path))
}
