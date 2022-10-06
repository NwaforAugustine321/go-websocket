package websocket

import (
	"fmt"
	"log"
	"net/http"

	"Go-websocket-project/v2/interfaces/ws"

	"github.com/gorilla/websocket"
)

type wsPayload struct {
	Action   string `json:"action"`
	Message  string `json:"message"`
	Username string `json:"username"`
}

type wsResponse struct {
	Action   string `json:"action"`
	Message  string `json:"message"`
	Username string `json:"username"`
	UserList
}

type UserList struct {
	Users []string `json:"users"`
}

type webConnection struct {
	conn *websocket.Conn
}

type Socket struct {
	ws *websocket.Conn
}

var wsChan = make(chan wsPayload)

var clients = make(map[webConnection]string)

func NewWebsocket() ws.ISocket {
	return &Socket{}
}

func (socket *Socket) Connection(response http.ResponseWriter, request *http.Request) error {
	var upGrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upGrader.Upgrade(response, request, nil)

	if err != nil {
		log.Println(err)
		return err
	}

	socket.ws = conn
	return nil
}

func (socket *Socket) ListenForWs() {

	defer func() {
		if err := recover(); err != nil {
			log.Println("Recovering from error", err)
		}
	}()

	for {
		var payload = wsPayload{Action: "", Message: ""}

		err := socket.ws.ReadJSON(&payload)

		if err != nil {
			log.Println("here", err)
			return
		}

		wsChan <- payload
	}
}

func (socket *Socket) ListenForWsChan() {

	response := wsResponse{}
	users := make([]string, 0)

	for {

		e := <-wsChan

		switch e.Action {
		case "connect":
			response.Action = e.Action
			response.Message = "Connected Successfully"
			broadcast(&response)
		case "message":
			response.Action = e.Action
			response.Message = e.Message
			response.Username = e.Username
			broadcast(&response)
		case "login":
			if e.Message != clients[webConnection{socket.ws}] {
				clients[webConnection{socket.ws}] = e.Message
				users = append(users, e.Message)

				response.UserList = UserList{Users: users}
				response.Action = e.Action
				response.Message = fmt.Sprintf("%s has joined the chat", e.Message)
				broadcast(&response)
			}
		}
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
