package main

import (
	"flag"
	"os"
	"time"
)

var paramPrimeNumbersAmount int
var paramJobsAmount int
var paramLogDebug bool

func main() {
	flag.IntVar(&paramPrimeNumbersAmount, "numbers", 100, "amount of prime numbers to generate  (approximated)")
	flag.IntVar(&paramJobsAmount, "jobs", 1, "number of jobs to parallel process")
	flag.BoolVar(&paramLogDebug, "debug", false, "show debug log (slow)")
	flag.Parse()

	err := ValidatePositiveNumber(paramPrimeNumbersAmount)
	if err != nil {
		LogError.Fatalf("Error validating number: %v", err)
	}

	err = ValidatePositiveNumber(paramJobsAmount)
	if err != nil {
		LogError.Fatalf("Error validating jobs: %v", err)
	}

	if paramLogDebug {
		LogDebug.SetOutput(os.Stdout)
	}

	executePrimeNumbers(paramPrimeNumbersAmount, paramJobsAmount)
}

func executePrimeNumbers(numbers, jobsAmount int) {
	LogInfo.Printf("starting generating first %d prime numbers using %d jobs\n", numbers, jobsAmount)
	start := time.Now()

	pc := PrimeNumberCalculator{}
	primeNumbers, errs := pc.GeneratePrimeNumbers(numbers, jobsAmount)
	for err := range errs {
		LogError.Printf("Errors calculating prime numbers: %v", err)
	}

	printPrimeNumbers(primeNumbers)

	elapsed := time.Since(start)
	LogInfo.Printf("took %s to generate %d numbers using %d jobs", elapsed, len(primeNumbers), jobsAmount)
}

func printPrimeNumbers(primeNumbers []int) {
	for _, i := range primeNumbers {
		LogInfo.Printf("prime %d\n", i)
	}
}
