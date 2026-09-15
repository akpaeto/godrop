package peer

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/akpaeto/godrop/internal/protocol"
	"github.com/akpaeto/godrop/internal/transfer"
)

func Connect(address string) {
	conn, err := net.Dial("tcp", address)

	if err != nil {
		fmt.Println("не удалось подключиться:", err)
		return
	}

	defer conn.Close()
	fmt.Println("подключились к peer", conn.RemoteAddr())

	hostname, err := os.Hostname()

	if err != nil {
		fmt.Println("не удалось получить имя устройства", err)
		return
	}

	message := protocol.Message{
		Type:       "hello",
		DeviceName: hostname,
	}

	data, err := json.Marshal(message)

	if err != nil {
		fmt.Println("не удалось получить json", err)
		return
	}

	_, err = conn.Write(data)

	if err != nil {
		fmt.Println("не удалось отправить сообщение", err)
		return
	}

	fmt.Println("отправили:", string(data))

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)

	if err != nil {
		fmt.Println("не удалось полуучить ответ", err)
		return
	}

	var response protocol.Message

	err = json.Unmarshal(buffer[:n], &response)

	if err != nil {
		fmt.Println("не удалось разоброать json", err)
		return
	}

	fmt.Println()
	fmt.Println("ответ сервера:")
	fmt.Println("тип:", response.Type)
	fmt.Println("текст:", response.Text)

	err = transfer.SendFile(conn, "text.txt")

	if err != nil {
		fmt.Println("ошибка отправки файла:", err)
		return
	}

}
