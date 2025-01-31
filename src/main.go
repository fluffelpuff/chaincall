package main

import (
	"fmt"
)

func main() {
	input := `
string msg = "Hello World";
int zahl = 42;

chcall("https://server2.example.com") {
  const string secret = "Geheim";
  string localVal = "Nur S2 soll das kennen";
}

string after = "Nach dem chcall";
`

	fmt.Println("SOURCE:\n", input)

	lexer := NewLexer(input)
	parser := NewParser(lexer)
	program := parser.ParseProgram()

	if len(parser.errors) > 0 {
		fmt.Println("Parser Errors:")
		for _, e := range parser.errors {
			fmt.Println(" -", e)
		}
	}

	fmt.Println("\n--- PARSED AST ---")
	fmt.Println(program.String())

	// Nun ermitteln wir den Hash über alle Statements, die NICHT im chcall-Block sind.
	mainBlockHash := ComputeMainBlockHash(program)
	fmt.Println("--- MAIN BLOCK HASH ---")
	fmt.Println(mainBlockHash)
}
