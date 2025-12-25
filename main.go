package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

func main() {
	file, err := os.Open("logs.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	var wg sync.WaitGroup

	jobs := make(chan string)
	results := make(chan LogEntry)
	numWorkers := 4
	analyzer := NewAnalyzer()

	for range make([]struct{}, numWorkers) {
		go worker(jobs, results)
	}
	go func() {
		for entry := range results {
			analyzer.Add(entry)
			wg.Done()
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		wg.Add(1)
		jobs <- line
	}
	close(jobs)
	wg.Wait()
	close(results)

	fmt.Println(analyzer.Summarize())
}
