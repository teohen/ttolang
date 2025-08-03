package utils

import (
	"bufio"
	"log"
	"os"
)

func LoadFile(path string) (error, string) {
	file, err := os.Open(path)

	if err != nil {
		return err, ""
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	code := ""

	for scanner.Scan() {
		code += scanner.Text()
	}

	return nil, code
}

func WriteFile(filePath string, text string) error {
	data := []byte(text)

	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
