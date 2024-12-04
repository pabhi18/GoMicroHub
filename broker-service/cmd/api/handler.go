package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
	Log    LogPayload  `json:"log,omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json: "password`
}

type LogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	_ = app.writeJSON(w, payload, http.StatusOK)

}

func (app *Config) HandleSubmit(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
	}

	switch requestPayload.Action {
	case "auth":
		app.authenticate(w, requestPayload.Auth)
	case "log":
		app.logItem(w, requestPayload.Log)

	default:
		app.errorJSON(w, errors.New("unknown action"))
		return
	}
}

func (app *Config) authenticate(w http.ResponseWriter, a AuthPayload) {
	// create some json we will send to auth service
	jsonData, _ := json.MarshalIndent(a, "", "\t")

	// call the service
	request, err := http.NewRequest("POST", "http://authentication-service:8081", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	defer response.Body.Close()

	// make sure we get correct status code

	if response.StatusCode != http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid credentials "))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling auth service "))
		return
	}

	// create a variable read json from response body
	var jsonFromService jsonResponse
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var payload jsonResponse

	payload.Error = false
	payload.Message = "Authenticated !"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, payload, http.StatusAccepted)
}

func (app *Config) logItem(w http.ResponseWriter, entry LogPayload) {
	jsonData, _ := json.MarshalIndent(entry, "", "/t")

	loggerRequestUrl := "http://logger-service:8082"

	request, err := http.NewRequest("POST", loggerRequestUrl, bytes.NewBuffer(jsonData))

	if err != nil {
		log.Println("error while calling logger-service in broker service")
		app.errorJSON(w, err)
		return
	}
	client := http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("error while executing client on logger-service in broker service by client")
		app.errorJSON(w, err)
		return
	}

	defer response.Body.Close()

	var jsonFromService jsonResponse
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("invalid status code"))
		log.Println("invalid request")
		return
	}

	var responsePayload jsonResponse
	responsePayload.Error = true
	responsePayload.Message = "Logged"

	responsePayload.Data = response.Body

	app.writeJSON(w, responsePayload, http.StatusAccepted)
}
