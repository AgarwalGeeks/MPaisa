package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	db "github.com/AgarwalGeeks/MPaisa/db/sqlc"
	"github.com/gin-gonic/gin"
)

type addSalarySplitRequest struct {
	UserID             string  `json:"user_id" binding:"required"`
	TotalSalary        float64 `json:"total_salary" binding:"required,gt=0"`
	Month              string  `json:"month" binding:"required"`
	Notes              string  `json:"notes"`
	IsFullyTransferred bool    `json:"is_fully_transferred"`
}

type addSalarySplitItemRequest struct {
	SplitID       int64   `json:"split_id" binding:"required"`
	CategoryName  string  `json:"category_name" binding:"required"`
	Amount        float64 `json:"amount" binding:"required,gt=0"`
	MoveTo        string  `json:"move_to"`
	IsTransferred bool    `json:"is_transferred"`
}

func (server *Server) addSalarySplit(ctx *gin.Context) {
	var req addSalarySplitRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	month, err := time.Parse("2006-01-02", req.Month)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.AddSalarySplitParams{
		UserID:             req.UserID,
		TotalSalary:        fmt.Sprintf("%.2f", req.TotalSalary),
		Month:              month,
		Notes:              sql.NullString{String: req.Notes, Valid: req.Notes != ""},
		IsFullyTransferred: sql.NullBool{Bool: req.IsFullyTransferred, Valid: true},
	}

	salarySplit, err := server.store.AddSalarySplit(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, salarySplit)
}

func (server *Server) getSalarySplitById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	salarySplit, err := server.store.GetSalarySplitById(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, salarySplit)
}

func (server *Server) deleteSalarySplit(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = server.store.DeleteSalarySplitById(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "salary split deleted successfully"})
}

func (server *Server) addSalarySplitItem(ctx *gin.Context) {
	var req addSalarySplitItemRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.AddSalarySplitItemParams{
		SplitID:       int32(req.SplitID),
		CategoryName:  req.CategoryName,
		Amount:        fmt.Sprintf("%.2f", req.Amount),
		MoveTo:        sql.NullString{String: req.MoveTo, Valid: req.MoveTo != ""},
		IsTransferred: sql.NullBool{Bool: req.IsTransferred, Valid: true},
	}

	salarySplitItem, err := server.store.AddSalarySplitItem(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, salarySplitItem)
}

func (server *Server) getSalarySplitItemById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	salarySplitItem, err := server.store.GetSalarySplitItemsBySplitId(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, salarySplitItem)
}

func (server *Server) deleteSalarySplitItem(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = server.store.DeleteSalarySplitItemsBySplitId(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "salary split item deleted successfully"})
}

// Add handler for GetLatestSalarySplitByUserId
func (server *Server) getLatestSalarySplitByUserId(ctx *gin.Context) {
	userID := ctx.Param("user_id")

	salarySplit, err := server.store.GetLatestSalarySplitByUserId(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, salarySplit)
}

// Add handler for MarkSalarySplitAsFullyTransferredById
func (server *Server) markSalarySplitAsFullyTransferredById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = server.store.MarkSalarySplitAsFullyTransferredById(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Salary split marked as fully transferred"})
}

// Add handler for UpDateSalarySplitTotalById
func (server *Server) updateSalarySplitTotalById(ctx *gin.Context) {
	var req struct {
		ID          int32 `json:"id" binding:"required"`
		TotalSalary int32 `json:"total_salary" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpDateSalarySplitTotalByIdParams{
		ID:          req.ID,
		TotalSalary: fmt.Sprintf("%d", req.TotalSalary),
	}

	err := server.store.UpDateSalarySplitTotalById(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Salary split total updated successfully"})
}

// Add handler for MarkSalarySplitItemAsTransferredById
func (server *Server) markSalarySplitItemAsTransferredById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = server.store.MarkSalarySplitItemAsTransferredById(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Salary split item marked as transferred"})
}

// Add handler for UpdateSalarySplitItemAmountById
func (server *Server) updateSalarySplitItemAmountById(ctx *gin.Context) {
	var req struct {
		ID     int32  `json:"id" binding:"required"`
		Amount string `json:"amount" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateSalarySplitItemAmountByIdParams{
		ID:     req.ID,
		Amount: req.Amount,
	}

	err := server.store.UpdateSalarySplitItemAmountById(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Salary split item amount updated successfully"})
}
