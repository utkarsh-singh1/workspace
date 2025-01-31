package main

import (
	"context"
	"fmt"
)

func main() {

	ctx := context.Background()
	
	fmt.Println("context",ctx)

	fmt.Println("context err",ctx.Err())
}
