package handler

import (
	"bufio"
	"fmt"
	"net"
)

func HandleConnection(conn net.Conn) {
    defer conn.Close()
    result := bufio.NewReader(conn)
    header,err := result.ReadString('\n')
    if err != nil {
        return
    }
    fmt.Println(header)
    conn.Write([]byte("GG sir"))
}

