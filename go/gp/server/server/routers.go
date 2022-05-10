package server

import "net/http"



func (s *server) routers()  {
	s.router.Handle(http.MethodGet, "/hello-world")
}