package testing

import (
	"fmt"
	"math"
	"net/http"
	"runtime"
	"sort"
	"sync"
	"testing"
	"time"
)

var numberOfRequests = 10000

type Service struct {
	Name string
	URL  string
}

type Result struct {
	Duration           time.Duration
	SuccessfulRequests int
	FailedRequests     int
	MinLatency         time.Duration
	MaxLatency         time.Duration
	AverageLatency     time.Duration
	P90Latency         time.Duration
	P99Latency         time.Duration
	P99999Latency      time.Duration
	RequestsPerSecond  float64
	Score              float64
}

func calculatePercentile(latencies []time.Duration, percentile float64) time.Duration {
	index := int(math.Ceil(float64(len(latencies))*percentile)) - 1
	return latencies[index]
}

func calculateScore(r Result) float64 {
	// Calculate the success rate (percentage of successful requests)
	successRate := float64(r.SuccessfulRequests) / float64(r.SuccessfulRequests+r.FailedRequests)

	// Convert latency durations to microseconds for more precise calculations
	minLatency := float64(r.MinLatency.Microseconds())
	maxLatency := float64(r.MaxLatency.Microseconds())
	avgLatency := float64(r.AverageLatency.Microseconds())
	p90Latency := float64(r.P90Latency.Microseconds())
	p99Latency := float64(r.P99Latency.Microseconds())
	p99999Latency := float64(r.P99999Latency.Microseconds())

	// Calculate the spread between max and min latency
	// A smaller spread indicates more consistent performance
	latencySpread := maxLatency - minLatency

	// Calculate the score using the formula:
	// (Success Rate * Requests Per Second * Scaling Factor) / (Sum of Latency Metrics)
	//
	// This formula aims to balance throughput (RPS) and latency:
	// - Higher success rates and RPS increase the score
	// - Lower latencies (including spread) increase the score
	//
	// The scaling factor (1,000,000) is used to make the score more readable
	// and to balance the magnitude difference between RPS and latency values
	score := (successRate * r.RequestsPerSecond * 1000000) /
		(minLatency + maxLatency + avgLatency + p90Latency + p99Latency + p99999Latency + latencySpread)

	// Explanation of the formula components:
	// - successRate: Rewards higher percentage of successful requests
	// - r.RequestsPerSecond: Rewards higher throughput
	// - 1000000: Scaling factor to balance RPS and latency magnitudes
	// - Sum of latencies in denominator: Penalizes higher latencies
	// - latencySpread in denominator: Penalizes inconsistent performance

	// A higher score indicates better overall performance, considering both
	// throughput (RPS) and various latency metrics

	return score
}

func BenchmarkRequests(b *testing.B) {
	fmt.Println("Deprecated method for benchmarking. Use the new method described in the README.md file instead")

	services := []Service{
		{"GoFr", "http://localhost:8080"},
		{"Bun", "http://localhost:8081"},
		{"Echo", "http://localhost:8082"},
		{"Fuego", "http://localhost:8083"},
		{"Fiber", "http://localhost:8084"},
		{"STD", "http://localhost:8085"},
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	results := make(map[Service]Result)

	for _, service := range services {
		b.Run(fmt.Sprintf("Requests to %s (%s)", service.Name, service.URL), func(b *testing.B) {
			runtime.GC()
			b.ResetTimer()

			var totalDuration time.Duration
			totalSuccessfulRequests := 0
			totalFailedRequests := 0
			var latencies []time.Duration
			var minLatency, maxLatency time.Duration

			for i := 0; i < b.N; i++ {
				var wg sync.WaitGroup
				resultChan := make(chan time.Duration, numberOfRequests)

				start := time.Now()

				for j := 0; j < numberOfRequests; j++ {
					wg.Add(1)
					go func() {
						defer wg.Done()
						reqStart := time.Now()
						resp, err := client.Get(service.URL)
						reqDuration := time.Since(reqStart)
						if err == nil && resp.StatusCode == http.StatusOK {
							resultChan <- reqDuration
							resp.Body.Close()
						} else {
							resultChan <- -reqDuration // Negative duration for failed requests
							if resp != nil {
								resp.Body.Close()
							}
						}
					}()
				}

				wg.Wait()
				close(resultChan)

				duration := time.Since(start)
				totalDuration += duration

				for latency := range resultChan {
					if latency >= 0 {
						totalSuccessfulRequests++
						latencies = append(latencies, latency)
						if minLatency == 0 || latency < minLatency {
							minLatency = latency
						}
						if latency > maxLatency {
							maxLatency = latency
						}
					} else {
						totalFailedRequests++
						latencies = append(latencies, -latency) // Store absolute value
					}
				}
			}

			totalRequests := totalSuccessfulRequests + totalFailedRequests
			averageLatency := totalDuration / time.Duration(totalRequests)

			// Calculate percentile latencies
			sort.Slice(latencies, func(i, j int) bool { return latencies[i] < latencies[j] })
			p90Latency := calculatePercentile(latencies, 0.90)
			p99Latency := calculatePercentile(latencies, 0.99)
			p99999Latency := calculatePercentile(latencies, 0.99999)

			// Calculate requests per second
			rps := float64(totalRequests) / totalDuration.Seconds()

			result := Result{
				Duration:           totalDuration,
				SuccessfulRequests: totalSuccessfulRequests,
				FailedRequests:     totalFailedRequests,
				MinLatency:         minLatency,
				MaxLatency:         maxLatency,
				AverageLatency:     averageLatency,
				P90Latency:         p90Latency,
				P99Latency:         p99Latency,
				P99999Latency:      p99999Latency,
				RequestsPerSecond:  rps,
			}
			result.Score = calculateScore(result)
			results[service] = result
		})
	}

	// Rank services
	var rankedServices []Service
	for service := range results {
		rankedServices = append(rankedServices, service)
	}
	sort.Slice(rankedServices, func(i, j int) bool {
		return results[rankedServices[i]].Score > results[rankedServices[j]].Score
	})

	fmt.Printf("\nRanked Results:\n")
	for i, service := range rankedServices {
		result := results[service]
		fmt.Printf("%d. %s (%s):\n", i+1, service.Name, service.URL)
		fmt.Printf("   Score: %.2f\n", result.Score)
		fmt.Printf("   Requests Per Second: %.2f\n", result.RequestsPerSecond)
		fmt.Printf("   Total Duration: %v\n", result.Duration)
		fmt.Printf("   Successful Requests: %d\n", result.SuccessfulRequests)
		fmt.Printf("   Failed Requests: %d\n", result.FailedRequests)
		fmt.Printf("   Min Latency: %v\n", result.MinLatency)
		fmt.Printf("   Max Latency: %v\n", result.MaxLatency)
		fmt.Printf("   Average Latency: %v\n", result.AverageLatency)
		fmt.Printf("   P90 Latency: %v\n", result.P90Latency)
		fmt.Printf("   P99 Latency: %v\n", result.P99Latency)
		fmt.Printf("   P99.999 Latency: %v\n\n", result.P99999Latency)
	}

	fmt.Printf("Best performing service: %s (%s) (Score: %.2f, RPS: %.2f)\n",
		rankedServices[0].Name,
		rankedServices[0].URL,
		results[rankedServices[0]].Score,
		results[rankedServices[0]].RequestsPerSecond)
}
