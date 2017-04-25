// author by @xiaoyusilen

package route

import (
	log "github.com/Sirupsen/logrus"
	r "gopkg.in/dancannon/gorethink.v2"

	"github.com/xiaoyusilen/realtime-fetch/model"
)

func FetchRealData(c *r.Session, ch chan model.Test) {

	// Fetch realtime data from rethinkdb
	resp, err := r.DB("test").Table("test").Map(map[string]interface{}{
		"id": r.Row.Field("id"),
	}).Run(c)

	// Deal with error
	if err != nil {
		log.Errorf("Query error: %s.", err)
	}

	// Change data struct from database to you need
	var res map[string]string

	err = resp.One(&res)

	if err != nil {
		log.Errorf("Change data struct err: %s.", err)
	}

	test := model.Test{
		ID:   res["id"],
		Name: res["name"],
	}

	// Write data to channel
	ch <- test
}
