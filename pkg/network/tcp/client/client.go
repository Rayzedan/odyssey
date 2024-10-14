package tcp

import (
    "net"

    "github.com/rayzedan/odyssey/pkg/network/tcp/common"
)

type Client struct {
    ID   int
    Host string
    Port string
    State common.State
    Conn *net.TCPConn
}

func NewClient(host string, port string) *Client {
    return &Client{
        ID:   0,
        Host: host,
        Port: port,
        State: common.StateHello,
        Conn: nil,
    }
}

func (c *Client) Connect() error {
    conn, err := new(c.Host, c.Port)
    if err != nil {
        return err
    }
    c.Conn = conn
    return  nil
}

func (c *Client) Disconnect() error {
    return c.Conn.Close()
}

func new(host string, port string) (*net.TCPConn, error) {
    tcpServer, err := net.ResolveTCPAddr("tcp", host+":"+port)
    if err != nil {
        return nil, err
    }
    conn, err := net.DialTCP("tcp", nil, tcpServer)
    if err != nil {
        return nil, err
    }

    return conn, nil
}

func (c *Client) doHello() error {
    _, err := c.Conn.Write([]byte("Hello"))
    if err != nil {
        return err
    }
    return nil
}

func (c *Client) Send(message []byte) error {
    switch c.State {
    case common.StateHello:
        err := c.doHello()
        if err != nil {
            return err
        }
        c.State = common.StateConnected
        break
    case common.StateConnected:
        _, err := c.Conn.Write(message)
        if err != nil {
            return err
        }
        break
    }
    return nil
}
