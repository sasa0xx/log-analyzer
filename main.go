package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("logs.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	analyzer := NewAnalyzer()
	for scanner.Scan() {
		line := scanner.Text()
		ParseLogLine(line, analyzer)
	}
	print(analyzer.Summarize())
}
