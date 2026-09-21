package transfer

import (
	"encoding/json"
	"fmt"
	"io"

	"net"
	"os"

	"github.com/akpaeto/godrop/internal/protocol"
)

func SendFile(conn net.Conn, path string) error {
	file, err := os.Open(path)

	if err != nil {
		return err
	}

	defer file.Close()

	info, err := file.Stat()

	if err != nil {
		return err
	}

	fileInfo := protocol.Message{
		Type:     "file_info",
		FileName: info.Name(),
		FileSize: info.Size(),
	}

	data, err := json.Marshal(fileInfo)
	if err != nil {
		return err
	}

	data = append(data, '\n')

	fmt.Println("отправляем JSON:", string(data))

	_, err = conn.Write(data)
	if err != nil {
		return err
	}

	fmt.Println("отправляем файл:", info.Name())
	fmt.Println("размер:", info.Size(), "байт")

	buffer := make([]byte, ChunkSize)
	var sent int64

	for {
		n, err := file.Read(buffer)

		if n > 0 {
			_, writeErr := conn.Write(buffer[:n])
			if writeErr != nil {
				return writeErr
			}

			sent += int64(n)

			fmt.Println("оптравлено:", sent, "/", info.Size(), "байт")
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}
	}

	return nil

}
