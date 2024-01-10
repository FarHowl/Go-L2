package dev08

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/shirou/gopsutil/process"
)

func Shell() {
	for {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
			return
		}

		input = strings.TrimSpace(input)
		inputParsed := strings.Split(input, " ")

		switch inputParsed[0] {
		case "cd":
			absolutePath, err := filepath.Abs(inputParsed[1])
			if err != nil {
				fmt.Println("Error while trying to get an absolute path: ", err)
			}

			err = os.Chdir(absolutePath)
			if err != nil {
				fmt.Println("Ошибка при изменении текущего каталога: ", err)
				return
			}

			fmt.Println(absolutePath)
		case "echo":
			for i := 1; i < len(inputParsed); i++ {
				fmt.Print(inputParsed[i], " ")
			}
		case "pwd":
			currentPath, err := os.Getwd()
			if err != nil {
				fmt.Println("Error while trying to get current directory path: ", err)
				return
			}
			fmt.Println(currentPath)
		case "kill":
			cmd := exec.Command("taskkill", "/F", "/PID", inputParsed[1])
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err := cmd.Run()
			if err != nil {
				fmt.Println("Ошибка при отправке сигнала:", err)
				return
			}

		case "ps":
			processList, err := process.Processes()
			if err != nil {
				fmt.Println("Error getting all processes: ", err)
			}

			for _, p := range processList {
				pid := p.Pid
				name, _ := p.Name()
				status, _ := p.Status()
				cmdline, _ := p.Cmdline()

				fmt.Printf("PID: %d\n", pid)
				fmt.Printf("Name: %s\n", name)
				fmt.Printf("Status: %s\n", status)
				fmt.Printf("Cmdline: %s\n", cmdline)
				fmt.Println("------------------------------")
			}
		}
	}
}
