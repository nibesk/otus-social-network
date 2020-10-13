package utils

import "github.com/go-playground/validator"

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func CreateValidator() {
	validate = validator.New()
}

func GetValidator() *validator.Validate {
	return validate
}
