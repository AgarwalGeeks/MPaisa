package api

import (
	"database/sql"
	"net/http"
	"time"

	db "github.com/AgarwalGeeks/MPaisa/db/sqlc"
	"github.com/gin-gonic/gin"
)

type addCreditCardRequest struct {
	BankName    string `json:"bank_name" binding:"required"`
	CardName    string `json:"card_name" binding:"required"`
	CardNumber  string `json:"card_number" binding:"required"`
	Cvv         int32  `json:"cvv" binding:"required"`
	Pin         int32  `json:"pin" binding:"required"`
	ExpiaryDate string `json:"expiary_date" binding:"required"` // Changed to string for easier handling
	Usage       string `json:"usage"`                           // Optional, so no binding:"required"
	UserID      string `json:"user_id" binding:"required"`
}

func (server *Server) addCreditCard(ctx *gin.Context) {
	var req addCreditCardRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	expiaryDate, err := time.Parse("2006-01-02", req.ExpiaryDate) // Parse the date string
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.AddCreditCardParams{
		BankName:    req.BankName,
		CardName:    req.CardName,
		CardNumber:  req.CardNumber,
		Cvv:         req.Cvv,
		Pin:         req.Pin,
		ExpiaryDate: expiaryDate,
		UserID:      req.UserID,
	}

	// Handle optional Usage field
	if req.Usage != "" {
		arg.Usage = sql.NullString{
			String: req.Usage,
			Valid:  true,
		}
	}

	creditCard, err := server.store.AddCreditCard(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, creditCard)
}

type getAllCreditCardsRequest struct {
	UserID string `uri:"user_id" binding:"required,min=1"`
}

func (server *Server) getAllCreditCards(ctx *gin.Context) {
	var req getAllCreditCardsRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	creditCards, err := server.store.GetAllCreditCards(ctx, req.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, creditCards)
}

type updateCreditCardRequest struct {
	Usage      string `json:"usage"` // Changed to string
	CardNumber string `json:"card_number"`
	UserID     string `json:"user_id" binding:"required"`
}

func (server *Server) updateCreditCardUsage(ctx *gin.Context) {
	var req updateCreditCardRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateCreditCardUsageParams{
		UserID:     req.UserID,
		CardNumber: req.CardNumber,
	}

	// Handle optional Usage field
	if req.Usage != "" {
		arg.Usage = sql.NullString{
			String: req.Usage,
			Valid:  true,
		}
	} else {
		arg.Usage = sql.NullString{
			Valid: false,
		}
	}

	err := server.store.UpdateCreditCardUsage(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "updated"})
}

type deleteCreditCardRequest struct {
	CardNumber string `json:"card_number" binding:"required"`
	UserID     string `json:"user_id" binding:"required"`
}

func (server *Server) deleteCreditCard(ctx *gin.Context) {
	var req deleteCreditCardRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.DeleteCreditCardParams{
		UserID:     req.UserID,
		CardNumber: req.CardNumber,
	}

	err := server.store.DeleteCreditCard(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

type getCreditCardByCardNumberRequest struct {
	CardNumber string `json:"card_number" binding:"required"`
	UserID     string `json:"user_id" binding:"required"`
}

func (server *Server) getCreditCardByCardNumber(ctx *gin.Context) {
	var req getCreditCardByCardNumberRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetCreditCardByCardNumberParams{
		UserID:     req.UserID,
		CardNumber: req.CardNumber,
	}

	creditCard, err := server.store.GetCreditCardByCardNumber(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, creditCard)
}

type getCreditCardByUsageRequest struct {
	Usage  string `json:"usage"` // Changed to string
	UserID string `json:"user_id" binding:"required"`
}

func (server *Server) getCreditCardByUsage(ctx *gin.Context) {
	var req getCreditCardByUsageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetCreditCardByUsageParams{
		UserID: req.UserID,
	}

	// Handle optional Usage field
	if req.Usage != "" {
		arg.Usage = sql.NullString{
			String: req.Usage,
			Valid:  true,
		}
	} else {
		arg.Usage = sql.NullString{
			Valid: false,
		}
	}

	creditCard, err := server.store.GetCreditCardByUsage(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, creditCard)
}

type updateCreditCardDetailsRequest struct {
	CardName    string `json:"card_name" binding:"required"`
	CardNumber  string `json:"card_number" binding:"required"`
	Cvv         int32  `json:"cvv" binding:"required"`
	Pin         int32  `json:"pin" binding:"required"`
	ExpiaryDate string `json:"expiary_date" binding:"required"` // Changed to string for easier handling
	UserID      string `json:"user_id" binding:"required"`
}

func (server *Server) updateCreditCardDetails(ctx *gin.Context) {
	var req updateCreditCardDetailsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	expiaryDate, err := time.Parse("2006-01-02", req.ExpiaryDate) // Parse the date string
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateCreditCardDetailsParams{
		UserID:      req.UserID,
		CardName:    req.CardName,
		CardNumber:  req.CardNumber,
		Cvv:         req.Cvv,
		Pin:         req.Pin,
		ExpiaryDate: expiaryDate,
	}

	err = server.store.UpdateCreditCardDetails(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "updated"})
}

type updateCreditCardPinRequest struct {
	CardNumber string `json:"card_number" binding:"required"`
	Pin        int32  `json:"pin" binding:"required"`
	UserID     string `json:"user_id" binding:"required"`
}

func (server *Server) updateCreditCardPin(ctx *gin.Context) {
	var req updateCreditCardPinRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateCreditCardPinParams{
		UserID:     req.UserID,
		CardNumber: req.CardNumber,
		Pin:        req.Pin,
	}

	err := server.store.UpdateCreditCardPin(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "updated"})
}
