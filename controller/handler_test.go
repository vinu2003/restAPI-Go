package controller

import (
	"bytes"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
)

func TestHandler_GetArticleByIDValidData(t *testing.T) {
	//data := []byte(`{"id":1,"title":"Global Warming","date":"2018-10-04","body":"Change in climate and vegetation","tags":["world","climate","nature"]}`)
	req, err := http.NewRequest("GET", "http://localhost:8984/articles/3", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetBasicAuth("test", "password")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	Router().ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	assert.Equal(t,
		`{
  	"ID": 3,
  	"title": "ABC",
  	"date": "2018-10-04",
  	"body": "My ABC",
  	"tags": [
  		"aaa",
  		"bbb",
  		"ccc"
  	]
  }`,
  rr.Body.String(),
  "handler returned unexpected body")
}

func TestHandler_GetArticleByIDInValidID(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8984/articles/32", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetBasicAuth("test", "password")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	Router().ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func TestHandler_GetArticleByTagNameDateInValidTag(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8984/tag/nnn/20181005", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetBasicAuth("test", "password")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	Router().ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func TestHandler_GetArticleByTagNameDateWithDateNotExists(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8984/tag/aaa/20221120", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetBasicAuth("test", "password")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	Router().ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func TestDatabase_GetArticleByTagDateInValidDate(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8984/tag/aaa/202210", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetBasicAuth("test", "password")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	Router().ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusUnprocessableEntity {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnprocessableEntity)
	}
}

func TestHandler_GetArticleByTagNameDateInValidDateInValidTag(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8984/tag/nnn/20221120", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetBasicAuth("test", "password")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	Router().ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func TestHandler_GetArticleByTagNameDate(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8984/tag/aaa/20181005", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetBasicAuth("test", "password")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	Router().ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	assert.Equal(t,
		`{
  	"tag": "aaa",
  	"count": 8,
  	"articles": [
  		"11",
  		"10",
  		"9",
  		"8",
  		"7",
  		"6",
  		"5",
  		"4"
  	],
  	"related_tags": [
  		"xxx",
  		"yyy",
  		"zzz",
  		"SSS",
  		"ooo",
  		"lll"
  	]
  }`,
		rr.Body.String(),
		"handler returned unexpected body")
}

func TestHandler_ArticlesHandlerInValidInput(t *testing.T) {
	data := []byte(`{"id":1,tle":"Global Warming","date":"2018-10-04","body":"Change in climate and vegetation","tags":["world","climate","nature"]}`)
	req, err := http.NewRequest("POST", "http://localhost:8984/articles",  bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}
	req.SetBasicAuth("test", "password")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	Router().ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusUnprocessableEntity {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnprocessableEntity)
	}
}

func TestHandler_ArticlesHandlerValidInput(t *testing.T) {
	data := []byte(`{"id":1,"title":"","date":"2018-03-14","body":"Change in climate and vegetation","tags":["world","climate","nature"]}`)
	req, err := http.NewRequest("POST", "http://localhost:8984/articles",  bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}
	req.SetBasicAuth("test", "password")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	Router().ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

func TestHandler_ArticlesHandlerDuplicateInput(t *testing.T) {
	data := []byte(`{"id":1,"title":"","date":"2018-03-14","body":"Change in climate and vegetation","tags":["world","climate","nature"]}`)
	req, err := http.NewRequest("POST", "http://localhost:8984/articles",  bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}
	req.SetBasicAuth("test", "password")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	Router().ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}
