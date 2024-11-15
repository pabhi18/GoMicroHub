package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

const webPort = "7070"

type Config struct {
	DB     *sql.DB
	models data.Models
}

func main() {
	log.Println("Starting Authentication Services")

	// connect db here

	// set config
	app := Config{}
	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Panicln(err)
		panic(err)
	}

}
