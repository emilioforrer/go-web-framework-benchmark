package main

import (
	"encoding/json"
	"fmt"
	"local/go-benchmarks/internal/benchmark"
	"log"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	// Path to the results-history directory
	directory := "results-history"

	// Read and group results by Name
	groupedResults, err := readResults(directory)
	if err != nil {
		log.Fatalf("Failed to read results: %v", err)
	}

	results := benchmark.CalculateP99(groupedResults)

	// Sort the results by score in descending order
	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	// Print or use the averaged results
	for _, result := range results {
		benchmark.PrintResult(result)
	}
}

// readResults reads all JSON files from the results-history directory and groups them by Name
func readResults(directory string) (map[string][]benchmark.Result, error) {
	groupedResults := make(map[string][]benchmark.Result)

	// Read all files from the directory
	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, fmt.Errorf("error reading directory: %w", err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			filePath := filepath.Join(directory, file.Name())
			content, err := os.ReadFile(filePath)
			if err != nil {
				return nil, fmt.Errorf("error reading file %s: %w", file.Name(), err)
			}

			var resultArray []benchmark.Result
			err = json.Unmarshal(content, &resultArray)
			if err != nil {
				return nil, fmt.Errorf("error unmarshalling JSON from file %s: %w", file.Name(), err)
			}

			// Group results by Name
			for _, result := range resultArray {
				groupedResults[result.Name] = append(groupedResults[result.Name], result)
			}
		}
	}

	return groupedResults, nil
}
