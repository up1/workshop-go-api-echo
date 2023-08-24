package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HelloResponse struct {
	Message string `json:"message"`
}

func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/hello", func(c echo.Context) error {
		res := HelloResponse{Message: "Hello world"}
		return c.JSON(http.StatusOK, res)
	})

	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 1)

	// Start the service listening for requests.
	go func() {
		log.Printf("main : API listening on 1323")
		serverErrors <- e.Start(":1323")
	}()

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// =========================================================================
	// Shutdown

	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		log.Fatalf("error: starting server: %s", err)

	case <-shutdown:
		log.Println("main : Start shutdown")

		// Give outstanding requests a deadline for completion.
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Asking listener to shutdown and load shed.
		err := e.Shutdown(ctx)
		if err != nil {
			log.Printf("main : Graceful shutdown did not complete in %v : %v", 5*time.Second, err)
			err = e.Close()
		}

		if err != nil {
			log.Fatalf("main : could not stop server gracefully : %v", err)
		}
	}
}
