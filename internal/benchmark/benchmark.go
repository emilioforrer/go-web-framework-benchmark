package benchmark

import "time"

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
