package main

import (
	"fmt"
	"sync"
)

type PrimeNumberCalculator struct {
	primeNumbersList []int
	errorList        []error
}

type PrimeNumberResult struct {
	Number  int
	IsPrime bool
	Error   error
}

func (p *PrimeNumberCalculator) IsPrimeNumber(number int) (bool, error) {
	if number <= 0 {
		return false, fmt.Errorf("invalid number %d", number)
	}

	if number == 1 {
		return false, nil
	}

	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false, nil
		}
	}

	return true, nil
}

func (p *PrimeNumberCalculator) createJobs(primeNumbersAmount int, jobChannel chan<- int, resultChannel <-chan PrimeNumberResult) {
	number := 2
	for len(p.primeNumbersList) < primeNumbersAmount {
		LogDebug.Printf("creating job %d (%d/%d)\n", number, len(p.primeNumbersList), primeNumbersAmount)
		jobChannel <- number
		LogDebug.Printf("job %d created\n", number)
		number++
	}
	LogDebug.Printf("closing job channel\n")
	close(jobChannel)
}

func (p *PrimeNumberCalculator) processResults(jobChannel <-chan int, resultChannel <-chan PrimeNumberResult) {
	for result := range resultChannel {
		LogDebug.Printf("processing result %d\n", result.Number)

		if result.Error != nil {
			p.errorList = append(p.errorList, result.Error)
			continue
		}

		if result.IsPrime {
			p.primeNumbersList = append(p.primeNumbersList, result.Number)
		}
	}
}

func (p *PrimeNumberCalculator) processJob(wg *sync.WaitGroup, jobChannel <-chan int, resultChannel chan<- PrimeNumberResult) {
	for number := range jobChannel {
		LogDebug.Printf("processing job %d\n", number)
		isPrimeNumber, err := p.IsPrimeNumber(number)
		resultChannel <- PrimeNumberResult{
			Number:  number,
			IsPrime: isPrimeNumber,
			Error:   err,
		}
		LogDebug.Printf("pushed job %d result\n", number)
	}
	LogDebug.Printf("processing job done\n")
	wg.Done()
}

func (p *PrimeNumberCalculator) createJobPool(jobsAmount int, jobChannel <-chan int, resultChannel chan<- PrimeNumberResult) {
	var wg sync.WaitGroup
	wg.Add(jobsAmount)
	for i := 0; i < jobsAmount; i++ {
		LogDebug.Printf("creating job processor %d\n", i)
		go p.processJob(&wg, jobChannel, resultChannel)
	}
	wg.Wait()
	LogDebug.Printf("closing result channel\n")
	close(resultChannel)
}

func (p *PrimeNumberCalculator) GeneratePrimeNumbers(primeNumbersAmount, jobsAmount int) ([]int, []error) {
	p.primeNumbersList = nil
	p.errorList = nil

	jobChannel := make(chan int, jobsAmount)
	resultChannel := make(chan PrimeNumberResult, jobsAmount)

	go p.createJobs(primeNumbersAmount, jobChannel, resultChannel)
	go p.processResults(jobChannel, resultChannel)
	p.createJobPool(jobsAmount, jobChannel, resultChannel)

	return p.primeNumbersList, p.errorList
}
