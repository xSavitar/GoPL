// Author: Alangi Derick
// Description: Fetchall fetches URLs in parallel and reports
//              their times and sizes

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now() // Get the time just when the program starts running.

	ch := make(chan string) // create a channel called "ch"

	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}

	fmt.Print("%.2fs elapsed\n", time.Since(start).Seconds()) // printing the total time elapsed by the longest
	// goroutine run.
}

// define the fetch() routine
func fetch(url string, ch chan<- string) {
	start := time.Now()

	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources :)

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds() // get time for each goroutine run on a URL
	ch <- fmt.Sprintf("%.2fs    %7d    %s", secs, nbytes, url)
}
