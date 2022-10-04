package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)



type Router struct {
  router *httprouter.Router
}


 func  NewHttpRouter() *Router {
   return &Router{httprouter.New()}
}


func (route *Router) Get(path string, handler func (response http.ResponseWriter, request *http.Request, params httprouter.Params)) {
  route.router.Handle(http.MethodGet, path, handler)
}


 func(route *Router) Websocket(path string, handler func (response http.ResponseWriter, request *http.Request, params httprouter.Params)){
   route.router.Handle(http.MethodGet,path, handler)
 }

func (route *Router) Serve(port string){
  app := &http.Server{  
    Addr:           port,
    Handler:        route.router,
    
  }

  app.ListenAndServe()
}