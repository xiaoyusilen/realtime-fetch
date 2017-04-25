# Realtime-fetch
A simple project of fetching realtime and write to channel then send to one Restful API. 🤔

# When you can use this repo
If you :
- have a table of realtime data
- not too much data
- you want to deal with the realtime data

Congratulations! 🎉 You can use this repo.

# About this repo
1. I use channel as a message queue, of course you can use `RabbitMQ` or others.
2. I use `RethinkDB` as my database, you can use others if you like.But you should change the code yourself.
3. I code a **simple** example.

# How to use
First, you should have a MacBook. Use command `cd` to the root directory of `realtime-fetch`. Then do like below:
1. `glide init`
2. `glide update`

Please make sure you can over the **wall**.

3. `brew install rethinkdb`
4. `rethinkdb`

Enter `localhost:8080` in your browser
5. `r.db("test").tableCreate("test")`
6. `r.db("test").table("test").insert({
      "name": "xiaoyusilen"
    })`
    
Do all of these, you can command as `go run main.go`, you will see:

>   time="2017-04-25T18:02:47+08:00" level=info msg="rethinkdb connect success!" 
>   time="2017-04-25T18:02:47+08:00" level=info msg="{39628f3c-8943-4bcf-a70f-98290de97580 }" 
>   time="2017-04-25T18:02:57+08:00" level=info msg="{39628f3c-8943-4bcf-a70f-98290de97580 }"
  
The msg will be the id of your msg.

Thx~
