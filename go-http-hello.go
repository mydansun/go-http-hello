package main

import (
    "fmt"
    "os"
    "strconv"
    "net/http"
    "strings"
)

func hello(w http.ResponseWriter, req *http.Request) {
    ip := req.Header.Get("X-Forwarded-For")
    if ip == "" {
        ip = req.Header.Get("X-Real-IP")
    }
    if ip == "" {
        ip = req.RemoteAddr
        ip = strings.Split(ip, ":")[0]
    }
    fmt.Fprintf(w, "Hello, your IP is %s\n", ip)
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
