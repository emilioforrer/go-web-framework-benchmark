package core

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"local/go-benchmarks/internal/benchmark"
	"os"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"
)

func ImportResults() ([]benchmark.Result, error) {
	var results []benchmark.Result
	statsFilePath := "stats.txt"
	stats, err := ReadContainerStatsFromFile(statsFilePath)
	if err != nil {
		fmt.Println("Error opening file:", statsFilePath)
	}

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

		stat, err := GetStatByName(stats, name)

		if err == nil {
			result.Memory = ParseUsage(stat.MemUsage)
			result.CPU = ParseUsage(stat.CPUPerc)
			result.Network = ParseUsage(stat.NetIO)
			result.Disk = ParseUsage(stat.BlockIO)
		}
		result.ResourceUtilizationScore = benchmark.CalculateResourceUtilizationScore(result)
		result.TotalScore = benchmark.CalculateTotalScore(result.PerformanceScore, result.ResourceUtilizationScore)

		results = append(results, result)
	}

	// Sort the results by score in descending order
	sort.Slice(results, func(i, j int) bool {
		return results[i].TotalScore > results[j].TotalScore
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

	if r.SuccessRate == 0.0 {
		fmt.Printf("Warning: Service '%s' has a success rate of 0. Please ensure that the service is up and running correctly.\n", r.Name)
		return r, nil
	}

	// Calculate the score
	r.PerformanceScore = benchmark.CalculateScore(r)
	return r, nil
}

// ReadContainerStatsFromFile reads the container stats from a file and returns a slice of ContainerStats or an error
func ReadContainerStatsFromFile(filePath string) ([]benchmark.ContainerStat, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var statsSlice []benchmark.ContainerStat

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Read each line of the file
		line := scanner.Text()

		// Unmarshal the JSON data into the struct
		var stats benchmark.ContainerStat
		err := json.Unmarshal([]byte(line), &stats)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal line: %w", err)
		}

		// Append the unmarshaled struct to the slice
		statsSlice = append(statsSlice, stats)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return statsSlice, nil
}

// GetStatByName searches for a container by its name using the slices package
func GetStatByName(stats []benchmark.ContainerStat, name string) (*benchmark.ContainerStat, error) {
	index := slices.IndexFunc(stats, func(stat benchmark.ContainerStat) bool {
		suffix := fmt.Sprintf("%s-1", name)

		return strings.HasSuffix(stat.Name, suffix)
	})

	if index != -1 {
		return &stats[index], nil
	}

	return nil, fmt.Errorf("container with name %s not found", name)
}

func ParseUsage(usage string) float64 {
	var value float64
	re := regexp.MustCompile(`\d+\.\d+`)

	// Find the first match
	match := re.FindString(usage)
	if match == "" {
		return 0.0
	}

	value, _ = strconv.ParseFloat(match, 64)

	pattern := `[KMGT]iB`
	re = regexp.MustCompile(pattern)
	formatMatch := re.FindString(usage)
	if match != "" {
		switch formatMatch {
		case "KiB":
			value = value / 1024
		case "GiB":
			value = value * 1024
		case "TiB":
			value = value * 1024 * 1024
		}
	}

	return value
}
