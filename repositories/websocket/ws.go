package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type wsPayload struct {
}

type Socket struct {
	ws *websocket.Conn
}

var wsChan = make(chan wsPayload)

var clients = make(map[*Socket]string)

func NewWebsocket(response http.ResponseWriter, request *http.Request) (*Socket, error) {
	var upGrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	ws, err := upGrader.Upgrade(response, request, nil)

	if err != nil {
		log.Println(err)
		return &Socket{ws}, err
	}

	return &Socket{ws}, nil
}

func (socket *Socket) Connection() *websocket.Conn {
	return socket.ws
}

func (socket *Socket) ListenForWs() {
	ws := socket.Connection()

	clients[&Socket{ws}] = ""
}
