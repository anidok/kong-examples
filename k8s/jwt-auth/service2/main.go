package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Handle root path for each API
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        switch r.URL.Path {
        case "/":
            fmt.Fprintf(w, "Hello from Service 2!")
        case "/api1":
            fmt.Fprintf(w, "Hello from Service 2 - API 1!")
        case "/api2":
            fmt.Fprintf(w, "Hello from Service 2 - API 2!")
        default:
            http.NotFound(w, r)
        }
    })
    
    fmt.Println("Service 2 starting on port 8082...")
    http.ListenAndServe(":8082", nil)
}
