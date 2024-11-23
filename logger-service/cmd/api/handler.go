package main

import (
	"fmt"
	"logger-service/cmd/data"
	"net/http"
)

type jsonPayload struct {
	Name string `json:"name"`
	Data string `json:"data`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	var requestPayload jsonPayload

	err := app.readJSON(w, r, requestPayload)
	if err != nil {
		fmt.Print("error reading jsonData", err)
		return
	}

	event := data.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}
	err = app.Models.LogEntry.Insert(event)
	if err != nil {
		fmt.Print("error inserting  jsonData", err)
		return
	}

	resp := jsonResponse{
		Error:   false,
		Message: "logged",
	}

	app.writeJSON(w, resp, http.StatusAccepted)

}
