package main

import (
	"fmt"
	// перепиши на gin
	"net"
)

// должны появится два файла с зависимостями go.mod, go.sum

func main() {

	// добавить ручки
	// getRoutes

	//http://localhost:8080 
	// Создаем слушающий сервер на порту 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Ошибка при создании сервера:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Сервер запущен и ждет соединения...")

	for {
		// Принимаем новое соединение
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ошибка при подключении:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func getRoutes() *gin.... {
	// добавить ручки
	
	// /hello POST response {"hello": "hi"}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Новое соединение установлено")

	// Читаем сообщение от клиента
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Ошибка при чтении сообщения:", err)
		return
	}
	fmt.Printf("Сообщение от клиента: %s\n", string(buffer[:n]))

	// Отправляем ответ клиенту
	response := "Привет от сервера!"
	conn.Write([]byte(response))
}
