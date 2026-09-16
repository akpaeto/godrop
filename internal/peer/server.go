package peer

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"

	"github.com/akpaeto/godrop/internal/protocol"
	"github.com/akpaeto/godrop/internal/transfer"
)

func StartServer(port string) {
	listener, err := net.Listen("tcp", ":"+port)

	if err != nil {
		fmt.Println("сервер отвалился пошел нахуй", err)
		return
	}

	defer listener.Close()

	manager := PeerManager{}

	fmt.Println("GoDrop server started!")
	fmt.Println("Listening on port:", port)

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("НЕ подключился", err)
			continue
		}

		fmt.Println("New peer connected:", conn.RemoteAddr())

		reader := bufio.NewReader(conn)

		line, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Println("ошибка чтения hello:", err)
			conn.Close()
			continue
		}

		var message protocol.Message

		err = json.Unmarshal(line, &message)

		if err != nil {
			fmt.Println("ошибка чтения json", err)
			conn.Close()
			continue
		}

		fmt.Println("тип сообщения:", message.Type)
		fmt.Println("устройство:", message.DeviceName)

		newPeer := &Peer{
			Name:    message.DeviceName,
			Address: conn.RemoteAddr().String(),
		}

		manager.AddPeer(newPeer)

		foundPeer := manager.Find(newPeer.Name)
		if foundPeer != nil {
			fmt.Println("нашли Peer", foundPeer.Name)
		} else {
			fmt.Println("peer не найден")
		}

		fmt.Println()
		fmt.Println("=== PEERS ===")

		for _, peer := range manager.Peers {
			fmt.Println(peer.Name, "-", peer.Address)
		}

		fmt.Println("=============")
		fmt.Println()

		response := protocol.Message{
			Type: "welcome",
			Text: "Welcome to GoDrop",
		}

		data, err := json.Marshal(response)
		if err != nil {
			fmt.Println("ошибка создания json:", err)
			conn.Close()
			continue
		}

		data = append(data, '\n')

		_, err = conn.Write(data)
		if err != nil {
			fmt.Println("ошибка отправки:", err)
			conn.Close()
			continue

		}

		fmt.Println("отправлен ответ:", string(data))

		fileLine, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Println("ошибка чтения file_info:", err)
			conn.Close()
			continue
		}

		var fileInfo protocol.Message

		err = json.Unmarshal(fileLine, &fileInfo)
		if err != nil {
			fmt.Println("ошибка чтения file_info:", err)
			conn.Close()
			continue
		}

		fmt.Println()
		fmt.Println("=== FILE INFO ===")
		fmt.Println("имя:", fileInfo.FileName)
		fmt.Println("размер:", fileInfo.FileSize, "байт")
		fmt.Println("=================")

		err = transfer.ReceiveFile(reader, fileInfo.FileName, fileInfo.FileSize)
		if err != nil {
			fmt.Println("ошибка получения файла:", err)
			conn.Close()
			continue
		}

		ack := protocol.Message{
			Type: "file_received",
			Text: "File received successfully",
		}

		ackData, err := json.Marshal(ack)
		if err != nil {
			fmt.Println("ошибка создания ACK:", err)
			conn.Close()
			continue
		}

		ackData = append(ackData, '\n')

		_, err = conn.Write(ackData)

		if err != nil {
			fmt.Println("ошибка отправки АСK:", err)
			conn.Close()
			continue
		}

		fmt.Println("отправлено подтверждение:", string(ackData))

		conn.Close()

	}

}
