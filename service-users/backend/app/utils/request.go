package utils

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"strings"
)

const MaxRequestSize = 1048576

type MalformedRequest struct {
	Status int
	Msg    string
}

func (mr *MalformedRequest) Error() string {
	return mr.Msg
}

func DecodeJSONBody(dst interface{}, w http.ResponseWriter, r *http.Request) error {
	if !IsJsonRequest(r) {
		return &MalformedRequest{Status: http.StatusUnsupportedMediaType, Msg: "Content-Type header is not application/json"}
	}

	r.Body = http.MaxBytesReader(w, r.Body, MaxRequestSize)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}

		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := fmt.Sprintf("Request body contains badly-formed JSON")
			return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}

		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}

		case errors.Is(err, io.EOF):
			return &MalformedRequest{Status: http.StatusBadRequest, Msg: "Request body must not be empty"}

		case err.Error() == "http: request body too large":
			msg := fmt.Sprintf("Request body must not be larger than %d bytes", MaxRequestSize)
			return &MalformedRequest{Status: http.StatusRequestEntityTooLarge, Msg: msg}

		default:
			return errors.WithStack(err)
		}
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return &MalformedRequest{Status: http.StatusBadRequest, Msg: "Request body must only contain a single JSON object"}
	}

	return nil
}

func IsJsonRequest(r *http.Request) bool {
	contentType := r.Header.Get("Content-Type")

	return "application/json" == contentType
}
