package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/nikp359/petstore/internal/server"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-c
		log.Printf("system call:%+v", oscall)
		cancel()
	}()

	if err := server.Serve(ctx); err != nil {
		log.Printf("failed to stop serve:+%v\n", err)
	}
}
