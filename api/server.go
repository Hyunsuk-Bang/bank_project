package api

import (
	db "bank/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store    // allow interacting with DataBase
	router *gin.Engine // Router, send each api to correct handler
}

//create new server instance
func NewServer(store db.Store) *Server {
	router := gin.Default()
	server := &Server{
		store:  store,
		router: router,
	}

	router.POST("/accounts", server.CreateAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
