// Prob 8

// Building upon the code from the previous problem:
// Add code to respond to the following METHODS & ROUTES: GET / GET /apply POST /apply

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {

	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Error =", err)
	}

	defer li.Close()
	
	for {

		conn , err := li.Accept()
		if err != nil {

			log.Fatalln("Error =",err)
		}

		go request(conn)
		
	}
}


func request(conn net.Conn) {

	scanner := bufio.NewScanner(conn)
	var method string
	var url string
	i := 0
	
	for scanner.Scan() {

		ln := scanner.Text()
		str := strings.Fields(ln)
		
		fmt.Println(ln)

		if ln == "" {
			break
		}

		if i == 0 {
			method = str[0]
			url = str[1]
			fmt.Println("Method ->" ,method )
			fmt.Println("URL ->", url )
		}

		i++
		
 	}

	fmt.Println(method, url)

	switch {

	case method == "GET" && url == "/":
		Index(conn)

	case method == "GET" && url == "/welcome" :
		Welcome(conn)
	}
	
	defer conn.Close()
}

func Index(conn net.Conn) {

		body := `
<!DOCTYPE HTML>
<html>
<head>
<meta charset="utf-8">
<meta name="viewport" content="intial-scale=device-width, intial-scale=1.0" />
</head>
<body>
<h1>Index Page</h1>
<ul>
<li> <a href="/"> Home </a> </li>
<li> <a href="/welcome"> Welcome </a> </li>
</ul>
</body>
</html>
`
		io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
		fmt.Fprint(conn, "Content-Type: text/html\r\n")
		io.WriteString(conn, "\r\n")
		io.WriteString(conn, body)
	
}

func Welcome(conn net.Conn) {

		body := `
<!DOCTYPE HTML>
<html>
<head>
<meta charset="utf-8">
<meta name="viewport" content="intial-scale=device-width, intial-scale=1.0" />
</head>
<body>
<h1>Welcome Page</h1>
<ul>
<li> <a href="/"> Home </a> </li>
<li> <a href="/welcome"> Welcome </a> </li>
</ul>
</body>
</html>
`
		io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
		fmt.Fprint(conn, "Content-Type: text/html\r\n")
		io.WriteString(conn, "\r\n")
		io.WriteString(conn, body)
	
}


