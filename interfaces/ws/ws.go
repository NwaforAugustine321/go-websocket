package ws

import (
	"net/http"
)

type ISocket interface {
	 Connection(response http.ResponseWriter, request *http.Request) error
	 ListenForWs()
	 ListenForWsChan()
}
