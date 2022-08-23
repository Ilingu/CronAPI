package main

import (
	"cron-api/cmd/utils"
	"encoding/json"
	"errors"
	"net/http"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, DELETE")
}

type RequestShape struct {
	CronFrequency string `json:"Frequency"`
	CallbackUrl   string `json:"CallbackUrl"`
}

func ReadBody(req *http.Request) (RequestShape, error) {
	var jsonBody RequestShape

	err := json.NewDecoder(req.Body).Decode(&jsonBody)
	if err != nil {
		return jsonBody, err
	}

	defer req.Body.Close()
	if utils.IsEmptyString(jsonBody.CronFrequency) || utils.IsEmptyString(jsonBody.CallbackUrl) {
		return RequestShape{}, errors.New("invalid args")
	}
	if !utils.IsValidUrl(jsonBody.CallbackUrl) {
		return RequestShape{}, errors.New("invalid callback_url")
	}

	return jsonBody, nil
}

type ResponseShape[T any] struct {
	Succeed bool `json:"success"`
	Data    T    `json:"data,omitempty"`
}

func WriteResponse[T any](w *http.ResponseWriter, code int, datas ...T) {
	ResponsePayload := ResponseShape[T]{}
	if code < 400 {
		ResponsePayload.Succeed = true
	}

	if len(datas) > 0 {
		ResponsePayload.Data = datas[0]
	}

	enableCors(w)
	json.NewEncoder(*w).Encode(ResponsePayload)
}
