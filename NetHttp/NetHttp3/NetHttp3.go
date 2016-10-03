package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http") {
			url += "http://"
		}
		resp, err := http.Get(url)
		fmt.Printf("Got status code as : %v, fot url : %v", resp.StatusCode, url)
		if err != nil {
			fmt.Printf("Error: %v occured while getting the url : %v data\n", err, url)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Printf("Error: %v occured while copying data for url : %v\n", err, url)
		}
		resp.Body.Close()
	}
}
