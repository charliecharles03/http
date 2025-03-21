package main

import (
	"bufio"
	"fmt"
	"net"
    "io"
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
        RequestContent := []string{}
        RequestContent = requestBreaker(result)
        fmt.Print(RequestContent)
        if err != nil {
            panic("some thing wrong baby")
        }
    }
}

func requestBreaker(result *bufio.Reader) []string {
    RequestContent := [] string{}
    for i:=0 ;i < 100 ; i++ {
        val,err := result.ReadBytes('\n') // nothing but a ascii  don't know the cost
        fmt.Println(len(val))
        if err != nil {
            if err == io.EOF{
                if len(val) >0 {
                   RequestContent = append(RequestContent,string(val));
                }
                break
            }
        }
        currentString := string(val)
        if len(currentString) == 0 {
            fmt.Println("you meet you end so die")
            break;
        }

        RequestContent = append(RequestContent, currentString)

        if err != nil{
            panic("processed it wrong bro")
        }
    } 
    return RequestContent
}
