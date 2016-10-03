package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error occured while getting the url : %v\n", url)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Printf("Error occured while trying to write the url : %v to os.Stdout\n", url)
		}
		resp.Body.Close()
	}
}
