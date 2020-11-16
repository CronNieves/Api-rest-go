package storage

import (
	"github.com/CronNieves/Api-rest-go/src/models"
)

const movieCollection = "movie"

func MovieAdd(movieData models.Movie) {
	session := getSession()
	session.DB(dbName).C(movieCollection).Insert(movieData)
}
