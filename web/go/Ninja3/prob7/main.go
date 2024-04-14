// Prob7

// Add code to WRITE to the connection.

// And

// Building upon the code from the previous problem:
// Before we WRITE our RESPONSE, let's WRITE to our RESPONSE the STATUS LINE and some RESPONSE HEADERS. Remember the request line and status line:

// REQUEST LINE GET / HTTP/1.1 method SP request-target SP HTTP-version CRLF https://tools.ietf.org/html/rfc7230#section-3.1.1

// RESPONSE (STATUS) LINE HTTP/1.1 302 Found HTTP-version SP status-code SP reason-phrase CRLF https://tools.ietf.org/html/rfc7230#section-3.1.2

// Write the following strings to the response - use io.WriteString for all of the following except the second and third:

// "HTTP/1.1 200 OK\r\n"

// fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))

// fmt.Fprint(c, "Content-Type: text/plain\r\n")

// "\r\n"

// Look in your browser "developer tools" under the network tab. Compare the RESPONSE HEADERS from the previous file with the RESPONSE HEADERS in your new solution.


// Prob8 is Also in here ->

// Building upon the code from the previous problem:
// Change your RESPONSE HEADER "content-type" from "text/plain" to "text/html"

// Change the RESPONSE from "CHECK OUT THE RESPONSE BODY PAYLOAD" (and everything else it contained: request method, request URI) to an HTML PAGE that prints "HOLY COW THIS IS LOW LEVEL" in

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main(){


	li, err := net.Listen("tcp", ":8080")
	if err != nil {

		log.Fatalln("Error =",err)
	}

	defer li.Close()

	for {


		conn, err := li.Accept()
		if err != nil {

			log.Fatalln("Error =",err)
		}


		go request(conn)

		go response(conn)
	}

}


func request (conn net.Conn) {

	scanner := bufio.NewScanner(conn)

	i := 0
	
	for scanner.Scan() {
		ln := scanner.Text()

		fmt.Println(ln)

		if ln == "" {
			break
		}

		if i == 0 {

			text := strings.Fields(ln)

			fmt.Println("Method:", text[0])

			fmt.Println("URL:", text[1])
		}

		i++
	}

	fmt.Println(" I see you are finally here ")
	
	// defer conn.Close()


}

func response(conn net.Conn) {

	body := `
<!DOCTYPE HTML>
<html>
<head>
<meta charset="utf-8">
<meta name="viewport" content="intial-scale=device-width, intial-scale=1.0" />
</head>
<body>
Haha you got me
</body>
</html>
`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
	fmt.Fprintln(conn, "Hahah you see me now -> 😀 ")
	
}
