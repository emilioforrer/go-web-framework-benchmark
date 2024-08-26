package main

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

func main() {
	file, err := os.Open("out.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Split the file content by "Name:"
	parts := strings.Split(string(content), "Name: ")
	var results []benchmark.Result

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

	// Print the results
	for _, result := range results {
		fmt.Printf("Name: %s\n", result.Name)
		fmt.Printf("Requests Per Second: %.2f\n", result.RequestsPerSecond)
		fmt.Printf("Aborted Due To Deadline: %d\n", result.AbortedDueToDeadline)
		fmt.Printf("Successful Requests: %d\n", result.SuccessfulRequests)
		fmt.Printf("P99.999 Latency: %s\n", result.P99999Latency)
		fmt.Printf("P99 Latency: %s\n", result.P99Latency)
		fmt.Printf("P90 Latency: %s\n", result.P90Latency)
		fmt.Printf("Average Latency: %s\n", result.AverageLatency)
		fmt.Printf("Min Latency: %s\n", result.MinLatency)
		fmt.Printf("Max Latency: %s\n", result.MaxLatency)
		fmt.Printf("Duration: %s\n", result.Duration)
		fmt.Printf("Score: %.2f\n", result.Score)
		fmt.Println("---------------------------------")
	}

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
