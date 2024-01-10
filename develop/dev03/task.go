package dev03

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Sort() {
	sortColumn := flag.Int("k", 0, "Номер колонки для сортировки (по умолчанию 0, разделитель - пробел)")
	sortNumeric := flag.Bool("n", false, "Сортировать по числовому значению")
	reverse := flag.Bool("r", false, "Сортировать в обратном порядке")
	unique := flag.Bool("u", false, "Не выводить повторяющиеся строки")
	flag.Parse()

	filePath := "C:/Users/dima3/OneDrive/Документы/GitHub/Go-L2/develop/dev03/example.txt"
	filePathToWrite := "C:/Users/dima3/OneDrive/Документы/GitHub/Go-L2/develop/dev03/sorted.txt"

	lines, err := readLines(filePath)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	if *sortColumn < 0 || *sortColumn >= len(lines[0]) {
		fmt.Println("Некорректный номер колонки для сортировки.")
		return
	}

	comparator := func(i, j int) bool {
		if *sortNumeric {
			var valI int
			for _, str := range lines[i] {
				val, err := strconv.Atoi(str)
				if err != nil {
					continue
				} else {
					valI = val
				}
			}

			var valJ int
			for _, str := range lines[j] {
				val, err := strconv.Atoi(str)
				if err != nil {
					continue
				} else {
					valJ = val
				}
			}

			return valI < valJ
		} else {
			return lines[i][*sortColumn] < lines[j][*sortColumn]
		}
	}

	sort.Slice(lines, comparator)

	if *reverse {
		reverseSlice(lines)
	}

	if *unique {
		lines = uniqueSlice(lines)
	}

	err = writeLines(filePathToWrite, lines)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
		return
	}

	fmt.Println("Сортировка завершена успешно.")
}

func readLines(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		lines = append(lines, words)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func writeLines(filePath string, lines [][]string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, words := range lines {
		line := strings.Trim(strings.Join(words, " "), "[]")
		fmt.Fprintln(writer, line)
	}

	return writer.Flush()
}

func reverseSlice(slice [][]string) {
	for i := 0; i < len(slice)/2; i++ {
		j := len(slice) - i - 1
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func uniqueSlice(slice [][]string) [][]string {
	encountered := make(map[string]struct{})
	result := [][]string{}

	for _, words := range slice {
		line := strings.Join(words, " ")
		if _, found := encountered[line]; !found {
			encountered[line] = struct{}{}
			result = append(result, words)
		}
	}

	return result
}
