package dev09

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Wget() {
	if len(os.Args) != 2 {
		fmt.Println("Использование: go run main.go <URL>")
		return
	}

	url := os.Args[1]
	err := downloadFile(url)
	if err != nil {
		fmt.Println("Ошибка загрузки файла:", err)
	}
}

func downloadFile(url string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("ошибка: %s", response.Status)
	}

	file, err := os.Create(getFileName(url))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Файл успешно загружен: %s\n", getFileName(url))
	return nil
}

func getFileName(url string) string {
	lastIndex := 0
	for i := len(url) - 1; i >= 0; i-- {
		if url[i] == '/' {
			lastIndex = i + 1
			break
		}
	}
	return url[lastIndex:]
}
