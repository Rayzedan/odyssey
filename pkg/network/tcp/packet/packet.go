package packet

type Packet struct {
    data []byte
    len  int
}

const (
    VERSION     = 1
    HEADER_SIZE = 4
)

func NewPacket(body []byte) *Packet {
}
