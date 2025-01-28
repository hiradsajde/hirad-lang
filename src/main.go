package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hiradsajde/hirad-lang/src/lexer"
	"github.com/hiradsajde/hirad-lang/src/parser"
	"github.com/sanity-io/litter"
)

func main() {
	sourceBytes, _ := os.ReadFile("test.lang")
	source := string(sourceBytes)
	start := time.Now()

	fmt.Println("\n------- Lexycal Analysis Tokens -------")
	tokens := lexer.Tokenize(source)
	litter.Dump(tokens)

	fmt.Println("\n------- Parse Tree -------")
	ast := parser.Parse(source)
	litter.Dump(ast)

	duration := time.Since(start)
	fmt.Printf("Duration: %v\n", duration)
}
