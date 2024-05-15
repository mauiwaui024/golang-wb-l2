package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$ ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		if input == "\\quit" {
			break
		}

		// Разделение ввода пользователя на команды конвейера
		commands := strings.Split(input, "|")
		var lastOutput []byte

		for _, cmdStr := range commands {
			cmdStr = strings.TrimSpace(cmdStr)
			args := strings.Fields(cmdStr)
			if len(args) == 0 {
				continue
			}

			switch args[0] {
			case "cd":
				if len(args) < 2 {
					fmt.Println("Usage: cd <directory>")
					continue
				}
				err := os.Chdir(args[1])
				if err != nil {
					fmt.Println("Error:", err)
				}
				continue
			case "pwd":
				dir, err := os.Getwd()
				if err != nil {
					fmt.Println("Error:", err)
					continue
				}
				fmt.Println(dir)
				continue
			case "echo":
				if len(args) < 2 {
					fmt.Println("Usage: echo <message>")
					continue
				}
				fmt.Println(strings.Join(args[1:], " "))
				continue
			case "kill":
				if len(args) < 2 {
					fmt.Println("Usage: kill <pid>")
					continue
				}
				pid := args[1]
				pidInt, err := strconv.Atoi(pid)
				if err != nil {
					fmt.Println("Error:", err)
					continue
				}
				err = syscall.Kill(pidInt, syscall.SIGKILL)
				if err != nil {
					fmt.Println("Error:", err)
				}
				continue
			case "ps":
				cmd := exec.Command("ps")
				output, err := cmd.Output()
				if err != nil {
					fmt.Println("Error:", err)
					continue
				}
				fmt.Println(string(output))
				continue
			}

			var cmd *exec.Cmd
			if len(commands) > 1 {
				cmd = exec.Command("bash", "-c", cmdStr)
			} else {
				cmd = exec.Command(args[0], args[1:]...)
			}

			// Если есть предыдущий вывод, установить его как ввод для этой команды
			if lastOutput != nil {
				cmd.Stdin = strings.NewReader(string(lastOutput))
			}

			// Запуск команды и обработка ошибок
			output, err := cmd.Output()
			if err != nil {
				fmt.Println("Error:", err)
				break
			}

			// Сохранение вывода для использования в следующей команде
			lastOutput = output
		}

		// Вывод результатов выполнения всех команд конвейера
		fmt.Println(string(lastOutput))
	}
}
