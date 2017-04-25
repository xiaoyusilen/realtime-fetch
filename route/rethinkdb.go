// author by @xiaoyusilen

package route

import (
	log "github.com/Sirupsen/logrus"
	r "gopkg.in/dancannon/gorethink.v2"
)

func NewRethinkdb() *r.Session {
	var err error

	session, err := r.Connect(r.ConnectOpts{
		Address:    "localhost:28015",
		InitialCap: 10,
		MaxOpen:    10,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Info("rethinkdb connect success!")
	return session
}
