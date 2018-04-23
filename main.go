package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"bytes"
	"os"
	"bufio"
)

func goCounter(url string, buf chan int) {

	var gocount int = 0

	req, err := http.Get(url)
	if err != nil {
	}
	defer req.Body.Close()

	date, err := ioutil.ReadAll(req.Body)
	if err != nil {
	}
	gocount = bytes.Count(date, []byte("Go"))

	fmt.Printf("Count for %s:\t%d\n", url, gocount)

	buf<-gocount
}


func main() {

	entrance := os.Stdin
	scanner := bufio.NewScanner(entrance)

	buf := make(chan int, 5)
	total := 0

	for scanner.Scan() {
		go goCounter(scanner.Text(), buf)
		total += <-buf
	}

	fmt.Printf("Total: %d ", total)
}
