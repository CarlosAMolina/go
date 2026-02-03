package main

import (
	"log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, HTTPS World with Self-Signed Certificate!"))
    })
    certFile := "server.cert"
    keyFile := "server.key"
	port := ":8443"
    log.Println("Starting server at https://localhost" + port)
    err := http.ListenAndServeTLS(port, certFile, keyFile, nil)
    if err != nil {
        log.Fatalf("ListenAndServeTLS failed: %v", err)
    }
}
