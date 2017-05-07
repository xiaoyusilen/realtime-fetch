# Realtime-fetch
A simple project of fetching real-time data and write to channel then send to one Restful API. 🤔
[Introduction](http://xiaoyu.world/2017/04/26/go-fetch-real-time-data/)⬅️

**What is real-time data?**
Real-time data means the data changed all the time. Data struct is unimportant.

# When you can use this repo
If you :
- have a table contains real-time data
- not too much data
- you want to deal with the real-time data

Congratulations! 🎉 You can use this repo.

# About this repo
1. I use channel as a message queue, of course you can use `RabbitMQ` or others.
2. I use `RethinkDB` as my database, you can use others if you like.But you should change the code yourself.
3. I just code a **simple** example.

# How to use
First, make sure you are under macOS system. Use command `cd` to the root directory of `realtime-fetch`. Then do like below,
1. `glide init`
2. `glide update`

Please make sure you can over the **wall**.

3. `brew install rethinkdb`
4. `rethinkdb`

Enter `localhost:8080` in your browser, you will see,

![WechatIMG150](http://ww2.sinaimg.cn/large/006tNc79ly1fez760h0frj31hu0y4q8r.jpg)

RethinkDB have a client driver itself, you can see the real-time read/write data on the board. Want more? [RethinkDB](https://www.rethinkdb.com/)⬅️

Click `Data Explorer` then execute sql as below,


5. `r.db("test").tableCreate("test")`
6. `r.db("test").table("test").insert({
      "name": "xiaoyusilen"
    })`

Do all of these, you can command as `go run main.go`, then you will see:

>   time="2017-04-25T18:02:47+08:00" level=info msg="rethinkdb connect success!" 
>
>   time="2017-04-25T18:02:47+08:00" level=info msg="{39628f3c-8943-4bcf-a70f-98290de97580 }" 
>
>   time="2017-04-25T18:02:57+08:00" level=info msg="{39628f3c-8943-4bcf-a70f-98290de97580 }"

The msg will be the id of your msg.

# Build Docker image

First of all, update the address in `route.rethinkdb.go` to your IP.
Then start rethink with your IP,
1. `rethinkdb --canonical-address your IP --bind all` 
2. `GOOS=linux GOARCH=amd64 go build main.go`
3. `docker build -t realtime-fetch .`
4. `docker images`

> ➜  realtime-fetch git:(develop) ✗ docker images
> 
>  REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
> 
>  realtime-fetch      latest              71a4028e8e4d        15 seconds ago      711 MB

5. `docker run -it realtime-fetch /bin/sh`

# TBD
- Add func `SendToAPI()` ✅
  - Get data from channel then send a request to a Restful API ✅
- Add time judge in `FetchRealData` to filter unchanged data at time interval ✅
- Add graceful shutdown maybe

# Nonsense
2017-4-28 16:32 Database will be changed to MongoDB. Why? My boss said it. 😑