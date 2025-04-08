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
            fmt.Fprintf(w, "Hello from Service 1!")
        case "/api1":
            fmt.Fprintf(w, "Hello from Service 1 - API 1!")
        case "/api2":
            fmt.Fprintf(w, "Hello from Service 1 - API 2!")
        default:
            http.NotFound(w, r)
        }
    })
    
    fmt.Println("Service 1 starting on port 8081...")
    http.ListenAndServe(":8081", nil)
}
