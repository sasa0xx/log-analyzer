package main

func worker(jobs <-chan string, results chan<- LogEntry) {
	for line := range jobs {
		entry := ParseLogLine(line)
		results <- entry
	}
}
