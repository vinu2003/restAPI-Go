// Implements all Handler Routines.
package controller

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"encoding/base64"
	"encoding/json"

	"github.com/gorilla/mux"
)

const (
	username = "test"
	password = "password"
)
type Handler struct {
	database Database
}

func prettyprint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "  ", "\t")
	return out.Bytes(), err
}

func writeJson(w http.ResponseWriter, data interface{}) {
	bJson, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	prettyB, _ := prettyprint(bJson)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(prettyB)
}

type authHandler func(w http.ResponseWriter, r *http.Request)

// Authentication authenticates the user.
func Authentication(pass authHandler) authHandler {
	return func(w http.ResponseWriter, r *http.Request) {

		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			http.Error(w, "Error: authorization failed", http.StatusUnauthorized)
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !validate(pair[0], pair[1]) {
			http.Error(w, "Error: authorization failed", http.StatusUnauthorized)
			return
		}

		pass(w, r)
	}
}

func validate(user, pass string) bool {
	if username == user && password == pass {
		return true
	}
	return false
}

// ArticlesHandler creates a new record - POST METHOD.
func (h *Handler) ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	log.Println(string(body))
	if err != nil {
		log.Println("Error: in adding article - ", err)
		http.Error(w, "Error: in adding article.", http.StatusUnprocessableEntity)
		return
	}

	var articleStruct Article
	if err := json.Unmarshal(body, &articleStruct); err != nil {
		log.Println("Error: ArticlesHandler - Unmarshalling data, ", err)
		http.Error(w, "Error: ArticlesHandler - Unmarshalling data", http.StatusUnprocessableEntity)
		return
	}
	log.Println(articleStruct)

	// write into database
	id, err := h.database.AddArticle(articleStruct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Add("Location", "articles/"+strconv.Itoa(id))
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Added the article successfully..."))
	return
}

// GetArticleByID retrives record by 'id' - GET METHOD.
func (h *Handler) GetArticleByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars)

	id := vars["id"]
	fmt.Println(id)

	articleID, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Error: getting the product ID, ", err)
		http.Error(w, "Error: getting the product ID.", http.StatusUnprocessableEntity)
		return
	}

	article, err := h.database.GetArticleByID(articleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		log.Println(err.Error())
		return
	}
	log.Println(article)

	writeJson(w, article)
	return
}

func unique(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// GetARticleByTagNameDate retrives records which match with user provide 'tag' and 'date' - GET METHOD.
func (h *Handler) GetArticleByTagNameDate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars)

	tagName := vars["tagName"]
	log.Println(tagName)

	dateInfo := vars["date"]
	log.Println(dateInfo)

	// assumption are made for the date as below
	// the complete date len is 8 ie 20160112 which is straight forward to split.
	// if provided as 2016112, then assumptions can be 20160112 or 20161102 or 20161120
	// for now lets request the user to provide valid date with eg:20160112
	if len(dateInfo) != 8 {
		http.Error(w, "Error: Invalid Date Entered.", http.StatusUnprocessableEntity)
		return
	}
	dateRune := []rune(dateInfo)
	date := string(dateRune[:4])+"-"+string(dateRune[4:6])+"-"+string(dateRune[6:])
	log.Println(date)

	articles, err := h.database.GetArticleByTagDate(tagName, date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		log.Println(err.Error())
		return
	}

	// fill in the ArticleTagDate model.
	var result ArticleTagDate
	result.Tag   = tagName
	result.Count = len(articles)
	var ct int = 0
	for i := range articles {
		if ct < 10 {
			result.Articles = append(result.Articles, strconv.Itoa(articles[len(articles)-1-i].ID))
			ct += 1
		} else {
			break
		}
	}
	var tagSlice []string
	for _, tags := range articles {
		for _, val := range tags.Tags {
			if val != tagName {
				tagSlice = append(tagSlice, val)
			}
		}
	}
	result.Related_tags = unique(tagSlice)
	writeJson(w, result)
}

func (h *Handler) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	log.Println(string(body))
	if err != nil {
		log.Println("Error: in adding article - ", err)
		http.Error(w, "Error: in adding article.", http.StatusUnprocessableEntity)
		return
	}

	var articleStruct Article
	if err := json.Unmarshal(body, &articleStruct); err != nil {
		log.Println("Error: ArticlesHandler - Unmarshalling data, ", err)
		http.Error(w, "Error: ArticlesHandler - Unmarshalling data", http.StatusUnprocessableEntity)
		return
	}
	log.Println(articleStruct)

	// delete from database
	_, err = h.database.DeleteArticle(articleStruct)
	if err != nil {
		log.Println("Error: DeleteHandler -", err.Error())
		if strings.Contains(err.Error(), "not found") {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted article successfully..."))
}