package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	ntpTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		// Вывод ошибки в STDERR
		fmt.Fprintln(os.Stderr, "Ошибка получения времени с NTP сервера:", err)
		// Возвращение ненулевого кода выхода в OS
		os.Exit(1)
	}

	// Форматирование времени для вывода
	currentTime := time.Now()
	fmt.Println("Текущее локальное время:", currentTime.Format("2007-01-02 15:04:05"))
	fmt.Println("Точное время с NTP сервера:", ntpTime.Format("2007-01-02 15:04:05"))
}
