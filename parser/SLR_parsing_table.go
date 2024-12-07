package parser

import "fmt"

type SLR_parsing_Table struct {
	actionTable map[int]map[string]*Action
	gotoToTable map[int]map[string]int
}

type Action struct {
	// Accept / Shift / Reduce
	actionType string
	value      int
}

func MakeSLRTable() *SLR_parsing_Table {

}

func MakeAction(ty string, value int) Action {
	newAction := new(Action)
	newAction.actionType = ty
	newAction.value = value
	return *newAction
}

func (table *SLR_parsing_Table) AddAction(state int, terminal string, actionType string, ActionValue int) {
	newAction := MakeAction(actionType, ActionValue)
	if table.actionTable[state][terminal] != nil{
		panic("Grammar does not seem to be SLR Parsable")
	} 
	table.actionTable[state][terminal] = &newAction
}

func (table *SLR_parsing_Table) AddGoTo(state int, symbol string, newstate int) {
	table.gotoToTable[state][symbol] = newstate
}

func (table *SLR_parsing_Table) PrintTable() {
	fmt.Println("Action: ")
	for i, m := range table.actionTable{
		fmt.Print(i)
		for str, action := range m{
			fmt.Print(str + " ")
			fmt.Print(action.actionType + " ")
			fmt.Print(action.value)
			fmt.Print(" ")
		} 
		fmt.Println()
	}

	fmt.Println("GoTo: ")
	for i, m := range table.gotoToTable{
		fmt.Print(i)
		for str, state := range m{
			fmt.Print(str + " ")
			fmt.Print(state)
			fmt.Print(" ")
		} 
		fmt.Println()
	}
}