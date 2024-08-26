package benchmark

import (
	"fmt"
	"sort"
	"time"
)

type Result struct {
	Name                 string
	Duration             time.Duration
	SuccessRate          float64
	SuccessfulRequests   int
	MinLatency           time.Duration
	MaxLatency           time.Duration
	AverageLatency       time.Duration
	P90Latency           time.Duration
	P99Latency           time.Duration
	P99999Latency        time.Duration
	AbortedDueToDeadline int
	RequestsPerSecond    float64
	Score                float64
}

type Summary struct {
	SuccessRate    float64 `json:"successRate"`
	Total          float64 `json:"total"`
	Slowest        float64 `json:"slowest"`
	Fastest        float64 `json:"fastest"`
	Average        float64 `json:"average"`
	RequestsPerSec float64 `json:"requestsPerSec"`
	TotalData      int64   `json:"totalData"`
	SizePerRequest int     `json:"sizePerRequest"`
	SizePerSec     float64 `json:"sizePerSec"`
}

type LatencyPercentiles struct {
	P10    float64 `json:"p10"`
	P25    float64 `json:"p25"`
	P50    float64 `json:"p50"`
	P75    float64 `json:"p75"`
	P90    float64 `json:"p90"`
	P95    float64 `json:"p95"`
	P99    float64 `json:"p99"`
	P9999  float64 `json:"p99.9"`
	P99999 float64 `json:"p99.99"`
}

type BenchmarkData struct {
	Summary                Summary            `json:"summary"`
	LatencyPercentiles     LatencyPercentiles `json:"latencyPercentiles"`
	StatusCodeDistribution map[string]int     `json:"statusCodeDistribution"`
	ErrorDistribution      map[string]int     `json:"errorDistribution"`
}

// CalculateAverages takes the grouped results and calculates the average for each service
func CalculateAverages(groupedResults map[string][]Result) []Result {
	averagedResults := make([]Result, 0, len(groupedResults))

	for name, results := range groupedResults {
		var totalDuration time.Duration
		var totalSuccessRate float64
		var totalSuccessfulRequests int
		var totalMinLatency time.Duration
		var totalMaxLatency time.Duration
		var totalAvgLatency time.Duration
		var totalP90Latency time.Duration
		var totalP99Latency time.Duration
		var totalP99999Latency time.Duration
		var totalAbortedDueToDeadline int
		var totalRequestsPerSecond float64
		var totalScore float64

		for _, result := range results {
			totalDuration += result.Duration
			totalSuccessRate += result.SuccessRate
			totalSuccessfulRequests += result.SuccessfulRequests
			totalMinLatency += result.MinLatency
			totalMaxLatency += result.MaxLatency
			totalAvgLatency += result.AverageLatency
			totalP90Latency += result.P90Latency
			totalP99Latency += result.P99Latency
			totalP99999Latency += result.P99999Latency
			totalAbortedDueToDeadline += result.AbortedDueToDeadline
			totalRequestsPerSecond += result.RequestsPerSecond
			totalScore += result.Score
		}

		count := len(results)
		averagedResults = append(averagedResults, Result{
			Name:                 name,
			Duration:             totalDuration / time.Duration(count),
			SuccessRate:          totalSuccessRate / float64(count),
			SuccessfulRequests:   totalSuccessfulRequests / count,
			MinLatency:           totalMinLatency / time.Duration(count),
			MaxLatency:           totalMaxLatency / time.Duration(count),
			AverageLatency:       totalAvgLatency / time.Duration(count),
			P90Latency:           totalP90Latency / time.Duration(count),
			P99Latency:           totalP99Latency / time.Duration(count),
			P99999Latency:        totalP99999Latency / time.Duration(count),
			AbortedDueToDeadline: totalAbortedDueToDeadline / count,
			RequestsPerSecond:    totalRequestsPerSecond / float64(count),
			Score:                totalScore / float64(count),
		})
	}

	return averagedResults
}

