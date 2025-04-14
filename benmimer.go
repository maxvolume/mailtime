package main

import (
    "fmt"
    "os"
    "bufio"
    "bytes"
    "net/mail"
    "mime"
)

func main() {
    fmt.Println("Hi there from benmimer...")
    file, err := os.Open("/home/ben/projects/gostuff/Takeout/Mail/benmail.mbox")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    decoder := new(mime.WordDecoder)
    msgbuffer := make([]byte, 200000)
    prefix := []byte("From ")
    count := 0
    for count < 2 {
        data, err := reader.ReadBytes('\n')
        if err != nil {
            fmt.Println(err)
            return
        }
        if bytes.HasPrefix(data, prefix) {
            if count != 0 {
                msg, err := mail.ReadMessage(bytes.NewReader(msgbuffer))
                if err != nil {
                    fmt.Println("Error parsing data msg buffer", err)
                    break
                }
                from, _ := decoder.DecodeHeader(msg.Header.Get("From"))
                fmt.Printf("From: %s\tTo: %s\tSubject: %s\n", from, msg.Header.Get("To"), msg.Header.Get("Subject"))
            }
            msgbuffer = msgbuffer[:0]
            fmt.Println("Capacity of msgbuffer: ", cap(msgbuffer))
            msgbuffer = append(msgbuffer, data...)
            count++
        } else {
            msgbuffer = append(msgbuffer, data...)
        }
    }
}
