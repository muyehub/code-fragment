package file

import (
	"log"
	"os"
)

// CreateFileByContent 新建文件
func CreateFileByContent(content []byte, fileName string) error {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	var written int
	written, err = file.Write(content)
	if err != nil {
		return err
	}
	log.Printf("Wrote %d bytes.\n", written)
	return nil
}