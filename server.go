package main
import (
    "fmt";
    "net";
    "bufio"
)

func main(){
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("error maa")
    }
    for {
        conn, err := ln.Accept()
       if err != nil {
            fmt.Println("error maa")
        }
        fmt.Println(conn)
        result := bufio.NewReader(conn)
        val,err := result.ReadBytes('\n')
        if err != nil {
            panic("some thing wrong baby")
        }
        fmt.Println(string(val))
    }
}

