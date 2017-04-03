package main

import (
	"fmt"
	"io"
	"os"
	"net/http"
	"time"
	"bufio"
)

func fetch(url string, contentType string, path string, c chan string, s chan bool) {
	start := time.Now()

	file, err := os.Open(path)
	defer file.Close()

	if err != nil{
		c <- fmt.Sprintf("Error : %v",err)
		s <- true
		return
	}

	r, err := http.Post(url, contentType ,bufio.NewReader(file))
	if err != nil {
		c <- fmt.Sprintf("Error : %v",err)
		s <- true
		return
	}
	nb, err := io.Copy(os.Stdout, r.Body)
	r.Body.Close()
	if err != nil {
		c <- fmt.Sprintf("Error : %v",err)
		s <- true
		return
	}
	c <- fmt.Sprintf("Start : %25v Bytes received : %8d URL : %s",start,nb,url)
	s <- true
}

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("Usage : %s <url> <content-type> <path-file>\n",os.Args[0])
		return
	}
	c := make(chan string)
	status := make(chan bool)
	url := os.Args[1]
	contentType := os.Args[2]
	path := os.Args[3]
	go fetch(url,contentType, path, c, status)

	for {
		select {
			case m:=<-c: fmt.Printf("Response : \n%s\n",m)
			case <-status: return
		}
	}
}
