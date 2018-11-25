package controller

// Each Article representation.
type Article struct {
	ID       int      `bson:"_id"`
	Title    string   `json:"title"`
	Date     string   `json:"date"`
	Body     string   `json:"body"`
	Tags     []string `json:"tags"`
}

// response model for tagName&Date query.
type ArticleTagDate struct {
	Tag            string   `json:"tag"`
	Count          int      `json:"count"`
	Articles       []string `json:"articles"`
	Related_tags   []string `json:"related_tags"`
}
// Articles is array of Article objects.
type ArticlesArr []Article
