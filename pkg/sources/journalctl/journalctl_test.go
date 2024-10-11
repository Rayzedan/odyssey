package journalctl_test

import (
    "testing"

    "github.com/rayzedan/odyssey/pkg/sources/journalctl"
)

func TestReadUnitWithInvalidArgs(t *testing.T) {
    _, err := journalctl.ReadUnit("", "", 5)
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
