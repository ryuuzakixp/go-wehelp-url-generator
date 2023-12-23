package survey

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

const BASE_URL = "https://app.wehelpsoftware.com/survey_persons/link"

type RequiredFieldError struct {
	message string
}

func Generate(data map[string]interface{}) (string, error) {
	err := validationRequiredFields(data)

	if err != nil {
		return "", err
	}

	dataBytes, err := json.Marshal(data)

	if err != nil {
		return "", err
	}

	accessToken := base64UrlEncode(dataBytes)
	return buildURL("?access_token=" + accessToken), nil
}

func validationRequiredFields(data map[string]interface{}) error {
	requiredKeys := []string{"code", "experience_id", "experience_date", "company_unit_code", "person"}

	for _, key := range requiredKeys {
		if _, ok := data[key]; !ok {
			return &RequiredFieldError{message: fmt.Sprintf("Required field %s not found.", key)}
		}
	}

	if person, ok := data["person"].(map[string]interface{}); ok {
		personRequiredKeys := []string{"name", "internal_code", "type", "company_unit_code"}

		for _, key := range personRequiredKeys {
			if _, ok := person[key]; !ok {
				return &RequiredFieldError{message: fmt.Sprintf("Required person field %s not found.", key)}
			}
		}
	} else {
		return &RequiredFieldError{message: "Error person must be an instance of map[string]interface{}."}
	}

	return nil
}

func base64UrlEncode(input []byte) string {
	inputBase64 := base64.URLEncoding.EncodeToString(input)
	return strings.Replace(strings.Replace(inputBase64, "+", "_", -1), "/", "-", -1)
}

func buildURL(value string) string {
	return BASE_URL + value
}

func (e *RequiredFieldError) Error() string {
	return e.message
}
