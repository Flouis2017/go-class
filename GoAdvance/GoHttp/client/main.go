package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	client := new(http.Client)
	req, _ :=http.NewRequest("POST", "http://localhost:9090/test", bytes.NewBuffer([]byte("456")))
	resp, _ := client.Do(req)
	body := resp.Body
	b, _ := io.ReadAll(body)
	fmt.Println(string(b))
}
