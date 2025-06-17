package helpers

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

func MapToStruct(m map[string]interface{}, out interface{}) error {
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, out)
}

var validate = validator.New()

func ValidateRequest(req interface{}) error {
	return validate.Struct(req)
}
