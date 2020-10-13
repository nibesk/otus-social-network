package utils

import (
	"encoding/json"
	"net/http"
)

type ResponseSuccessStruct struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

type ResponseErrorsStruct struct {
	Status bool     `json:"status"`
	Errors []string `json:"errors"`
}

func ResponseSuccessMessage(message string) ResponseSuccessStruct {
	return ResponseSuccessStruct{
		Status: true,
		Data: map[string]string{
			"message": message,
		},
	}
}

func ResponseData(data interface{}) ResponseSuccessStruct {
	return ResponseSuccessStruct{
		Status: true,
		Data:   data,
	}
}

func ResponseErrors(errors interface{}) ResponseErrorsStruct {
	return ResponseErrorsStruct{
		Status: false,
		Errors: errors.([]string),
	}
}

func ResponseErrorMessage(message string) ResponseErrorsStruct {
	return ResponseErrorsStruct{
		Status: false,
		Errors: []string{message},
	}
}

func SendResponseJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func SendResponseJsonWithStatusCode(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
