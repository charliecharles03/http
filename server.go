package main

import (
	"fmt"
	handler "http/lib/utils"
	"net"
)

func main(){
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("error maa")
    }
    defer ln.Close()
    for {
        conn, err := ln.Accept()
       if err != nil {
            fmt.Println("error maa")
            return 
        }
        fmt.Println(conn)
        handler.HandleConnection(conn) // getting request content
        if err != nil {
            panic("some thing wrong baby")
        }
    }
}

