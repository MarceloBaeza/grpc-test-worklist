package main

import (
	"log"
	"net/http"

	"github.com/MarceloBaeza/grpc-test-worklist/client/configuration"
	"github.com/MarceloBaeza/grpc-test-worklist/client/services"
)

func main() {
	ConnectionStruct := configuration.NewConnection()
	Services := services.NewServiceStruct(ConnectionStruct.Client)
	// Server
	http.HandleFunc("/", Services.NewWork)

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
