package main

import (
	"Go-websocket-project/v2/repositories/router"
	"net/http"

	"Go-websocket-project/v2/repositories/websocket"

	"github.com/julienschmidt/httprouter"
)

func main() {
	httpRoutes := router.NewHttpRouter()

	httpRoutes.Get("/", func(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
		response.Write([]byte("Hello"))
	})

	httpRoutes.Websocket("/ws", func(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
		ws, _ := websocket.NewWebsocket(response, request)

		ws.Connection().WriteJSON("Hello from websocket")

	})

	httpRoutes.Serve(":4000")
}
