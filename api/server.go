package api

import (
	"github.com/gin-gonic/gin"
	simpleBankDB "simpleBank/db/sqlc"
)

type Server struct {
	store  simpleBankDB.Store
	router *gin.Engine
}

func NewServer(store simpleBankDB.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/account", server.createAccount)
	router.GET("account/:id", server.getAccount)
	router.GET("/account", server.listAccount)
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"err": err.Error()}
}
