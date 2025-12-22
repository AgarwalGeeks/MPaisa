package api

import (
	db "github.com/AgarwalGeeks/MPaisa/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{
		store: store,
	}
	router := gin.Default()

	router.POST("/credit_card", server.addCreditCard)
	router.GET("/credit_cards/:user_id", server.getAllCreditCards)
	router.GET("/credit_cards/usage", server.getCreditCardByUsage)
	router.GET("/credit_cards/number", server.getCreditCardByCardNumber)
	router.POST("/credit_card/replace", server.updateCreditCardDetails)
	router.DELETE("/credit_card", server.deleteCreditCard)
	router.POST("/credit_card/usage", server.updateCreditCardUsage)
	router.POST("/credit_card/pin", server.updateCreditCardPin)

	router.POST("/users", server.addUser)
	router.GET("/users/:id", server.getUserById)
	router.GET("/users/email", server.getUserByEmail)
	router.DELETE("/users", server.deleteUserByEmail)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"Error": err.Error()}
}
