package dev06

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func Cut() {
	fields := flag.String("f", "", "Select fields (columns)")
	delimiter := flag.String("d", "\t", "Use a different delimiter")
	separated := flag.Bool("s", false, "Only output lines containing delimiter")
	flag.Parse()

	selectedFields := parseFields(*fields)
	selectedDelimiter := *delimiter

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if *separated && !strings.Contains(line, selectedDelimiter) {
			continue
		}

		fields := strings.Split(line, selectedDelimiter)
		selectedFieldsValues := selectFields(fields, selectedFields)

		fmt.Println(strings.Join(selectedFieldsValues, selectedDelimiter))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
		os.Exit(1)
	}
}

func parseFields(fieldsStr string) []int {
	var selectedFields []int

	if fieldsStr == "" {
		return selectedFields
	}

	fields := strings.Split(fieldsStr, ",")
	for _, field := range fields {
		if fieldRange := strings.Split(field, "-"); len(fieldRange) > 1 {
			// Handle field range (e.g., 1-3)
			start, end := parseInt(fieldRange[0]), parseInt(fieldRange[1])
			for i := start; i <= end; i++ {
				selectedFields = append(selectedFields, i)
			}
		} else {
			// Handle single field
			selectedFields = append(selectedFields, parseInt(field))
		}
	}

	return selectedFields
}

func selectFields(fields []string, selectedFields []int) []string {
	var selectedFieldsValues []string
	for _, fieldIndex := range selectedFields {
		if fieldIndex > 0 && fieldIndex <= len(fields) {
			selectedFieldsValues = append(selectedFieldsValues, fields[fieldIndex-1])
		} else {
			selectedFieldsValues = append(selectedFieldsValues, "")
		}
	}
	return selectedFieldsValues
}

func parseInt(s string) int {
	var result int
	_, err := fmt.Sscanf(s, "%d", &result)
	if err != nil {
		return 0
	}
	return result
}
