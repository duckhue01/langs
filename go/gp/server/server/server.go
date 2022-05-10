package server

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type (
	server struct {
		db     *sql.DB
		router *gin.Engine
	}
)


func (s *server) HanderHelloWorld()  {
	
}