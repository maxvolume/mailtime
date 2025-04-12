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
    for range 4 {
        data, err := reader.ReadBytes('\n')
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println(string(data))
    }
}
