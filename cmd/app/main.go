package main

import (
	"fmt"
	"github.com/ojhughes/carvel-tilt-example-go/pkg/api"
	"log"
	"net/http"
	"os"
	"time"
)

const defaultPort = 8083

func main() {
	logger := log.New(os.Stdout, "", 0)
	err := http.ListenAndServe(getAddress(), api.NewHandler(api.LogWith(logger)))
	if err != nil {
		logger.Fatal("Error starting HTTP server on address %s", getAddress())
		return
	}
}

func newHTTPServer(address string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:         address,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}

func getAddress() string {
	if port := os.Getenv("PORT"); port != "" {
		return fmt.Sprint(":%s", port)
	} else {
		return fmt.Sprint(":%s", defaultPort)
	}

}
