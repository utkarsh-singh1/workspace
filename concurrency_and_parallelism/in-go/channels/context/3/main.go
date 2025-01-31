package main

import (
	"context"
	"fmt"
)

func main() {

	ctx := context.Background()

	fmt.Println("Current context:\t",ctx)
	fmt.Println("Error with context:\t",ctx.Err())
	fmt.Printf("type of context:\t\t%T\n",ctx)


	ctx,_ = context.WithCancel(ctx)

	fmt.Println("\nCurrent context:\t",ctx)
	fmt.Println("Error with context:\t",ctx.Err())
	fmt.Printf("type of context:\t%T\n",ctx)
}
