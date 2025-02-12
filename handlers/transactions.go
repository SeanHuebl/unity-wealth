package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/seanhuebl/unity-wealth/helpers"
	"github.com/seanhuebl/unity-wealth/models"
)

func (h *Handler) NewTransaction(ctx *gin.Context) {
	userID, err := userIDFromClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	var req models.NewTransactionRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	txn, err := h.transactionService.CreateTransaction(ctx.Request.Context(), userID.String(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create transaction",
		})
		return
	}

	response := h.transactionService.ConvertToResponse(txn)

	ctx.JSON(http.StatusCreated, gin.H{
		"data": response,
	})

}

func (h *Handler) GetTransactionsByUserID(ctx *gin.Context) {
	userID, err := userIDFromClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	cursorDate := ctx.Query("cursor_date")
	cursorID := ctx.Query("cursor_id")
	pageSize := int64(50)

	transactions, nextCursorDate, nextCursorID, hasMoreData, err :=
		h.transactionService.ListUserTransactions(ctx.Request.Context(), userID, &cursorDate, &cursorID, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to get transactions",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{

			"transactions":     transactions,
			"next_cursor_date": nextCursorDate,
			"next_cursor_id":   nextCursorID,
			"has_more_data":    hasMoreData,
		},
	})
}

func (h *Handler) GetTransactionByID(ctx *gin.Context) {
	userID, err := userIDFromClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	id := ctx.Param("id")

	txn, err := h.transactionService.GetTransactionByID(ctx.Request.Context(), userID.String(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to get transaction",
		})
		return

	}

	response := h.transactionService.ConvertToResponse(txn)

	ctx.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

func (h *Handler) UpdateTransaction(ctx *gin.Context) {
	userID, err := userIDFromClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	var req models.NewTransactionRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	txn, err := h.transactionService.UpdateTransaction(ctx.Request.Context(), id, userID.String(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update transaction",
		})
		return
	}

	response := h.transactionService.ConvertToResponse(txn)

	ctx.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

func (h *Handler) DeleteTransaction(ctx *gin.Context) {
	userID, err := userIDFromClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	err = h.transactionService.DeleteTransaction(ctx.Request.Context(), id, userID.String())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error deleting transaction",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{

			"transaction_deleted": "success",
		},
	})
}

// Helpers
func userIDFromClaims(ctx *gin.Context) (uuid.UUID, error) {
	claims, err := helpers.ValidateClaims(ctx)
	if err != nil {
		return uuid.Nil, err
	}

	return uuid.Parse(claims.Subject)
}
