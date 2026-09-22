package main

import (
	"fmt"
	"os"

	"github.com/akpaeto/godrop/internal/peer"
	"github.com/akpaeto/godrop/internal/transfer"
)

//go run ./cmd/godrop server  команда для запуска сервера экей первого терминала
//go run ./cmd/godrop connect localhost:8080 это для запуска клиента экей второго терминала

func main() {
	fmt.Println("================================")
	fmt.Println("         GoDrop 🚀")
	fmt.Println("   P2P File Sharing Platform")
	fmt.Println("================================")

	hostname, err := os.Hostname()

	if err != nil {
		fmt.Println("иди нахуй уебок")
		return
	}

	fmt.Println()
	fmt.Println("device:", hostname)
	fmt.Println()

	if len(os.Args) < 2 {
		fmt.Println("использование")
		fmt.Println(" godrop serber")
		fmt.Println(" godrop connect<address>")
		return
	}

	command := os.Args[1]

	switch command {

	case "server":
		fmt.Println("status: online🟢")
		peer.StartServer("8080")

	case "connect":
		if len(os.Args) < 3 {
			fmt.Println("укажите адрес для полюкчения")
			fmt.Println("например  godrop connect localhost:8080")
			return
		}

		address := os.Args[2]

		peer.Connect(address)

	case "read":
		err := transfer.ReadFile("test.txt")

		if err != nil {
			fmt.Println("error:", err)
		}

	case "send":
		if len(os.Args) < 4 {
			fmt.Println("использование:")
			fmt.Println(" godrop send <address> <file>")
			return
		}

		address := os.Args[2]
		filePath := os.Args[3]

		err := peer.SendFileToPeer(address, filePath)
		if err != nil {
			fmt.Println("ошибка отправки  файла:", err)
			return
		}

	default:
		fmt.Println("Неизвестная команда:", command)

	}

}
