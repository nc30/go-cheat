package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	go func() {
		// too many time function
		time.Sleep(time.Second * 10)
	}()

	<-ctx.Done()

	log.Println("done")
}
