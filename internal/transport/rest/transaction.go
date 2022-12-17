package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/skorolevskiy/wallet-backend-go/internal/domain"
	"net/http"
	"strconv"
)

func (h *Handler) createTransaction(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	walletId, err := strconv.Atoi(c.Param("wallet_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid wallet id param")
		return
	}

	var input domain.Transaction
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateTransaction(userId, int64(walletId), input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllTransactions(c *gin.Context) {

}

func (h *Handler) getTransactionById(c *gin.Context) {

}

func (h *Handler) updateTransaction(c *gin.Context) {

}

func (h *Handler) deleteTransaction(c *gin.Context) {

}
