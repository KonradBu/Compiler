grammar kbgrammar;

// Program Structure
NAMESPACE: 'namespace' NAME '{' CLASS '}';
USING_BLOCK: 'using' NAME ';' USING_BLOCK | ; // epsilon
CLASS: 'public' 'class' NAME '(' MAIN_ARGS ')' '{' MAIN FUNC_BLOCK '}';
MAIN: 'public' 'static' 'void' 'main';
FUNC_BLOCK: FUNC FUNC_BLOCK | ; // epsion

// Functions
FUNC: 'public' 'static' RETURN_TYPE NAME '(' MAIN_ARGS ')' '{' STATEMENT_BLOCK '}';
FUNC_CALL: NAME '(' CALL_ARGS ')';
MAIN_ARGS: '['']' NAME STRING;
INPUT_ARGS: INPUT_ARG;
INPUT_ARG: ;
CALL_ARGS: ;
CALL_ARG: ;


// Statements
STATEMENT_BLOCK: STATEMENT STATEMENT_BLOCK;
STATEMENT: VARIABLE_DECLARATION | VAR_ASSIGN | FUNC_CALL | WHILE | IF » RETURN;

VARIABLE_DECLARATION: TYPE NAME '=' EXPRESSION;

EXPRESSION: ;



// Controll Flow
WHILE: 'while' '(' EXPRESSION ')' '{' STATEMENT_BLOCK '}';
IF: 'if' '(' EXPRESSION ')' '{' STATEMENT_BLOCK '}';


// Parts

RETURN_TYPE: VOID | TYPE;
TYPE: INT | STRING | BOOL | DOUBLE;

NAME: ;


// Literals

