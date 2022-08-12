package main

import (
	"net/http"
	"websocket/imlsp"
	"websocket/redis"
)

func main() {

	err := redis.InitWithEnv()
	if err != nil {
		panic("缓存连接异常:" + err.Error())
	}
	http.HandleFunc("/ws", imlsp.WsHandler)
	http.ListenAndServe("0.0.0.0:7777", nil)
}


