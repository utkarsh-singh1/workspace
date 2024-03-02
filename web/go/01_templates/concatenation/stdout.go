package concatenation

import "fmt"

func Stdout() {

	name := "Myself"

	print := `
                    <!DOCTYPE html>
                    <html lang="en">
                    <head>
                        <meta name="viewporrt" content="width=device-width initial-scale=1.0" >
                        <meta charset="utf-8" >
                    </head>
                    <body>
                          <h1> Hello This is `+ name +` </h1>
                    </body>
                    </html>
                    `
	fmt.Println(print)
}
