package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

// type jsonResponse struct {
// 	Error   bool   `json:"error"`
// 	Message string `json:"message"`
// 	Data    any    `json:"data,omitempty"`
// }
//

type RequestPayLoad struct {
	Action string      `json:"action"`
	Auth   AuthPayLoad `json:"auth,omitempty`
}

type AuthPayLoad struct {
	Email    string `json:"email`
	Password string `json:"password"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)

	// out, _ := json.MarshalIndent(payload, "", "\t")
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusAccepted)
	// w.Write(out)
}

func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayLoad RequestPayLoad

	err := app.readJSON(w, r, &requestPayLoad)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	switch requestPayLoad.Action {
	case "auth":
		app.authenticate(w, requestPayLoad.Auth)

	default:
		app.errorJSON(w, errors.New("La action disconocida"))
	}
}

func (app *Config) authenticate(w http.ResponseWriter, a AuthPayLoad) {
	jsonData, _ := json.MarshalIndent(a, "", "\t")

	request, err := http.NewRequest("POST", "http://authentication-service/authenticate", bytes.NewBuffer(jsonData))
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

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("datos invalidos"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error al llamar servicio de autenicacion"))
		return
	}

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

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Authenticated!"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusAccepted, payload)
}
