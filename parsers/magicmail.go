/*
Copyright © 2025 maxvolume <ben@schonbeck.io>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package parsers

import (
	"bufio"
	"bytes"
	"fmt"
	"mime"
	"net/mail"
	"os"
)

func ParseMbox(filename string, peekNumber int) {
	fmt.Println("Hi there from benmimer...")
	file, err := os.Open(filename)
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
	for count < peekNumber {
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
                mailTo := msg.Header.Get("To")
                msgSubject := msg.Header.Get("Subject")
				fmt.Printf("From: %s\tTo: %s\tSubject: %s\n", from, mailTo, msgSubject)
                extractText(msg)
			}
			msgbuffer = msgbuffer[:0]
            // NOTE: maybe add this behind a verbose flag fmt.Println("Capacity of msgbuffer: ", cap(msgbuffer))
			msgbuffer = append(msgbuffer, data...)
			count++
		} else {
			msgbuffer = append(msgbuffer, data...)
		}
	}
}

func extractText (msg *mail.Message) {
    fmt.Printf("Content Type: %s\n\n", msg.Header.Get("Content-Type")) 
}
