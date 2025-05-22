package main

import (
	"fmt"
	"math"
	"slices"
	"sync"
	"time"
)

// WorkerProgress holds progress information for a worker
type WorkerProgress struct {
	WorkerID int
	Progress float64
}

// isPrime checks if a number is prime.
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	// Optimization: only check up to square root of n
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 5; i <= sqrtN; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// isPrimeWorker checks a range of numbers for primality and sends found primes to a channel.
func isPrimeWorker(start, end int, primeChan chan<- int, wg *sync.WaitGroup, progressChan chan<- WorkerProgress, workerID int) {
	defer wg.Done()

	total := end - start + 1
	processed := 0

	for num := start; num <= end; num++ {
		if isPrime(num) {
			primeChan <- num
		}
		processed++

		if processed%1000 == 0 {
			progress := float64(processed) / float64(total) * 100
			progressChan <- WorkerProgress{WorkerID: workerID, Progress: progress}
		}
	}
	// Send final 100% progress update for this worker
	progressChan <- WorkerProgress{WorkerID: workerID, Progress: 100.0}
}

func main() {
	start := 1
	end := 100_000
	numWorkers := 4
	progressStep := 10.0
	sampleSize := 5

	primeChan := make(chan int)
	progressChan := make(chan WorkerProgress)
	var wg sync.WaitGroup

	startTime := time.Now()

	rangePerWorker := (end - start + 1) / numWorkers

	// Goroutine to track and display progress
	go func() {
		progressMap := make(map[int]float64)
		lastDisplayedProgress := 0.0

		for progress := range progressChan {
			progressMap[progress.WorkerID] = progress.Progress

			totalProgress := 0.0
			for _, p := range progressMap {
				totalProgress += p
			}
			totalProgress /= float64(numWorkers)

			// Display progress every progressPercentage
			if totalProgress >= lastDisplayedProgress+progressStep {
				fmt.Printf("\rProgress: %.2f%%\n", totalProgress)
				lastDisplayedProgress = totalProgress
			}
		}
	}()

	// Spawn worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		workerStart := start + i*rangePerWorker
		workerEnd := workerStart + rangePerWorker - 1

		if i == numWorkers-1 {
			workerEnd = end
		}

		go isPrimeWorker(workerStart, workerEnd, primeChan, &wg, progressChan, i)
	}

	// Close primeChan when all workers are done
	go func() {
		wg.Wait()
		close(primeChan)
		close(progressChan)
	}()

	// Collect and store primes
	primes := []int{}
	for prime := range primeChan {
		primes = append(primes, prime)
	}

	// Print results
	fmt.Println("\n\nTime taken:", time.Since(startTime))
	fmt.Printf("Total prime numbers found: %d\n", len(primes))

	// Peek the results
	slices.Sort(primes)
	fmt.Printf("First %d prime numbers found: %v\n", sampleSize, primes[:sampleSize])
	fmt.Printf("Last %d prime numbers found: %v\n", sampleSize, primes[len(primes)-sampleSize:])
}
