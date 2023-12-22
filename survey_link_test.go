package survey

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	data := map[string]interface{}{
		"code":              "123",
		"experience_id":     "456",
		"experience_date":   "2023-01-01",
		"company_unit_code": "789",
		"person": map[string]interface{}{
			"name":              "John Doe",
			"internal_code":     "12345",
			"type":              "employee",
			"company_unit_code": "789",
		},
	}
	expectedURL := "https://app.wehelpsoftware.com/survey_persons/link?access_token="
	sut, _ := Generate(data)
	assert.True(t, strings.HasPrefix(sut.GetUrl(), expectedURL), "Unexpected URL prefix")
	fmt.Println(sut.GetUrl())
}
