package survey

import (
	"encoding/base64"
	"encoding/json"
	"strings"
)

type SurveyLink struct {
	queryParams string
}

const BASE_URL = "https://app.wehelpsoftware.com/survey_persons/link"

func Generate(data map[string]interface{}) (*SurveyLink, error) {
	dataBytes, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	accessToken := base64UrlEncode(dataBytes)

	return &SurveyLink{queryParams: "?access_token=" + accessToken}, nil
}

func base64UrlEncode(input []byte) string {
	inputBase64 := base64.URLEncoding.EncodeToString(input)
	return strings.Replace(strings.Replace(inputBase64, "+", "_", -1), "/", "-", -1)
}

func (s *SurveyLink) GetUrl() string {
	return BASE_URL + s.queryParams
}
