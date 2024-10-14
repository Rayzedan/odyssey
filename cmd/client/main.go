package main

import (
    "log"
    "os"
    "os/signal"
    "time"

    "github.com/rayzedan/odyssey/pkg/network/tcp/client"
    "github.com/rayzedan/odyssey/pkg/log_sources/journalctl"
)

func main() {
    client := tcp.NewClient("127.0.0.1", "8080")
    err := client.Connect()
    if err != nil {
        log.Fatalf("failed to connect: %v", err)
    }
    signals := make(chan os.Signal, 1)
    signal.Notify(signals, os.Interrupt)
    go func() {
        <-signals
        client.Disconnect()
        log.Println("closing connection by interrupt")
        os.Exit(0)
    } ()
    defer client.Disconnect()
    ticker := time.NewTicker(1 * time.Second)
    var cursor string
    for {
        for range ticker.C {
            event, err := journalctl.ReadUnit("", cursor, 1)
            if err != nil {
                log.Fatalf("failed to read unit: %v", err)
            }
            if len(event) == 0 || event[0].IsEmpty() {
                continue
            }
            if cursor != event[0].Cursor {
                cursor = event[0].Cursor
            }
            err = client.Send([]byte(event[0].String()))
            if err != nil {
                log.Fatalf("failed to send message: %v", err)
            }
        }
    }
}
