
package tcp

import (
    "log"
    "net"
)

type Server struct {
    Host     string
    Port     string
    Listener net.Listener
}

func NewServer(host string, port string) *Server {
    return &Server{
        Host: host,
        Port: port,
        Listener: nil,
    }
}

func (s *Server) Listen() error {
    server, err := net.Listen("tcp", s.Host+":"+s.Port)
    if err != nil {
        return err
    }
    s.Listener = server
    defer s.Listener.Close()

    log.Printf("Server listening on %s:%s\n", s.Host, s.Port)

    for {
        conn, err := s.Listener.Accept()
        if err != nil {
            log.Println("Error accepting connection:", err)
            continue
        }
        go handleRequest(conn)
    }
}

func handleRequest(conn net.Conn) {
    defer conn.Close()

    buffer := make([]byte, 1024)

    for {
        n, err := conn.Read(buffer)
        if err != nil {
            log.Println("Error reading from connection:", err)
            return
        }

        if n == 0 {
            log.Println("Connection closed by client")
            return
        }

        log.Printf("Received message: %s\n", string(buffer[:n]))
    }
}

