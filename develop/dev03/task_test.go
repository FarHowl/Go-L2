package dev03

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestXxx(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "-n", "-k", "-r")

	cmd.Env = append(os.Environ(), "GOFLAGS=-mod=mod")

	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Ошибка выполнения команды: %v\n%s", err, out)
	}

	if !strings.Contains(string(out), "Сортировка завершена успешно.") {
		t.Error("Ожидалась успешная сортировка, но не найдено сообщение.")
	}
}
