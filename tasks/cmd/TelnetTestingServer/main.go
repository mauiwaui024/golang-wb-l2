package main

import (
	"fmt"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Бесконечный цикл для чтения данных из соединения
	for {
		// Буфер для чтения данных из сокета
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}

		// Вывод полученных данных в STDOUT
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	// Порт, на котором будет запущен сервер
	port := "8080"

	// Запуск сервера
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Server listening on port", port)

	// Бесконечный цикл для принятия и обработки соединений
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			os.Exit(1)
		}
		fmt.Println("Connection accepted from", conn.RemoteAddr())

		// Обработка соединения в отдельной горутине
		go handleConnection(conn)
	}
}
