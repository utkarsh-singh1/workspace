package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)


func main() {

	li, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Panic("Error =",err)
	}

	for {

		conn, err := li.Accept()
		if err != nil {
			log.Println("Error =",err)
		}

		go handle(conn)
	}
	
}

func handle(conn net.Conn) {

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		li := strings.ToLower(scanner.Text())
		bs := []byte(li)
		r := rot13(bs)
		fmt.Fprintf(conn, "%s- %s\n",li,r)
	}
	
}

func rot13(bs []byte) []byte {

	r13 := make([]byte ,len(bs) )

	for k , v := range bs {
		if v <= 109 {
			r13[k] = v + 13
		}else {
			r13[k] = v - 13
		}
	}

	return r13
}
