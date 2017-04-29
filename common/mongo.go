// author by @xiaoyusilen

package common

import (
	log "github.com/Sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
)

// In main func, init the MongoDB, then you will get a *mgo.Session,
// If you use a Database pool, you can use session.Copy(),
// then close it when your query done.

func NewMongo() *mgo.Session {
	session, err := mgo.Dial("localhost:27017")

	if err != nil {
		log.Fatalln(err.Error())
		return nil
	}

	log.Debug("mongodb connect successed.")
	return session
}
