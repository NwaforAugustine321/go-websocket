package router

import "net/http"

type IRouter interface {
	Get(path string, h http.HandlerFunc)
  Websocket(path string, h http.HandlerFunc)
	Serve(port string)
}

