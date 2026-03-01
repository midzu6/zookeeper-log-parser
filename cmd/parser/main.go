package main

import (
	"fmt"
	"logParse/internal/parser"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Использование: go run cmd/parser/main.go <testdata/zookiper.log>")
		os.Exit(1)
	}

	filePath := os.Args[1]

	count, err := parser.ReadAndCountEntries(filePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
	}
	fmt.Printf("\nГотово! Найдено %d записей в логе.\n", count)
}
