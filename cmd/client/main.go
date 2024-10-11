package main

import (
    "log"
    "strconv"
    "time"

    "github.com/rayzedan/odyssey/pkg/network/tcp/client"
    "github.com/rayzedan/odyssey/pkg/sources/journalctl"
)

func main() {
    client := tcp.NewClient("127.0.0.1", "8080")
    err := client.Connect()
    if err != nil {
        log.Fatalf("failed to connect: %v", err)
    }
    defer client.Close()
    ticker := time.NewTicker(1 * time.Second)
    var cursor string
    for {
        for range ticker.C {
            log.Printf("Reading unit...")
            event, err := journalctl.ReadUnit("systemd-journald.service", cursor, 1)
            if err != nil {
                log.Fatalf("failed to read unit: %v", err)
            }
            if len(event) == 0 || event[0].IsEmpty() {
                log.Printf("event is empty skip")
                continue
            }
            if cursor != event[0].Cursor {
                cursor = event[0].Cursor
                log.Printf("cursor: %s", cursor)
            }
            err = client.Send([]byte(event[0].Unit + " " + event[0].MessageSource + " " + strconv.FormatInt(event[0].Timestamp, 10)))
            if err != nil {
                log.Fatalf("failed to send message: %v", err)
            }
        }
    }
}
