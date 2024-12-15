# NEON
Compiler from a subset of C# to JVM Bytecode (Jasmin)

For the course compiler construction at HHU Düsseldorf

By Konrad Burgi


## How to run and compile
Build:
go build neon.go

Run:
./main -compile [filepath]

-liveness for variable liveness analysis
-constant for constant propogation analysis

## Info
Uses go 1.23.2 
Implements a SRL(0) parser from scratch

## File Explaination:
lexer:
lexer.go Takes a file and generates the corresponding tokens

parser:
parser.go Manages the Parser and Grammar Construction. Takes Tokens and gives them into the constructed SLR Parsing Table

grammar.go defines the grammar struct and includes several helper functions, including FIRST and FOLLOW

grammarConstructor.go Handels the actual grammar used transforms the rules into the nececary structs etc.

parser_constructor.go Provides a interface for parser.go to define build the different grammar features

slr_automata.go defines the SLR Automata and functions for its constructions

slr_parsing_table.go defines the parsing table and utility functions, including taking the SLR automata and transforming it into the table

stack.go provides a stack for parsing with the parsing table