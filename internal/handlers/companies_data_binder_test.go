package handlers

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestBindCompanyData(t *testing.T) {
	e := echo.New()

	testData := map[string]string{"name": "my_test_company", "code": "23323", "country": "Ukraine", "website": "https://something.com", "phone": "23323"}
	testDataM, _ := json.Marshal(testData)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(echo.POST, "/company/add", bytes.NewReader(testDataM))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	c := e.NewContext(req, rec)

	companyRequest, err := bindCompanyData(c)

	if assert.NoError(t, err) {
		assert.Equal(t, companyRequest.Name, testData["name"])
		assert.Equal(t, companyRequest.Code, testData["code"])
		assert.Equal(t, companyRequest.Country, testData["country"])
		assert.Equal(t, companyRequest.Website, testData["website"])
		assert.Equal(t, companyRequest.Phone, testData["phone"])
	}
}

func TestBindCompanyDataInvalidField(t *testing.T) {
	tests := []struct {
		name     string
		testData map[string]string
		err      string
	}{
		{
			name:     "Company name validation",
			testData: map[string]string{"name": "", "code": "23323", "country": "Ukraine", "website": "https://something.com", "phone": "23323"},
			err:      "Company validation failed",
		},
		{
			name:     "Code validation",
			testData: map[string]string{"name": "my_test_company", "code": "", "country": "Ukraine", "website": "https://something.com", "phone": "23323"},
			err:      "Code validation failed",
		},
		{
			name:     "Country validation",
			testData: map[string]string{"name": "my_test_company", "code": "23323", "country": "", "website": "https://something.com", "phone": "23323"},
			err:      "Country validation failed",
		},
		{
			name:     "Website validation",
			testData: map[string]string{"name": "my_test_company", "code": "23323", "country": "Ukraine", "website": "", "phone": "23323"},
			err:      "Website validation failed",
		},
		{
			name:     "Phone validation",
			testData: map[string]string{"name": "my_test_company", "code": "23323", "country": "Ukraine", "website": "https://something.com", "phone": ""},
			err:      "Phone validation failed",
		},
		{
			name:     "Empty file",
			testData: map[string]string{},
			err:      "Loaded empty file",
		},
	}

	e := echo.New()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testDataM, _ := json.Marshal(tt.testData)

			rec := httptest.NewRecorder()
			req := httptest.NewRequest(echo.POST, "/company/add", bytes.NewReader(testDataM))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			c := e.NewContext(req, rec)

			_, err := bindCompanyData(c)
			if err == nil {
				t.Fatal(tt.err)
			}
		})
	}
}
