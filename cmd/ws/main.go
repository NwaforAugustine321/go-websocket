package main

import (
	"Go-websocket-project/v2/repositories/router"
	"net/http"

	"Go-websocket-project/v2/repositories/websocket"

	"github.com/julienschmidt/httprouter"
)

func main() {
	httpRoutes := router.NewHttpRouter()
	ws := websocket.NewWebsocket()

	go ws.ListenForWsChan()

	httpRoutes.Get("/", func(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
		response.Write([]byte("Hello"))
	})

	httpRoutes.Websocket("/ws", func(response http.ResponseWriter, request *http.Request, params httprouter.Params) {

		ws.Connection(response, request)
		go ws.ListenForWs()
	})

	httpRoutes.Serve(":4000")
}
