package utils

import "github.com/gookit/validate"

func ValidateStruct(data any) validate.Errors {
	val := validate.Struct(data)
	val.Validate()
	return val.Errors
}
