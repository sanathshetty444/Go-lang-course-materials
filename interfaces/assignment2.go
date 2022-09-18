package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("No of bytes:", len(bs))
	return len(bs), nil
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error %+v", err)
		os.Exit(1)
	}
	io.Copy(logWriter{}, file)
}
