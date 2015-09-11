package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", Route)
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}
