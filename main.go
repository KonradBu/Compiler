package compiler

import (
	"compiler/util"


	_ "compiler/tokenizer"
	_ "errors"
	_ "fmt"
)

func main() {
	newNFA, err := util.MakeAutomata()
	if (err != nil){
		panic("Wrong Input")
	}

}
