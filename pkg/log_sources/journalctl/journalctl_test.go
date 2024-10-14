package journalctl_test

import (
    "testing"

    "github.com/rayzedan/odyssey/pkg/log_sources/journalctl"
)

func TestReadUnitWithInvalidArgs(t *testing.T) {
    _, err := journalctl.ReadUnit("", "", 0)
    if err == nil {
        t.Fatal("expected error")
    }
}

func TestReadUnitWithoutCursor(t *testing.T) {
    event, err := journalctl.ReadUnit("systemd-logind.service", "", 5)
    if err != nil {
        t.Fatal(err)
    }
    if len(event) == 0 || event[0].IsEmpty() {
        t.Fatal("event is empty")
    }
}

func TestReadUnitWithCursor(t *testing.T) {
    event, err := journalctl.ReadUnit("", "", 2)
    if err != nil {
        t.Fatal(err)
    }
    if len(event) == 0 || event[0].IsEmpty() {
        t.Fatal("event is empty")
    }
    cursor := event[0].Cursor
    event, err = journalctl.ReadUnit("", cursor, 2)
    if err != nil {
        t.Fatal(err)
    }
    if len(event) == 0 || event[0].IsEmpty() {
        t.Fatal("event is empty")
    }
}
