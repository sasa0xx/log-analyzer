package main

import (
	"strconv"
	"strings"
)

type Analyzer struct {
	result map[string]int
}

func NewAnalyzer() *Analyzer {
	tmp := Analyzer{result: map[string]int{"ERROR": 0, "INFO": 0, "DEBUG": 0, "WARN": 0}}
	return &tmp
}

func (a *Analyzer) Add(entry LogEntry) {
	a.result[entry.Level] += 1
}

func (a *Analyzer) Summarize() string {
	var builder strings.Builder
	builder.WriteString("Summary:\n")
	for k, v := range a.result {
		if v != 0 {
			builder.WriteString(k + " : " + strconv.Itoa(v) + "\n")
		}
	}
	return builder.String()
}
