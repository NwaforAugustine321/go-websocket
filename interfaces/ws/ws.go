package ws

import "net/http"

type ISocket interface {
	 Connection(path string, h http.HandlerFunc)
	 ListenForWs()
}
