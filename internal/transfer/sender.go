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

	_, err = io.CopyN(conn, file, info.Size())
	if err != nil {
		return err
	}

	fmt.Println("файл отправлен полностью")

	return nil

}
