package api

import (
	db "github.com/AgarwalGeeks/MPaisa/db/sqlc"
	"github.com/AgarwalGeeks/MPaisa/middleware"
	"github.com/AgarwalGeeks/MPaisa/tokens"
	"github.com/AgarwalGeeks/MPaisa/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store      *db.Store
	router     *gin.Engine
	tokenMaker tokens.Maker
}

func NewServer(store *db.Store) *Server {
	server := &Server{
		store: store,
	}

	server.setTokenMaker()
	router := gin.Default()

	router.POST("/users", server.addUser)
	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/").Use(middleware.AuthMiddleware(server.tokenMaker))
	server.addUserRequestsToRouter(authRoutes.(*gin.RouterGroup))
	server.addCreditCardRequestsToRouter(authRoutes.(*gin.RouterGroup))

	server.router = router

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) addCreditCardRequestsToRouter(authRoutes *gin.RouterGroup) {
	authRoutes.POST("/credit_card", server.addCreditCard)
	authRoutes.GET("/credit_cards/:user_id", server.getAllCreditCards)
	authRoutes.GET("/credit_cards/usage", server.getCreditCardByUsage)
	authRoutes.GET("/credit_cards/number", server.getCreditCardByCardNumber)
	authRoutes.POST("/credit_card/replace", server.updateCreditCardDetails)
	authRoutes.DELETE("/credit_card", server.deleteCreditCard)
	authRoutes.POST("/credit_card/usage", server.updateCreditCardUsage)
	authRoutes.POST("/credit_card/pin", server.updateCreditCardPin)
}

func (server *Server) addUserRequestsToRouter(authRoutes *gin.RouterGroup) {
	authRoutes.GET("/users/:id", server.getUserById)
	authRoutes.GET("/users/email", server.getUserByEmail)
	authRoutes.DELETE("/users", server.deleteUserByEmail)
}

func (server *Server) setTokenMaker() {
	tokenKey, err := util.RandomString(32)
	if err != nil {
		panic(err)
	}
	server.tokenMaker = tokens.NewPasetoMaker(tokenKey)
}

func errorResponse(err error) gin.H {
	return gin.H{"Error": err.Error()}
}
