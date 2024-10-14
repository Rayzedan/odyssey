package journalctl

import (
    "fmt"
    "time"
)

type Event struct {
    Cursor        string      `json:"__CURSOR"`
    Timestamp     int64       `json:"__REALTIME_TIMESTAMP,string"`
    Message       string      `json:"-"`
    MessageSource string      `json:"MESSAGE"`
    Unit          string      `json:"_SYSTEMD_UNIT"`
    Priority      string      `json:"PRIORITY"`
    UnitResult    string      `json:"UNIT_RESULT"`
    SourceData    string      `json:"-"`
}

func (e *Event) IsEmpty() bool {
    return e.Cursor == ""
}

func (e Event) String() string {
    return fmt.Sprintf("%v %s %s %s", time.Unix(0, e.Timestamp), e.Priority, e.Unit, e.MessageSource)
}
