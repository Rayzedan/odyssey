package common

type Packet struct {
    Header []byte
    Body   []byte
}

func NewPacket(header []byte, body []byte) *Packet {
    return &Packet{
        Header: header,
        Body:   body,
    }
}

func (p *Packet) Len() int {
    return len(p.Header) + len(p.Body)
}

func (p *Packet) String() string {
    return string(p.Header) + string(p.Body)
}

type State int

const (
    StateHello = iota
    StateConnected
)
