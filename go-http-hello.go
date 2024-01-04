package main

import (
    "fmt"
    "os"
    "strconv"
    "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {
    port := 8090
    args := os.Args[1:]
    if len(args) > 0 {
        argPort, err := strconv.Atoi(args[0])
        if err == nil {
            port = argPort
        }
    }
    http.HandleFunc("/", hello)
    http.HandleFunc("/headers", headers)
    address := fmt.Sprintf("0.0.0.0:%d", port)
    fmt.Println("Listening on:", address)
    http.ListenAndServe(address, nil)
}