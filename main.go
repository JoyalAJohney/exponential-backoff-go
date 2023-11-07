package main


import (
	"fmt"
	"time"
)

func operationThatMayFail() error {
	return nil
}


func main() {
	maxRetries := 5
	baseDelay := time.Second
	maxDelay := time.Minute

	for attempt := 0; attempt < maxRetries; attempt++ {
		err := operationThatMayFail()
		if err != nil {
			fmt.Println(err)
			break
		}
		
		nextDelay := exponentialBackoff(baseDelay, maxDelay, attempt)
		fmt.Printf("Operation failed, next retry in %v\n", nextDelay)
		time.Sleep(nextDelay)
	}

	fmt.Println("Operation failed after maximum retries")
}