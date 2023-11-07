package main

import (
	"time"
	"math/rand"
)

func exponentialBackoff(baseDelay time.Duration, maxDelay time.Duration, attempt int) time.Duration {
	// left shift 1 (0001) by number of attempts
	exponentialDelay := baseDelay * (1 << attempt)

	// Exponentail delay capped at maxDelay
	if (exponentialDelay > maxDelay) {
		exponentialDelay = maxDelay
	}

	// Adding Jitter
	jitter := time.Duration(rand.Int63n(int64(exponentialDelay) / 2))
	delayeWithJitter := exponentialDelay + jitter

	return delayeWithJitter
}