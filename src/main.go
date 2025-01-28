package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hiradsajde/hirad-lang/src/parser"
	"github.com/sanity-io/litter"
)

func main() {
	sourceBytes, _ := os.ReadFile("test.lang")
	source := string(sourceBytes)
	start := time.Now()
	ast := parser.Parse(source)
	duration := time.Since(start)
	litter.Dump(ast)
	fmt.Printf("Duration: %v\n", duration)
}
