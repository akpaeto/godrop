package main

import (
	"fmt"
	"os"

	"github.com/akpaeto/godrop/internal/peer"
	"github.com/akpaeto/godrop/internal/transfer"
)

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

	default:
		fmt.Println("Неизвестная команда:", command)

	}

}
