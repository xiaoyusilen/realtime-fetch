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

	// Close RethinkDB
	defer c.Close()

	// New a channel to save the realtime data
	ch := make(chan model.Test, 10000)

	// This part of code is ugly! I will refactor it someday.

	t := time.Now()

	go route.FetchRealData(c, ch, t, true)

	time.Sleep(time.Second * 10)

	for {

		// 获取最新数据
		go route.FetchRealData(c, ch, t, false)

		t = time.Now()

		// 每10s执行一次
		time.Sleep(time.Second * 10)
	}
}
