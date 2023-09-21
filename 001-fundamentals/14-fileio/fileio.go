package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	filename := "index.txt"
	data := "This is the content to be written into file."
	err := createfile(filename)
	if err != nil {
		log.Fatalln("CreateFile", err)
	}
	err = writetofile(filename, data)
	if err != nil {
		log.Fatalln("WriteToFile", err)
	}
	msg, err := readfromfile(filename)
	if err != nil {
		log.Fatalln("ReadFromFile", err)
	}
	fmt.Println(msg)

	err = bufferedWriteToFile("newfile.txt", "This is new content")
	if err != nil {
		log.Fatalln("bufferedWriteToFile", err)
	}
	fmt.Println(bufferedReaderFromFile("newfile.txt"))

}

func createfile(filename string) error {
	file, err := os.Create(filename)

	if err != nil {
		return err
	}
	file.Close()
	return nil
}

func writetofile(filename string, data string) error {
	f, err := os.OpenFile(filename, os.O_APPEND, 0644)
	_, err = f.Write([]byte(data))
	f.Close()
	return err
}

func readfromfile(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	fi, err := f.Stat()
	if err != nil {
		return "", err
	}
	data := make([]byte, fi.Size())
	_, err = io.ReadFull(f, data)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
