package main

import (
	"fmt"
	"logger-service/data"
	"net/http"
	"time"
)

type JsonPayload struct {
	Name string `json:"name"`
	Data string `json:"data`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	var requestPayload JsonPayload

	err := app.readJSON(w, r, requestPayload)
	if err != nil {
		fmt.Println("error reading jsonData in logger controller", err)
		return
	}

	event := data.LogEntry{
		Name:      requestPayload.Name,
		Data:      requestPayload.Data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = app.Models.LogEntry.Insert(event)
	if err != nil {
		fmt.Println("error inserting  jsonData", err)
		return
	}

	resp := jsonResponse{
		Error:   false,
		Message: "Logged in",
		Data:    "log data is inserted successfully",
	}

	app.writeJSON(w, resp, http.StatusAccepted)

}
