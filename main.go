package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Welcome to My Dummy Website</h1><p>This is a test website served by Go! wonder if this actually updated, also pipeline to build achkhsually wroks :3</p><h1>it is finished</h1><p>you cant see this if the fail safe works unless it dont. Just updated</p><h5>if you dont see this, meaningit goes to rescue mode :), testing update.</h5>")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server started at http://localhost:8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}
