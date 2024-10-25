package main

import (
    "fmt"
    "log"
    "net/http"
)

// Kezelő függvény a főoldalhoz
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    // Beállítjuk a route-ot, hogy a "/" URL-t kezelje a helloHandler függvény
    http.HandleFunc("/", helloHandler)

    // Indítjuk a web szervert a 8080-as porton
    fmt.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
