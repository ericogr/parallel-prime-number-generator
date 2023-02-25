package main

import (
	"flag"
	"os"
	"time"
)

var paramPrimeNumbersAmount int
var paramParallelJobs int
var paramLogDebug bool

func main() {
	flag.IntVar(&paramPrimeNumbersAmount, "primes", 100, "amount of prime numbers to generate (approximated)")
	flag.IntVar(&paramParallelJobs, "jobs", 1, "number of parallel jobs to generate prime numbers")
	flag.BoolVar(&paramLogDebug, "debug", false, "show debug log (slow)")
	flag.Parse()

	err := ValidateGreaterThanZero(paramPrimeNumbersAmount)
	if err != nil {
		LogError.Fatalf("Error validating prime numbers amount: %v", err)
	}

	err = ValidateGreaterThanZero(paramParallelJobs)
	if err != nil {
		LogError.Fatalf("Error validating parallel jobs: %v", err)
	}

	if paramLogDebug {
		LogDebug.SetOutput(os.Stdout)
	}

	generatePrimeNumbers(paramPrimeNumbersAmount, paramParallelJobs)
}

func generatePrimeNumbers(numbers, parallelJobs int) {
	LogInfo.Printf("starting generating first %d prime numbers using %d jobs\n", numbers, parallelJobs)
	start := time.Now()

	pg := PrimeNumberGenerator{}
	primeNumbers, errs := pg.GeneratePrimeNumbers(numbers, parallelJobs)
	for err := range errs {
		LogError.Printf("Errors calculating prime numbers: %v", err)
	}

	printPrimeNumbers(primeNumbers)

	elapsed := time.Since(start)
	LogInfo.Printf("took %s to test %d numbers and generate %d prime numbers using %d jobs", elapsed, pg.lastNumberProcessed, len(primeNumbers), parallelJobs)
}

func printPrimeNumbers(primeNumbers []int) {
	for _, i := range primeNumbers {
		LogInfo.Printf("prime %d\n", i)
	}
}
