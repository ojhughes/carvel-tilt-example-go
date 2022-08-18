package main

import (
	"fmt"
	"github.com/ojhughes/carvel-tilt-example-go/pkg/api"
	"log"
	"net/http"
	"os"
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

func getAddress() string {
	if port := os.Getenv("PORT"); port != "" {
		return fmt.Sprintf(":%s", port)
	} else {
		return fmt.Sprintf(":%d", defaultPort)
	}

}
