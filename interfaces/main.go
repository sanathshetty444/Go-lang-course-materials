package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type bot interface {
	greeting() string
}
type englishbot struct{}
type spanishbot struct{}

type todoResponse struct {
	userId    int    `json:"userid"`
	id        int    `json:"id"`
	title     string `json:"title"`
	completed bool   `json:"completed"`
}

type logWriter struct{}

func main() {
	eb := englishbot{}
	sb := spanishbot{}
	printGreeting(eb)
	printGreeting(sb)
	requestToGoogle()
}

func printGreeting(b bot) {
	fmt.Println(b.greeting())
}
func (eb englishbot) greeting() string {
	return "Hi there"
}
func (sb spanishbot) greeting() string {
	return "Hola"
}
func requestToGoogle() {
	res, err := http.Get("http://google.com")
	// res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// b := make([]byte, 99999)
	// res.Body.Read(b)
	// s := string(b)
	// fmt.Print(s)
	// io.Copy(os.Stdout, res.Body) //Here stdout implement Writer interface

	l_w := logWriter{}
	io.Copy(l_w, res.Body) //Here l_w has Write function which is a prerequisite of Writer interface
}
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("No of bytes:", len(bs))
	return len(bs), nil
}
