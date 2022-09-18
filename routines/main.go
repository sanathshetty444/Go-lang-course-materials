package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://golang.org",
		"http://amazon.com",
		"http://aerem.co",
	}

	c := make(chan string)
	f, _ := os.OpenFile("sitecheck.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	for _, link := range links {
		go checkLink(link, c, f)
	}

	// for {
	// 	l := <-c
	// 	fmt.Println(l)
	// 	f.Write([]byte(l + "\n"))
	// 	go checkLink(strings.Split(l, " ")[0], c)
	// }

	for l := range c {
		fmt.Println(l)
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c, f)
		}(l)
	}
}

func checkLink(link string, c chan string, f *os.File) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down")
		f.Write([]byte(link + " is down \n"))

	}
	f.Write([]byte(link + " is up \n"))
	c <- link

}
