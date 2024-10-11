package tcp

import (
    "net"
)

type Client struct {
    Host string
    Port string
    Conn *net.TCPConn
}

func NewClient(host string, port string) *Client {
    return &Client{
        Host: host,
        Port: port,
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

func (c *Client) Close() error {
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

func (c *Client) Send(message []byte) error {
    _, err := c.Conn.Write(message)
    if err != nil {
        return err
    }
    return nil
}