func PrintResult(result Result) {
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

// CalculateP99 takes the grouped results and calculates the P99 for each service
func CalculateP99(groupedResults map[string][]Result) []Result {
	p99Results := make([]Result, 0, len(groupedResults))

	for name, results := range groupedResults {
		// Sort each metric to find P99
		p99Result := Result{Name: name}

		// Sorting functions for each metric
		sort.Slice(results, func(i, j int) bool {
			return results[i].Duration < results[j].Duration
		})
		p99Result.Duration = calculatePercentile(results, func(r Result) time.Duration { return r.Duration })

		sort.Slice(results, func(i, j int) bool {
			return results[i].SuccessRate < results[j].SuccessRate
		})
		p99Result.SuccessRate = calculatePercentileFloat64(results, func(r Result) float64 { return r.SuccessRate })

		sort.Slice(results, func(i, j int) bool {
			return results[i].SuccessfulRequests < results[j].SuccessfulRequests
		})
		p99Result.SuccessfulRequests = calculatePercentileInt(results, func(r Result) int { return r.SuccessfulRequests })

		sort.Slice(results, func(i, j int) bool {
			return results[i].MinLatency < results[j].MinLatency
		})
		p99Result.MinLatency = calculatePercentile(results, func(r Result) time.Duration { return r.MinLatency })

		sort.Slice(results, func(i, j int) bool {
			return results[i].MaxLatency < results[j].MaxLatency
		})
		p99Result.MaxLatency = calculatePercentile(results, func(r Result) time.Duration { return r.MaxLatency })

		sort.Slice(results, func(i, j int) bool {
			return results[i].AverageLatency < results[j].AverageLatency
		})
		p99Result.AverageLatency = calculatePercentile(results, func(r Result) time.Duration { return r.AverageLatency })

		sort.Slice(results, func(i, j int) bool {
			return results[i].P90Latency < results[j].P90Latency
		})
		p99Result.P90Latency = calculatePercentile(results, func(r Result) time.Duration { return r.P90Latency })

		sort.Slice(results, func(i, j int) bool {
			return results[i].P99Latency < results[j].P99Latency
		})
		p99Result.P99Latency = calculatePercentile(results, func(r Result) time.Duration { return r.P99Latency })

		sort.Slice(results, func(i, j int) bool {
			return results[i].P99999Latency < results[j].P99999Latency
		})
		p99Result.P99999Latency = calculatePercentile(results, func(r Result) time.Duration { return r.P99999Latency })

		sort.Slice(results, func(i, j int) bool {
			return results[i].AbortedDueToDeadline < results[j].AbortedDueToDeadline
		})
		p99Result.AbortedDueToDeadline = calculatePercentileInt(results, func(r Result) int { return r.AbortedDueToDeadline })

		sort.Slice(results, func(i, j int) bool {
			return results[i].RequestsPerSecond < results[j].RequestsPerSecond
		})
		p99Result.RequestsPerSecond = calculatePercentileFloat64(results, func(r Result) float64 { return r.RequestsPerSecond })

		sort.Slice(results, func(i, j int) bool {
			return results[i].Score < results[j].Score
		})
		p99Result.Score = calculatePercentileFloat64(results, func(r Result) float64 { return r.Score })

		p99Results = append(p99Results, p99Result)
	}

	return p99Results
}

// Helper function to calculate P99 for time.Duration metrics
func calculatePercentile(results []Result, getValue func(Result) time.Duration) time.Duration {
	index := int(0.99 * float64(len(results)))
	return getValue(results[index])
}

// Helper function to calculate P99 for float64 metrics
func calculatePercentileFloat64(results []Result, getValue func(Result) float64) float64 {
	index := int(0.99 * float64(len(results)))
	return getValue(results[index])
}

// Helper function to calculate P99 for int metrics
func calculatePercentileInt(results []Result, getValue func(Result) int) int {
	index := int(0.99 * float64(len(results)))
	return getValue(results[index])
}
