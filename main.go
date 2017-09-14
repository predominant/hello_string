package main

import (
    "encoding/json"
    "fmt"
    "os"
    "net/http"
)

type Config struct {
    Message string
}

func handler(w http.ResponseWriter, r *http.Request) {
    file, _ := os.Open("conf.json")
    decoder := json.NewDecoder(file)
    config := Config{}
    err := decoder.Decode(&config)
    if err != nil {
        fmt.Println("error:", err)
    }

    fmt.Fprintf(w, "Hello: %s", config.Message)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
