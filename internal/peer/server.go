package peer

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/akpaeto/godrop/internal/protocol"
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

		buffer := make([]byte, 1024)

		n, err := conn.Read(buffer)

		if err != nil {
			fmt.Println("ошибка чтения")
			continue
		}

		var message protocol.Message

		err = json.Unmarshal(buffer[:n], &message)

		if err != nil {
			fmt.Println("ошибка чтения json", err)
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
			fmt.Println("ошибка создания json", err)
			conn.Close()
			continue
		}

		_, err = conn.Write(data)

		if err != nil {
			fmt.Println("ошибка отправки", err)
		}

		fmt.Println("отправлен ответ:", string(data))
		conn.Close()

	}

}
