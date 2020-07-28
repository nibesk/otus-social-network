package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseMessage(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func ResponseData(data interface{}) map[string]interface{} {
	return map[string]interface{}{"status": true, "data": data}
}

func ResponseErrors(errors interface{}) map[string]interface{} {
	return map[string]interface{}{"status": false, "errors": errors}
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
