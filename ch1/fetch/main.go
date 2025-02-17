package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	// fetch1()
	fetchAll()
}

func fetch1() {
	for _, url := range os.Args[1:] {
		url_ := ""
		if !strings.HasPrefix(url, "http://") {
			url_ = "http://" + url
		}
		fmt.Println(url_)
		resp, err := http.Get(url_)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		body, err := io.Copy(os.Stdout, resp.Body) //(resp.Body)
		status := resp.Status
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: чтение %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%d\t%s", body, status)
	}
}

func fetchAll() {
	fo, err := os.Create("output.txt")
	if err != nil {
		fmt.Printf("fetchAll: %v", err)
		return
	}
	defer fo.Close()

	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetchUrl(url, ch)
	}
	for range os.Args[1:] {
		result := <-ch // получение из канала ch
		fo.Write([]byte(result + "\n"))
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetchUrl(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close() // исключение утечки ресурсов
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d Bytes %s", secs, nbytes, url)
}
