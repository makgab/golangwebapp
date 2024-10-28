package main

import (
    "fmt"
    "log"
    "net/http"
)

// Handler function
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    // séta the route for /
    http.HandleFunc("/", helloHandler)

    // run server
    fmt.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
