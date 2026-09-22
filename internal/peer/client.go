package peer

import (
	"bufio"
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

	data = append(data, '\n')

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

	err = transfer.SendFile(conn, "test.txt")

	if err != nil {
		fmt.Println("ошибка отправки файла:", err)
		return
	}

	buffer = make([]byte, 1024)

	n, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("не удалось получить подтверждение:", err)
		return
	}

	var ack protocol.Message

	err = json.Unmarshal(buffer[:n], &ack)
	if err != nil {
		fmt.Println("не удалось разобрать АСК:", err)
		return
	}

	fmt.Println("сервер подтвердил получение:")
	fmt.Println("тип:", ack.Type)
	fmt.Println("текст:", ack.Text)

}

func SendFileToPeer(address string, filePath string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	defer conn.Close()

	fmt.Println("подлкючились к peer:", conn.RemoteAddr())
	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	message := protocol.Message{
		Type:       "hello",
		DeviceName: hostname,
	}

	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	data = append(data, '\n')

	_, err = conn.Write(data)
	if err != nil {
		return err
	}

	fmt.Println("отправили:", string(data))

	reader := bufio.NewReader(conn)

	line, err := reader.ReadBytes('\n')
	if err != nil {
		return err
	}

	var response protocol.Message

	err = json.Unmarshal(line, &response)
	if err != nil {
		return err
	}

	fmt.Println("ответ сервера:")
	fmt.Println("тип:", response.Type)
	fmt.Println("текст:", response.Text)

	err = transfer.SendFile(conn, filePath)
	if err != nil {
		return err
	}

	line, err = reader.ReadBytes('\n')
	if err != nil {
		return err
	}

	var ack protocol.Message

	err = json.Unmarshal(line, &ack)
	if err != nil {
		return err
	}

	fmt.Println("сервер подтвердил получение:")
	fmt.Println("тип:", ack.Type)
	fmt.Println("текст:", ack.Text)

	return nil

}
