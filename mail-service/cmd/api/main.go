package main

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct {
}

const webPort = "8083"

func main() {
	app := Config{}

	log.Println("mail-service is starting...")

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("something is wrong to start server")
		panic(err)
	}
	log.Printf("server is running on port : %s", webPort)

}
