package main

import (
    "fmt"
    "os"
    "bufio"
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
    data := make([]byte, 500)
    n, err := reader.Read(data)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("Read %d bytes from mail data...\n", n)
    fmt.Println(string(data))
}
