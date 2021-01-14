package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func Serve(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello Dude")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v\n", err)
		}
	}()

	fmt.Println("Server started 🚀")

	<-ctx.Done()

	fmt.Println("Server stopped ✋")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second) //nolint:gomnd
	defer func() {                                                                  //nolint:wsl
		cancel()
	}()

	if err := srv.Shutdown(ctxShutDown); err != nil {
		return fmt.Errorf("server Shutdown Failed: %w", err)
	}

	log.Printf("server exited properly 👌")

	return nil
}
