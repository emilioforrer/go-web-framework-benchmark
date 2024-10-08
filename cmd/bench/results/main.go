package main

import (
	"fmt"
	"local/go-benchmarks/internal/benchmark"
	"local/go-benchmarks/internal/core"
)

func main() {
	results, err := core.ImportResults()
	if err != nil {
		fmt.Println("Error at ImportResults:", err)
		return
	}

	// Print the results
	for _, result := range results {
		benchmark.PrintResult(result)
	}
}
