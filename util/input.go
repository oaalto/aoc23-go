package util

import (
	"log"
	"os"
)

func OpenFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Fatal(err)
	}
}
