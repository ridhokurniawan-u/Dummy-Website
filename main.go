package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Welcome to My Dummy Website</h1><p>This is a test website served by Go! wonder if this actually updated, also pipeline to build achkhsually wroks :3</p><h1>it is finished</h1>")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server started at http://localhost:8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}
