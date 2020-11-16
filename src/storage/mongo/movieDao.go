package storage

import (
	"github.com/CronNieves/Api-rest-go/src/models"
	"gopkg.in/mgo.v2/bson"
)

const movieCollection = "movie"

// MovieAdd add one movie in MongoDB
func MovieAdd(movieData models.Movie) error {
	session := getSession()
	err := session.DB(dbName).C(movieCollection).Insert(movieData)
	return err
}

// MovieList Return all movie from MongoDB
func MovieList() (models.Movies, error) {
	var result models.Movies
	session := getSession()
	err := session.DB(dbName).C(movieCollection).Find(nil).Sort("-_id").All(&result)
	return result, err
}

// MovieList Return all movie from MongoDB
func MovieFindOneById(id bson.ObjectId) (models.Movie, error) {
	var result models.Movie
	session := getSession()
	err := session.DB(dbName).C(movieCollection).FindId(id).One(&result)
	return result, err
}

func UpdateMovie(oid bson.ObjectId, movie models.Movie) error {
	session := getSession()
	document := bson.M{"_id": oid}
	change := bson.M{"$set": movie}
	err := session.DB(dbName).C(movieCollection).Update(document, change)
	return err
}
