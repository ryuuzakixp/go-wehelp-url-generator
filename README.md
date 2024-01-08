## Description
[![Go](https://github.com/ryuuzakixp/go-wehelp-url-generator/actions/workflows/test.yml/badge.svg)](https://github.com/ryuuzakixp/go-wehelp-url-generator/actions/workflows/test.yml)

Function to help generate valid url for surveys on Wehelp.

## Installation
```bash
$ go get github.com/ryuuzakixp/go-wehelp-url-generator/survey
```

## Usage

Example minimum data required

```Go
   
package main

import (
  "fmt"
  "github.com/ryuuzakixp/go-wehelp-url-generator/survey"
)

func main() {
  data := map[string]interface{}{
    "code":              "123",//change to valid code
    "experience_id":     "456",
    "experience_date":   "2023-01-01 11:00:00",//UTC
    "company_unit_code": "789",
    "person": map[string]interface{}{
	"name":              "John Doe",
	"internal_code":     "12345",
	"type":              "CUSTOMER",//CUSTOMER,COLLABORATOR
	"company_unit_code": "789",
	},
    }

  encryptKey := "your-encryption-key"// change to valid encryptKey

  surveyUrl, err := survey.Generate(data, encryptKey)

  if err != nil {
    fmt.Println("Error:", err)
    return
  }

  fmt.Println("Generated survey url:", surveyUrl)
}
```

Example full data

```Go
   
package main

import (
  "fmt"
  "github.com/ryuuzakixp/go-wehelp-url-generator/survey"
)

func main() {
  data := map[string]interface{}{
    "code":              "123",//change to valid code
    "experience_id":     "",
    "experience_date":   "2023-01-01 11:00:00",//UTC
    "company_unit_code": "789",
    "person": map[string]interface{}{
	"name":              "John Doe",
	"internal_code":     "12345",
	"type":              "CUSTOMER",//CUSTOMER,COLLABORATOR
	"company_unit_code": "789",
        "email":             "xxxxxx",
        "phone":             "xxxxx",
        "created_at":        "2020-01-01",//Y-m-d
        "date_of_birth":     "1970-07-06",//Y-m-d
        "language":          "xxxx",//PORTUGUESE,SPANISH,ENGLISH 
    },
    "cf": map[string]interface{}{
	"1": "value",//id:value
	"2": "value",
	"3": "value",
	"4": "value",
	},
    }

    encryptKey := "your-encryption-key"// change to valid encryptKey

    surveyUrl, err := survey.Generate(data, encryptKey)

    if err != nil {
	fmt.Println("Error:", err)
	return
    }

    fmt.Println("Generated survey url:", surveyUrl)
}
```
## Additional Information

If your company uses a unique code such as email or phone, you must enter these fields in your person object
