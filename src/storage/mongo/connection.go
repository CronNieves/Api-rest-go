package storage

import "gopkg.in/mgo.v2"

const dbName = "Api-Rest-Go"

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}

	return session
}
