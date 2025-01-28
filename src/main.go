package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hiradsajde/hirad-lang/src/lexer"
	"github.com/sanity-io/litter"
	"github.com/tlaceby/hiradsajde/hirad-lang/src/parser"
)

func main() {
	sourceBytes, _ := os.ReadFile("test.lang")
	source := string(sourceBytes)
	start := time.Now()
	ast := parser.Parse(source)
	tokens := lexer.Tokenize(source)
	duration := time.Since(start)
	fmt.Println("Lexical Analysis (Tokens):")

	fmt.Println("Parse Tree: ")
	litter.Dump(ast)
	fmt.Printf("Duration: %v\n", duration)
}
