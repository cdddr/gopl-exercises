// Fetch prints the content found at a URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	outfile, err := os.Create("ex1_9_out.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "opening ex1_9_out.txt for output.")
	}
	for _, url := range os.Args[1:] {
		go fetch(url, outfile, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	outfile.Close()
}

func fetch(url string, outfile *os.File, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(outfile, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

/*
	When a website is unresponsive, the return hangs, but the other requests are able to complete.
$ ./fetchall https://www.google.com https://youtube.com https://facebook.com https://badu.com https://wikipedia.org https://yahoo.com https://twitter.com https://amazon.com https://sohu.com
0.30s   11010 https://www.google.com
0.52s  303352 https://twitter.com
0.62s  102882 https://facebook.com
0.82s  200112 https://amazon.com
0.85s   81235 https://wikipedia.org
1.02s  467903 https://yahoo.com
1.14s     102 https://sohu.com
1.88s  558955 https://youtube.com
Get https://badu.com: dial tcp 103.51.144.81:443: i/o timeout
30.00s elapsed
*/
