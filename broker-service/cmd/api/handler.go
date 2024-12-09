package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
	Log    LogPayload  `json:"log,omitempty"`
	Mail   MailPayload `json:"mail,omitempty"`
}

type MailPayload struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
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
	case "mail":
		app.sendMail(w, requestPayload.Mail)

	default:
		app.errorJSON(w, errors.New("unknown action"))
		return
	}
}

func (app *Config) authenticate(w http.ResponseWriter, a AuthPayload) {
	// create some json we will send to auth service
	jsonData, _ := json.MarshalIndent(a, "", "\t")

	// call the service
	request, err := http.NewRequest("POST", "http://authentication-service:8081/auth", bytes.NewBuffer(jsonData))
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
	jsonData, _ := json.MarshalIndent(entry, "", "\t")

	loggerRequestUrl := "http://logger-service:8082/log"

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

	fmt.Println("printing logger response in broker service", jsonFromService)

	var responsePayload jsonResponse
	responsePayload.Error = false
	responsePayload.Message = "Logged in successfully"
	responsePayload.Data = jsonFromService
	app.writeJSON(w, responsePayload, http.StatusAccepted)
}

func (app *Config) sendMail(w http.ResponseWriter, mail MailPayload) {
	jsonData, _ := json.MarshalIndent(mail, "", "\t")

	emilRequestURL := "http://mailer-service:8083/send"

	req, err := http.NewRequest("POST", emilRequestURL, bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	response, err := client.Do(req)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
	}

	defer response.Body.Close()

	fmt.Println("mailer response in broker service", response)

	if response.StatusCode != http.StatusCreated {
		app.errorJSON(w, errors.New("invalid status code"))
		log.Println("invalid request")
		return
	}

	var res jsonResponse

	res.Error = false
	res.Message = "email sent" + mail.To
	res.Data = response.Body

	app.writeJSON(w, res, http.StatusAccepted)
}
