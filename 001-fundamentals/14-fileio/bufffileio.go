package main

import (
	"bufio"
	"io"
	"os"
)

func bufferedReaderFromFile(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	bs, err := io.ReadAll(f)

	if err != nil {
		return "", err
	}
	f.Close()
	return string(bs), nil
}

func bufferedWriteToFile(filename, data string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	bw := bufio.NewWriter(file)

	_, err = bw.WriteString(data)
	if err != nil {
		return err
	}
	bw.Flush()
	file.Close()
	return err
}
