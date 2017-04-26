// author by @xiaoyusilen

package route

import (
	"github.com/xiaoyusilen/realtime-fetch/model"

	log "github.com/Sirupsen/logrus"
)

func CheckChannel(ch chan model.Test) {
	ok := true
	var msg model.Test
	for ok {
		if msg, ok = <-ch; ok {
			// You can do what you want to the realtime data.
			log.Println(msg)

			// send msg to Restful API
			res := SendToAPI(msg)
			if res {
				log.Debugf("send to api successed.")
			} else {
				log.Debugf("send to api failed.")
			}
		}
	}
}
