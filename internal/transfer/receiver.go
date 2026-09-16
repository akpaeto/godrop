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

	written, err := io.CopyN(file, reader, fileSize)
	if err != nil {
		return err
	}

	if written != fileSize {
		return fmt.Errorf(
			"получено %d байт, ожидалось %d",
			written,
			fileSize,
		)
	}

	fmt.Println("файл получен:", fileName)
	fmt.Println("размер:", written, "байт")

	return nil

}
