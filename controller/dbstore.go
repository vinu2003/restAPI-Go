// Package implements routines that executes mongo query pipelines.
package controller

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"awesomeProject/errors"
)

type Database struct {
	mutex sync.Mutex

	articlesID int
}

const (
	DBNAME = "ffdatabase"
	COLLECTION = "NewArtStore"
	)



// checkDuplicate checks if record provided by user already exists in database.
func checkDuplicate(data Article, db *mgo.Collection) bool {
	r := Article{}
	pipeline := []bson.M{{"$match":bson.M{"date":data.Date}},{"$match":bson.M{"title":data.Title}},{"$match":bson.M{"body":data.Body}},{"$match":bson.M{"tags": bson.M{"$in":data.Tags}}}}
	err := db.Pipe(pipeline).One(&r)
	if err != nil && strings.Contains(err.Error(), "not found") {
			log.Println("INFO: Record does not exist, ", err)
			return false
	}
	return true
}

// AddArticles insert the record into datbase - POST METHOD.
func (d Database) AddArticle(data Article) (int, error) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		return -1, errors.New(fmt.Sprintf("Error: Failed to establish connection to mongoDB server, %v", err))
	}
	defer session.Close()

	session.SetSafe(&mgo.Safe{})
	db := session.DB(DBNAME).C(COLLECTION)

	// first verify if the entry provided is duplicate.
	isExists := checkDuplicate(data, db)
	if isExists {
		log.Println("Info: Data enter already exists in database, \n",data)
		return -1, errors.New("Info: Data enter already exists in database")
	}

	numRec, err := db.Count()
	log.Println("numRec:", numRec)
	d.mutex.Lock()
	defer d.mutex.Unlock()
	if err != nil {
		log.Println("Error: Unable to get number of records exists from Database, ", err)
		d.articlesID += 1
		data.ID = d.articlesID
	} else {
		d.articlesID = numRec + 1
		data.ID = d.articlesID
	}

	err = db.Insert(data)
	if err != nil {
		return -1, errors.New(fmt.Sprintf("Error: adding the article, %v", err))
	}
	log.Println("Added new Article with id :", d.articlesID)
	return d.articlesID, nil
}

// GetArticleByID retrives the article with 'id' specified by user from datbase - GET METHOD.
func (d Database) GetArticleByID(id int) (Article, error) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		return Article{}, errors.New(fmt.Sprintf("Error: Failed to establish connection to mongoDB server, %v", err))
	}
	defer session.Close()

	db := session.DB(DBNAME).C(COLLECTION)

	result := Article{}

	if err := db.FindId(id).One(&result); err != nil {
		return result, errors.New(fmt.Sprintf("Error: Failed to retrive the article with ID, %v", err))
	}

	log.Println("Successfully retrived the article with ID - ", id)
	return result, nil
}

// GetArticleByTagDate retrieves array of Articles that matches the 'tag' and 'date' provided by user - GET METHOD.
func (d Database) GetArticleByTagDate(tagStr, dateStr string) (ArticlesArr, error) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		return ArticlesArr{}, errors.New(fmt.Sprintf("Error: Failed to establish connection to mongoDB server, %v", err))
	}
	defer session.Close()

	db := session.DB(DBNAME).C(COLLECTION)

	var result ArticlesArr

	tagArr := []string{tagStr}

	pipeline := []bson.M{{"$match":bson.M{"date":dateStr}},{"$match":bson.M{"tags": bson.M{"$in":tagArr}}}}
	log.Println(pipeline)
	err = db.Pipe(pipeline).All(&result)
	if err != nil || len(result)==0 {
		// <TODO> - when query returns nothing for invalid entries the pipe return nil.
		// need to debug this.
		return result, errors.New(fmt.Sprintf("Error: Failed to retrive the articles for date&Tag, %v", err))
	}
	return result, nil
}
