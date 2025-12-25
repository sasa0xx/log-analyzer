package main

import (
	"log"
	"slices"
	"strings"
)

func ParseLogLine(line string, analyzer *Analyzer) LogEntry {
	splits := strings.SplitN(line, " ", 2)
	level := splits[0]
	if !slices.Contains([]string{"ERROR", "INFO", "DEBUG", "WARN"}, level) {
		log.Fatalf("Could not recognize level \"%s\"", level)
	}

	entry := LogEntry{level, splits[1]}
	analyzer.Add(entry)
	return entry
}
