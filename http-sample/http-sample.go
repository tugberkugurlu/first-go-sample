package main

import "net/http"
import "fmt"

func main() {
	resp, err := http.Get("http://example.com/")
	fmt.Println(resp)
	fmt.Println(err)
}