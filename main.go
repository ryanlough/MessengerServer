package main

import (
    "log"
    "net/http"
    "fmt"
)

func main() {
    router := NewRouter()

    fmt.Printf("Connecting to port :8080...")

    log.Fatal(http.ListenAndServe(":8080", router))
}
