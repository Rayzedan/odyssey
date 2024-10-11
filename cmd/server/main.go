package main

import (
    "log"

    "github.com/rayzedan/odyssey/pkg/network/tcp/server"
)

func main() {
    server := tcp.NewServer("127.0.0.1", "8080")
    err := server.Listen()
    if err != nil {
        log.Fatal(err)
    }
}
