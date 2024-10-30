package main

import (
	"fmt"
	"net"
)

func main() {
	// Подключаемся к серверу на порту 8080
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Ошибка подключения к серверу:", err)
		return
	}
	defer conn.Close()

	// Отправляем сообщение серверу
	message := "Привет от клиента!"
	conn.Write([]byte(message))
	fmt.Println("Сообщение отправлено серверу")

	// Читаем ответ от сервера
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа от сервера:", err)
		return
	}
	fmt.Printf("Ответ от сервера: %s\n", string(buffer[:n]))
}
