package api

import (
	"context"
	"log"
	"net/http"
	"time"

	db "github.com/AgarwalGeeks/MPaisa/db/sqlc"
	"github.com/AgarwalGeeks/MPaisa/middleware"
	"github.com/AgarwalGeeks/MPaisa/tokens"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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

	server.addPlatformRoutes(router)

	router.POST("/users", server.addUser)
	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/").Use(middleware.AuthMiddleware(server.tokenMaker))
	server.addUserRequestsToRouter(authRoutes.(*gin.RouterGroup))
	server.addCreditCardRequestsToRouter(authRoutes.(*gin.RouterGroup))
	server.addSalarySplitRequestsToRouter(authRoutes.(*gin.RouterGroup))
	server.addSalarySplitItemRequestsToRouter(authRoutes.(*gin.RouterGroup))

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

func (server *Server) addSalarySplitRequestsToRouter(authRoutes *gin.RouterGroup) {
	authRoutes.POST("/salary_splits", server.addSalarySplit)
	authRoutes.GET("/salary_splits/:id", server.getSalarySplitById)
	authRoutes.DELETE("/salary_splits/:id", server.deleteSalarySplit)
	authRoutes.GET("/salary_splits/latest/:user_id", server.getLatestSalarySplitByUserId)
	authRoutes.POST("/salary_splits/mark_transferred/:id", server.markSalarySplitAsFullyTransferredById)
	authRoutes.PUT("/salary_splits/update_total", server.updateSalarySplitTotalById)
}

func (server *Server) addSalarySplitItemRequestsToRouter(authRoutes *gin.RouterGroup) {
	authRoutes.POST("/salary_split_items", server.addSalarySplitItem)
	authRoutes.GET("/salary_split_items/:id", server.getSalarySplitItemById)
	authRoutes.DELETE("/salary_split_items/:id", server.deleteSalarySplitItem)
	authRoutes.POST("/salary_split_items/mark_transferred/:id", server.markSalarySplitItemAsTransferredById)
	authRoutes.PUT("/salary_split_items/update_amount", server.updateSalarySplitItemAmountById)
}

// addPlatformRoutes registers platform / infrastructure focused endpoints.
func (server *Server) addPlatformRoutes(router *gin.Engine) {
	router.GET("/healthz", server.healthCheck)
	router.GET("/readyz", server.readinessCheck)
}

func (server *Server) healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (server *Server) readinessCheck(ctx *gin.Context) {
	requestCtx, cancel := context.WithTimeout(ctx.Request.Context(), 2*time.Second)
	defer cancel()

	if err := server.store.Ping(requestCtx); err != nil {
		log.Printf("readiness check failed: %v", err)
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"status": "unhealthy"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "ready"})
}

func (server *Server) setTokenMaker() {
	tokenKey := viper.GetString("PASETO_SYMMETRIC_KEY")
	if len(tokenKey) < 32 {
		log.Fatal("PASETO_SYMMETRIC_KEY must be set and at least 32 characters long")
	}

	server.tokenMaker = tokens.NewPasetoMaker(tokenKey)
}

func errorResponse(err error) gin.H {
	return gin.H{"Error": err.Error()}
}
