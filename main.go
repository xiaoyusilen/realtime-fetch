// author by @xiaoyusilen

package main

import (
	"time"

	"github.com/xiaoyusilen/realtime-fetch/model"
	"github.com/xiaoyusilen/realtime-fetch/route"
)

func main() {

	// Start RethinkDB
	c := route.NewRethinkdb()

	// New a channel to save the realtime data
	ch := make(chan model.Test, 10000)

	// One goroutine: read from channel
	go route.CheckChannel(ch)

	for {
		// Another goroutine: fetch realtime data
		go route.FetchRealData(c, ch)

		// Time sleep control the fetch time, I set it as 10 seconds
		time.Sleep(time.Second * 10)
	}

}
