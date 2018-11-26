package controller

import (
	"bytes"
	"strconv"
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

	// Check error message
	if !bytes.Equal(rr.Body.Bytes(),
		[]byte("Error: Failed to retrive the article with ID, not found\n")) {
		t.Errorf("handler returned unexpected body: got %v want \n%v",
			rr.Body.Bytes(),
			[]byte("Error: Failed to retrive the article with ID, not found\n"))
	}
}

func TestHandler_GetArticleByIDInValidIDValue(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8984/articles/$$", nil)
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

	// Check error message
	if !bytes.Equal(rr.Body.Bytes(),
		[]byte("Error: getting the product ID.\n")) {
		t.Errorf("handler returned unexpected body: got %v want \n%v",
			rr.Body.Bytes(),
			[]byte("Error: getting the product ID.\n"))
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

	// Check error message
	if !bytes.Equal(rr.Body.Bytes(),
		[]byte("Error: Failed to retrive the articles for date&Tag, <nil>\n")) {
		t.Errorf("handler returned unexpected body: got %v want \n%v",
			rr.Body.Bytes(),
			[]byte("Error: Failed to retrive the articles for date&Tag, <nil>\n"))
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

	// Check error message
	if !bytes.Equal(rr.Body.Bytes(),
		[]byte("Error: Failed to retrive the articles for date&Tag, <nil>\n")) {
		t.Errorf("handler returned unexpected body: got %v want \n%v",
			rr.Body.Bytes(),
			[]byte("Error: Failed to retrive the articles for date&Tag, <nil>\n"))
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

	// Check error message
	if !bytes.Equal(rr.Body.Bytes(),
		[]byte("Error: Invalid Date Entered.\n")) {
		t.Errorf("handler returned unexpected body: got %v want \n%v",
			rr.Body.Bytes(),
			[]byte("Error: Invalid Date Entered.\n"))
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

	// Check error message
	if !bytes.Equal(rr.Body.Bytes(),
		[]byte("Error: Failed to retrive the articles for date&Tag, <nil>\n")) {
		t.Errorf("handler returned unexpected body: got %v want \n%v",
			rr.Body.Bytes(),
			[]byte("Error: Failed to retrive the articles for date&Tag, <nil>\n"))
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
	req, err := http.NewRequest("POST", "http://localhost:8984/articles", bytes.NewBuffer(data))
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

	// Check the error message.
	if !bytes.Equal(rr.Body.Bytes(),
		[]byte("Error: ArticlesHandler - Unmarshalling data : invalid character 't' looking for beginning of object key string\n")) {
		t.Errorf("handler returned unexpected body: got %v want \n%v",
			rr.Body.Bytes(),
			[]byte("Error: ArticlesHandler - Unmarshalling data : invalid character 't' looking for beginning of object key string\n"))
	}
}

func TestHandler_ArticlesHandlerValidInput(t *testing.T) {
	var handler = &Handler{Database{}}
	data := []byte(`{"id":1,"title":"","date":"2018-03-14","body":"Change in climate and vegetation","tags":["world","climate","nature"]}`)

	req, err := http.NewRequest("POST", "http://localhost:8984/articles", bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}
	req.SetBasicAuth("test", "password")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	h := http.HandlerFunc(Authentication(handler.ArticlesHandler))

	h.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// check header
	if location := rr.Header().Get("Location"); location == "articles/"+strconv.Itoa(handler.database.articlesID) {
		t.Errorf("handler returned wrong Location : got %v want %v",
			location, "articles/"+strconv.Itoa(handler.database.articlesID))
	}

	// Check the header message.
	if !bytes.Equal(rr.Body.Bytes(), []byte("Added the article successfully...")) {
		t.Errorf("handler returned unexpected body: got %v want \n%v",
			rr.Body.Bytes(),
			[]byte("Added the article successfully..."))
	}
}

func TestHandler_ArticlesHandlerDuplicateInput(t *testing.T) {
	var handler = &Handler{Database{}}
	data := []byte(`{"id":1,"title":"","date":"2018-03-14","body":"Change in climate and vegetation","tags":["world","climate","nature"]}`)

	req, err := http.NewRequest("POST", "http://localhost:8984/articles", bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}
	req.SetBasicAuth("test", "password")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	h := http.HandlerFunc(Authentication(handler.ArticlesHandler))

	h.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

	// Check the error message.
	if !bytes.Equal(rr.Body.Bytes(), []byte("Info: Article already exists in database, 15\n")) {
		t.Errorf("handler returned unexpected body: got %v want \n%v",
			rr.Body.Bytes(),
			[]byte("Info: Article already exists in database, 15\n"))
	}
}

func TestHandler_DeleteArticleInValidData(t *testing.T) {
	data := []byte(`{"id":1,"title":"my music","date":"3000-92-40","body":"music","tags":["songs"]}`)

	// first delete entry created during previous unit test execution.
	req, err := http.NewRequest("DELETE", "http://localhost:8984/article", bytes.NewBuffer(data))
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

	// Check the error message.
	if !bytes.Equal(rr.Body.Bytes(), []byte("Error: Data enter not found in database, not found\n")) {
		t.Errorf("handler returned unexpected body: got %v want \n%v",
			rr.Body.Bytes(),
			[]byte("Error: Data enter not found in database, not found\n"))
	}
}

func TestHandler_DeleteArticleValidInput(t *testing.T) {
	data := []byte(`{"id":1,"title":"","date":"2018-03-14","body":"Change in climate and vegetation","tags":["world","climate","nature"]}`)

	// first delete entry created during previous unit test execution.
	req, err := http.NewRequest("DELETE", "http://localhost:8984/article", bytes.NewBuffer(data))
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

	// Check the header message.
	if !bytes.Equal(rr.Body.Bytes(), []byte("Deleted article successfully...")) {
		t.Errorf("handler returned unexpected body: got %v want \n%v",
			rr.Body.Bytes(),
			[]byte("Deleted article successfully..."))
	}
}
