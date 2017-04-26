// author by @xiaoyusilen

package route

import (
	"time"

	log "github.com/Sirupsen/logrus"
	r "gopkg.in/dancannon/gorethink.v2"

	"github.com/xiaoyusilen/realtime-fetch/model"
)

func FetchRealData(c *r.Session, ch chan model.Test, t time.Time, flag bool) {

	// Fetch realtime data from rethinkdb
	// Query all ids
	resp, err := r.DB("test").Table("test").Map(map[string]interface{}{
		"id": r.Row.Field("id"),
	}).Run(c)

	// Close the query
	defer resp.Close()

	// Deal with error, exit the goroutine
	if err != nil {
		if resp.IsNil() {
			log.Errorf("Have no data like this, err is: %s.", err)
			return
		}
		log.Errorf("Query error: %s.", err)
		return
	}

	// Change data struct from database to you need
	var res []map[string]string

	err = resp.All(&res)

	if err != nil {
		log.Errorf("Change data struct err: %s.", err)
		return
	}

	// Judge time
	for i := 0; i < len(res); i++ {
		resp, err = r.DB("lecar").Table("realtime_veh_info").Filter(map[string]string{
			"id": res[i]["id"],
		}).Run(c)

		var re model.Test
		err = resp.One(&re)

		if err != nil {
			log.Errorf("err is: %s", err)

		}

		// if flag == false, means if the data is unchanged, it won't be added to channel
		if !flag {
			if re.UpdatedAt.After(t) {
				ch <- re
			}
		} else {
			// if flag == true, every data will be added to channel
			ch <- re
		}
	}

	// return to stop the goroutine
	return
}
