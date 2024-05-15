package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	// Получение аргументов командной строки (хост и порт)
	if len(os.Args) < 3 {
		fmt.Println("Usage: go-telnet <host> <port> [--timeout <timeout>]")
		os.Exit(1)
	}
	host := os.Args[1]
	port := os.Args[2]

	// Обработка таймаута
	var timeout time.Duration = 10 * time.Second
	if len(os.Args) == 5 && os.Args[3] == "--timeout" {
		timeoutStr := os.Args[4]
		customTimeout, err := time.ParseDuration(timeoutStr)
		if err != nil {
			fmt.Println("Invalid timeout format:", err)
			os.Exit(1)
		}
		timeout = customTimeout
	}

	// Установка соединения с сервером с учетом таймаута
	conn, err := net.DialTimeout("tcp", host+":"+port, timeout)
	if err != nil {
		if opErr, ok := err.(*net.OpError); ok && opErr.Timeout() {
			fmt.Println("Connection timeout:", err)
		} else {
			fmt.Println("Error connecting:", err)
		}
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connected to", host+":"+port)

	// Обработка сигнала завершения (Ctrl+D)
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input := scanner.Text()
			fmt.Fprintln(conn, input)
		}

		if scanner.Err() != nil {
			fmt.Println("Error reading from stdin:", scanner.Err())
		}

		fmt.Println("Closing connection...")
		conn.Close()
		os.Exit(0)
	}()

	// Отправка ввода пользователя в сокет
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input := scanner.Text()
			fmt.Fprintln(conn, input)
		}
	}()

	// Вывод данных, полученных из сокета, в STDOUT
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from server:", err)
	}
}
