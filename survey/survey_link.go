package survey

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

const baseUrl = "https://app.wehelpsoftware.com/survey_persons/link"

type requiredFieldError struct {
	message string
}

func (e *requiredFieldError) Error() string {
	return e.message
}

func Generate(data map[string]interface{}, encryptKey string) (string, error) {
	err := validationRequiredFields(data)

	if err != nil {
		return "", err
	}

	headerMap := map[string]string{
		"alg": "HS256",
		"typ": "JWT",
	}

	headerBytes, err := json.Marshal(headerMap)

	if err != nil {
		return "", err
	}

	dataBytes, err := json.Marshal(data)

	if err != nil {
		return "", err
	}

	payload := string(dataBytes)
	header := string(headerBytes)

	hmacBytes := computeHMAC([]byte(header+payload), []byte(encryptKey))

	accessToken := base64UrlEncode([]byte(header)) + "." + base64UrlEncode([]byte(payload)) + "." + base64Encode(hmacBytes)
	return buildURL("?access_token=" + accessToken), nil
}

func validationRequiredFields(data map[string]interface{}) error {
	requiredKeys := []string{"code", "experience_id", "experience_date", "company_unit_code", "person"}

	for _, key := range requiredKeys {
		if _, ok := data[key]; !ok {
			return &requiredFieldError{message: fmt.Sprintf("Required field %s not found.", key)}
		}
	}

	if person, ok := data["person"].(map[string]interface{}); ok {
		personRequiredKeys := []string{"name", "internal_code", "type", "company_unit_code"}

		for _, key := range personRequiredKeys {
			if _, ok := person[key]; !ok {
				return &requiredFieldError{message: fmt.Sprintf("Required person field %s not found.", key)}
			}
		}
	} else {
		return &requiredFieldError{message: "Error person must be an instance of map[string]interface{}."}
	}

	return nil
}

func base64UrlEncode(input []byte) string {
	inputBase64 := base64.URLEncoding.EncodeToString(input)
	return strings.Replace(strings.Replace(inputBase64, "+", "_", -1), "/", "-", -1)
}

func base64Encode(input []byte) string {
	inputBase64 := base64.StdEncoding.EncodeToString(input)
	return strings.Replace(strings.Replace(inputBase64, "+", "_", -1), "/", "-", -1)
}

func buildURL(value string) string {
	return baseUrl + value
}

func computeHMAC(data, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	return h.Sum(nil)
}
