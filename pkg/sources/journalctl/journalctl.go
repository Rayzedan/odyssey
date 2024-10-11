package journalctl

import (
    "encoding/json"
    "errors"
    "fmt"
    "os/exec"
    "strings"
    "time"
)

func ReadUnit(name string, cursor string, count int) ([]Event, error) {
    if count == 0 {
        return []Event{}, errors.New("invalid arguments" )
    }
    arg := fmt.Sprintf("journalctl -o json -n %d", count)
    if name != "" {
        arg = fmt.Sprintf("%s --unit %s", arg, name)
    }
    if cursor != "" {
        arg = fmt.Sprintf("%s --after-cursor=\"%s\"", arg, cursor)
    }
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()

    buf, err := exec.Command("bash", "-c", arg).Output()
    if err != nil {
        return []Event{}, err
    }
    var events []Event
    for _, item := range strings.Split(string(buf), "\n") {
        if item == "" {
            continue
        }
        var out Event
        if err := json.Unmarshal([]byte(item), &out); err != nil {
            return []Event{}, err
        }
        out.Timestamp *= int64(time.Microsecond)
        events = append(events, out)
    }
    return events, nil
}
