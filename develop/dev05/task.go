package dev05

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

func Grep() {
	after := flag.Int("A", 0, "Print N lines after each match")
	before := flag.Int("B", 0, "Print N lines before each match")
	context := flag.Int("C", 0, "Print ±N lines around each match")
	count := flag.Bool("c", false, "Print only a count of selected lines")
	ignoreCase := flag.Bool("i", false, "Ignore case distinctions")
	invert := flag.Bool("v", false, "Invert the sense of matching")
	fixed := flag.Bool("F", false, "Interpret pattern as a literal string")
	lineNumber := flag.Bool("n", false, "Print line numbers")

	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Usage: grep [OPTIONS] PATTERN FILE [FILE...]")
		flag.PrintDefaults()
		os.Exit(1)
	}

	patternStr := args[0]
	filePath := "develop/dev05/" + args[1]

	var pattern *regexp.Regexp
	if *fixed {
		patternStr = regexp.QuoteMeta(patternStr)
	}
	if *ignoreCase {
		pattern = regexp.MustCompile("(?i)" + patternStr)
	} else {
		pattern = regexp.MustCompile(patternStr)
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	var lineNumberCounter int

	for scanner.Scan() {
		line := scanner.Text()
		lineNumberCounter++

		if pattern.MatchString(line) != *invert {
			if *count {
				continue
			}

			if *lineNumber {
				line = fmt.Sprintf("%d:%s", lineNumberCounter, line)
			}

			if *before > 0 || *after > 0 || *context > 0 {
				lines = append(lines, line)
				printLines(lines, *before, *after, *context)
				lines = nil
			} else {
				fmt.Println(line)
			}
		} else {
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func printLines(lines []string, before, after, context int) {
	for i, line := range lines {
		if i >= before && i < len(lines)-after {
			fmt.Println(line)
		}
	}
}
