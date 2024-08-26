package core

import (
	"encoding/json"
	"fmt"
	"io"
	"local/go-benchmarks/internal/benchmark"
	"os"
	"sort"
	"strings"
	"time"
)

func ImportResults() ([]benchmark.Result, error) {
	var results []benchmark.Result

	file, err := os.Open("out.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return results, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return results, err
	}

	// Split the file content by "Name:"
	parts := strings.Split(string(content), "Name: ")

	for _, part := range parts[1:] { // Skip the first empty split
		lines := strings.SplitN(part, "\n", 2)
		name := strings.TrimSpace(lines[0])
		data := lines[1]

		// Parse the JSON data
		result, err := parseBenchmarkData(name, []byte(data))
		if err != nil {
			fmt.Printf("Error parsing data for %s: %v\n", name, err)
			continue
		}

		results = append(results, result)
	}

	// Sort the results by score in descending order
	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})
	return results, nil
}

func ExportResults(results []benchmark.Result) {
	// Get the current time
	timestamp := time.Now().Format("2006-01-02-15-04-05")
	// Construct the file name with the timestamp
	fileName := fmt.Sprintf("results-history/result-%s.json", timestamp)

	jsonData, err := json.Marshal(results)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	// Create a new file for writing
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	file.Write(jsonData)

	defer file.Close()
}

func parseBenchmarkData(name string, data []byte) (benchmark.Result, error) {
	var benchmarkData benchmark.BenchmarkData
	if err := json.Unmarshal(data, &benchmarkData); err != nil {
		return benchmark.Result{}, err
	}

	r := benchmark.Result{
		Name:                 name,
		Duration:             time.Duration(benchmarkData.Summary.Total * float64(time.Second)),
		AbortedDueToDeadline: benchmarkData.ErrorDistribution["aborted due to deadline"],
		SuccessfulRequests:   benchmarkData.StatusCodeDistribution["200"],
		SuccessRate:          float64(benchmarkData.Summary.SuccessRate),
		MinLatency:           time.Duration(benchmarkData.Summary.Fastest * float64(time.Second)),
		MaxLatency:           time.Duration(benchmarkData.Summary.Slowest * float64(time.Second)),
		AverageLatency:       time.Duration(benchmarkData.Summary.Average * float64(time.Second)),
		P90Latency:           time.Duration(benchmarkData.LatencyPercentiles.P90 * float64(time.Second)),
		P99Latency:           time.Duration(benchmarkData.LatencyPercentiles.P99 * float64(time.Second)),
		P99999Latency:        time.Duration(benchmarkData.LatencyPercentiles.P99999 * float64(time.Second)),
		RequestsPerSecond:    benchmarkData.Summary.RequestsPerSec,
	}

	// Calculate the score
	r.Score = benchmark.CalculateScore(r)
	return r, nil
}
