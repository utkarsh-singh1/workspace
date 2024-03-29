package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic("Error =",err)
	}
	defer li.Close()
	
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println("Error =",err)
		}
		go handle(conn)
	}

}

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
}

func request(conn net.Conn)  {

	i := 0
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			mux(conn, ln)			
		}
		if ln == "" {
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {

	m := strings.Fields(ln)[0]
	u := strings.Fields(ln)[1]

	if m == "GET" && u =="/" {
		index(conn)		
	}

	if m == "GET" && u =="/about" {
		about(conn)
	}

	if m == "GET" && u == "/contact" {
		contact(conn)
	}

	
	
}

func index(conn net.Conn) {


			body := `<!DOCTYPE html>
<html lang="en-us">
<head>
<meta charset="utf-8">
<title></title>
</head>
<body>
<STRONG>Index</STRONG>
<a href="/">index</a><br>
<a href="/about">about</a><br>
<a href="/contact">contact</a><br>
<a href="/apply">apply</a><br>
</body>
</html>
`
		fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
		fmt.Fprint(conn, "Content-Type: text/html\r\n")
		fmt.Fprint(conn, "\r\n")
		fmt.Fprint(conn, body)
	
}


func about(conn net.Conn) {

			body := `<!DOCTYPE html>
<html lang="en-us">
<head>
<meta charset="utf-8">
<title></title>
</head>
<body>
<a href="/">index</a><br>
<a href="/about">about</a><br>
<a href="/contact">contact</a><br>
</body>
</html>
`
		fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
		fmt.Fprint(conn, "Content-Type: text/html\r\n")
		fmt.Fprint(conn, "\r\n")
		fmt.Fprint(conn, body)

}

func contact(conn net.Conn){

	body := `<!DOCTYPE html>
<html lang="en-us">
<head>
<meta charset="utf-8">
<title></title>
</head>
<body>
<a href="/">index</a><br>
<a href="/about">about</a><br>
<a href="/contact">contact</a><br>
</body>
</html>
`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)

}

