package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type wsPayload struct {
	Action string `json:"action"`
	Message string `json:"message"`
}

type wsResponse struct {
	Action string `json:"action"`
	Message string `json:"message"`
}

type webConnection struct {
  conn *websocket.Conn
}

type Socket struct {
	ws *websocket.Conn
}

var wsChan = make(chan wsPayload)

var clients = make(map[webConnection]string)

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
	

	clients[webConnection{socket.ws}] = ""

	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	for {
		var payload wsPayload

		
		err := socket.ws.ReadJSON(&payload)

		if err != nil {
			
			return
		}

		wsChan <- payload
	}
}


func  ListenForWsChan() {

	response := wsResponse{}
	

	for{
		 e := <- wsChan
        response.Action = fmt.Sprintf(`hello from server %v`,e)
		response.Message = "hello from server"
		log.Println(e)
        broadcast(&response)
	}
}

func broadcast(response *wsResponse) {
	for client := range clients {
		err := client.conn.WriteJSON(response)

		if err != nil {
			delete(clients, client)
			return
		}
	}
}