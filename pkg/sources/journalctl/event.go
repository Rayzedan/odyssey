package journalctl

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
