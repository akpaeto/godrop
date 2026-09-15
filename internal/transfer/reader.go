package transfer

import (
	"fmt"
	"io"
	"os"
)

const ChunkSize = 32 * 1024

func ReadFile(path string) error {
	file, err := os.Open(path)

	if err != nil {
		return err
	}
	defer file.Close()
	buffer := make([]byte, ChunkSize)

	var total int64

	for {
		n, err := file.Read(buffer)

		if n > 0 {
			total += int64(n)
			fmt.Println("прочитано байт:", total)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}

	fmt.Println("файл прочитан полностью")
	return nil
}
