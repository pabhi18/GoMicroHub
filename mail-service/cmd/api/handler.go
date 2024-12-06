package main

import (
	"fmt"
	"net/http"
)

func (app *Config) SendMail(w http.ResponseWriter, r *http.Request) {
	type mailMessage struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}

	var requestPayload mailMessage

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		fmt.Println("error reading client data", err)
		app.errorJSON(w, err)
		return
	}

	msg := Message{
		From:    requestPayload.From,
		To:      requestPayload.To,
		Subject: requestPayload.Subject,
		Data:    requestPayload.Message,
	}

	err = app.Mailer.SendSMPTMessage(msg)
	if err != nil {
		fmt.Println("error sending email to the client", err)
		app.errorJSON(w, err)
		return
	}

	payLoad := jsonResponse{
		Error:   false,
		Message: "sent to" + requestPayload.To,
	}

	app.writeJSON(w, payLoad, http.StatusCreated)

}
