package api

import (
	"net/http"

	db "github.com/AgarwalGeeks/MPaisa/db/sqlc"
	"github.com/AgarwalGeeks/MPaisa/util"
	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Email        string `json:"email" binding:"required,email"`
	UserPassword string `json:"user_password" binding:"required,min=6"`
	Username     string `json:"username" binding:"required,alphanum,lowercase"`
}

func (server *Server) addUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.UserPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Email:        req.Email,
		UserPassword: hashedPassword,
		Username:     req.Username,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := gin.H{
		"id":         user.ID,
		"created_at": user.CreatedAt,
	}

	ctx.JSON(http.StatusOK, response)
}

type getUserByIDRequest struct {
	ID string `uri:"id" binding:"required,min=16"`
}

func (server *Server) getUserById(ctx *gin.Context) {
	var req getUserByIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUserById(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type getUserByEmailRequest struct {
	Email string `json:"email" binding:"required,email"`
}

func (server *Server) getUserByEmail(ctx *gin.Context) {
	var req getUserByEmailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type deleteUserByEmailRequest struct {
	Email string `json:"email" binding:"required,email"`
}

func (server *Server) deleteUserByEmail(ctx *gin.Context) {
	var req deleteUserByEmailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, errorResponse(err))
		return
	}

	err := server.store.DeleteUserByEmail(ctx, req.Email)
	if err != nil {
		ctx.JSON(500, errorResponse(err))
		return
	}

	ctx.JSON(200, gin.H{"message": "user deleted successfully"})
}
