package api

import (
	"github.com/gin-gonic/gin"
	db "multitenant.com/app/db/sqlc"
)

//Server serves HTTP requests for user service
type Server struct {
	store  db.Store
	router *gin.Engine
}

//NewServer creates a new HTTP server and setup routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//add routes to the router
	router.POST("/users", server.createUser)
	router.GET("/users/:id", server.getUser)
	router.GET("/users", server.getUsers)

	server.router = router
	return server
}

//Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
