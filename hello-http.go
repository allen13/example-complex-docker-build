package main

import (
    "fmt"
    "net/http"
    "os"
    "log"
)

func hello(w http.ResponseWriter, req *http.Request) {
    helloEnv := os.Getenv("HELLO")
    
    if helloEnv == "" {
        helloEnv = "hello"
    }

    arg := os.Args[1]
    if arg == "" {
        arg = "friend"
    }
    
    fmt.Fprintf(w, "%s %s\n", helloEnv, arg)
}

func main() {

    http.HandleFunc("/hello", hello)
    log.Println("Starting server on 0.0.0.0:8090")
    http.ListenAndServe(":8090", nil)
}