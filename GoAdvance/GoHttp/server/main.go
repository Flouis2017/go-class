package main

import (
	"fmt"
	"io"
	"net/http"
)

func testHandler(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
		case "GET":
			resp.Write([]byte("This is Get Response"))
			break
		case "POST":
			var bytes, _ = io.ReadAll(req.Body)
			fmt.Println(string(bytes))
			resp.Write([]byte("This is Post Response"))
			break
	}
}

func main() {
	//var mux = http.NewServeMux()
	http.HandleFunc("/test", testHandler)
	http.ListenAndServe(":9090", nil)
}
