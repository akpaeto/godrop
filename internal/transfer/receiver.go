package transfer

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReceiveFile(reader *bufio.Reader, fileName string, fileSize int64) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	buffer := make([]byte, ChunkSize)
	var received int64

	for received < fileSize {
		remaining := fileSize - received

		if int64(len(buffer)) > remaining {
			buffer = buffer[:remaining]
		}

		n, err := reader.Read(buffer)

		if n > 0 {
			_, writeErr := file.Write(buffer[:n])
			if writeErr != nil {
				return writeErr
			}

			received += int64(n)
			fmt.Println("получено:", received, "/", fileSize, "байт")
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}
	}

	if received != fileSize {
		return fmt.Errorf(
			"получено %d байт, ожидалось %d",
			received,
			fileSize,
		)
	}
	fmt.Println("файл получен:", fileName)
	fmt.Println("размер:", received, "байт")

	return nil

}
