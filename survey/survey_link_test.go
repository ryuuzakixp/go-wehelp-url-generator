package survey

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	encryptKey := "123"
	data := map[string]interface{}{
		"code":              "123",
		"experience_id":     "456",
		"experience_date":   "2023-01-01 11:00:00",
		"company_unit_code": "789",
		"person": map[string]interface{}{
			"name":              "John Doe",
			"internal_code":     "12345",
			"type":              "CUSTOMER",
			"company_unit_code": "789",
		},
	}

	expectedURL := "https://app.wehelpsoftware.com/survey_persons/link?access_token="
	sut, error := Generate(data, encryptKey)

	assert.NoError(t, error)
	assert.True(t, strings.HasPrefix(sut, expectedURL), "Unexpected URL prefix")
}

func TestError(t *testing.T) {
	encryptKey := "123"

	data := map[string]interface{}{
		"code":              "123",
		"experience_id":     "456",
		"experience_date":   "2023-01-01 11:00:00",
		"company_unit_code": "789",
		"person": map[string]interface{}{
			"name":              "John Doe",
			"internal_code":     "12345",
			"type":              "CUSTOMER",
			"company_unit_code": "789",
		},
	}

	data1 := copyMap(data)
	delete(data1, "code")
	_, error1 := Generate(data1, encryptKey)
	assert.EqualError(t, error1, "Required field code not found.")

	data2 := copyMap(data)
	delete(data2, "experience_id")
	_, error2 := Generate(data2, encryptKey)
	assert.EqualError(t, error2, "Required field experience_id not found.")

	data3 := copyMap(data)
	delete(data3, "experience_date")
	_, error3 := Generate(data3, encryptKey)
	assert.EqualError(t, error3, "Required field experience_date not found.")

	data4 := copyMap(data)
	delete(data4, "company_unit_code")
	_, error4 := Generate(data4, encryptKey)
	assert.EqualError(t, error4, "Required field company_unit_code not found.")

	data5 := copyMap(data)
	delete(data5, "person")
	_, error5 := Generate(data5, encryptKey)
	assert.EqualError(t, error5, "Required field person not found.")

	data6 := copyMap(data)
	if person6, ok := data6["person"].(map[string]interface{}); ok {
		delete(person6, "name")
		_, error6 := Generate(data6, encryptKey)
		assert.EqualError(t, error6, "Required person field name not found.")
	} else {
		t.Errorf("key person not found")
	}

	data7 := copyMap(data)
	if person7, ok := data7["person"].(map[string]interface{}); ok {
		delete(person7, "internal_code")
		_, error7 := Generate(data7, encryptKey)
		assert.EqualError(t, error7, "Required person field internal_code not found.")
	} else {
		t.Errorf("key person not found")
	}

	data8 := copyMap(data)
	if person8, ok := data8["person"].(map[string]interface{}); ok {
		delete(person8, "type")
		_, error8 := Generate(data8, encryptKey)
		assert.EqualError(t, error8, "Required person field type not found.")
	} else {
		t.Errorf("key person not found")
	}

	data9 := copyMap(data)
	if person9, ok := data9["person"].(map[string]interface{}); ok {
		delete(person9, "company_unit_code")
		_, error9 := Generate(data9, encryptKey)
		assert.EqualError(t, error9, "Required person field company_unit_code not found.")
	} else {
		t.Errorf("key person not found")
	}
}

func copyMap(original map[string]interface{}) map[string]interface{} {
	copy := make(map[string]interface{})

	for key, value := range original {
		switch v := value.(type) {
		case map[string]interface{}:
			copy[key] = copyMap(v)
		default:
			copy[key] = v
		}
	}

	return copy
}
