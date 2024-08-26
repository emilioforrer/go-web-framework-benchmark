package benchmark

func CalculateScore(r Result) float64 {
	// Assume 100% success rate since it's given
	successRate := r.SuccessRate

	// Convert latency durations to microseconds for more precise calculations
	minLatency := float64(r.MinLatency.Microseconds())
	maxLatency := float64(r.MaxLatency.Microseconds())
	avgLatency := float64(r.AverageLatency.Microseconds())
	p90Latency := float64(r.P90Latency.Microseconds())
	p99Latency := float64(r.P99Latency.Microseconds())
	p99999Latency := float64(r.P99999Latency.Microseconds())
	requestsPerSecond := r.RequestsPerSecond

	// Calculate the spread between max and min latency
	latencySpread := maxLatency - minLatency

	// Apply weights to P99.999 latency and Requests Per Second
	weightedRequestsPerSecond := requestsPerSecond * 2.5
	weightedP99999Latency := p99999Latency * 3.0 // Increase the weight of P99.999 latency
	weightedP99Latency := p99Latency * 1.1       // Increase the weight of P99 latency
	weightedP90Latency := p90Latency * 1.0       // Increase the weight of P90 latency
	weightedAvgLatency := avgLatency * 0.9       // Increase the weight of Average Latency
	reducedLatencySpread := latencySpread * 0.8  // // Reduce the impact of latency spread

	// Optimize the impact of SuccessfulRequests and AbortedDueToDeadline
	successFactor := float64(r.SuccessfulRequests) * 1.2
	abortPenalty := float64(r.AbortedDueToDeadline) * 1.5

	// Calculate the score using the modified formula:
	// score := (successRate * (weightedRequestsPerSecond*1000000 + float64(r.SuccessfulRequests))) /
	// 	(minLatency + maxLatency + weightedAvgLatency + weightedP90Latency + weightedP99Latency + weightedP99999Latency + reducedLatencySpread)

	score := (successRate * (weightedRequestsPerSecond + successFactor - abortPenalty)) /
		(weightedAvgLatency + weightedP90Latency + weightedP99Latency + weightedP99999Latency + reducedLatencySpread)

	return score
}
