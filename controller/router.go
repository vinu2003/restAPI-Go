// Router Initialization.
package controller

import (
	"github.com/gorilla/mux"
)

var handler = &Handler{Database{}}

func Router() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/articles", Authentication(handler.ArticlesHandler))
	r.HandleFunc("/articles/{id}", Authentication(handler.GetArticleByID))
	r.HandleFunc("/tag/{tagName}/{date}", Authentication(handler.GetArticleByTagNameDate))
	return r
}
