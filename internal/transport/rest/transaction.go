package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/skorolevskiy/wallet-backend-go/internal/domain"
	"net/http"
	"strconv"
)

//	@Summary		Create transaction
//	@Security		ApiKeyAuth
//	@Tags			transactions
//	@Description	create transaction
//	@ID				create-transaction
//	@Accept			json
//	@Produce		json
//	@Param			input		body		domain.Transaction	true	"wallet info"
//	@Param			walletId	path		int					true	"Wallet ID"
//	@Success		200			{integer}	integer				1
//	@Failure		400,404		{object}	errorResponse
//	@Failure		500			{object}	errorResponse
//	@Failure		default		{object}	errorResponse
//	@Router			/api/wallet/{walletId}/transaction [post]
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

type getAllTransactionResponse struct {
	Data []domain.Transaction `json:"data"`
}

//	@Summary		Get All Transactions
//	@Security		ApiKeyAuth
//	@Tags			transactions
//	@Description	get all transactions
//	@ID				get-all-transactions
//	@Accept			json
//	@Produce		json
//	@Param			walletId	path		int	true	"Wallet ID"
//	@Success		200			{object}	getAllTransactionResponse
//	@Failure		400,404		{object}	errorResponse
//	@Failure		500			{object}	errorResponse
//	@Failure		default		{object}	errorResponse
//	@Router			/api/wallet/{walletId}/transaction [get]
func (h *Handler) getAllTransactions(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	walletId, err := strconv.Atoi(c.Param("wallet_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid wallet id param")
		return
	}

	transactions, err := h.services.GetAllTransactions(userId, int64(walletId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllTransactionResponse{
		Data: transactions,
	})
}

//	@Summary		Get transaction By ID
//	@Security		ApiKeyAuth
//	@Tags			transactions
//	@Description	get transaction by id
//	@ID				get-transaction-by-id
//	@Accept			json
//	@Produce		json
//	@Param			walletId		path		int	true	"Wallet ID"
//	@Param			transactionId	path		int	true	"Transaction ID"
//	@Success		200				{object}	domain.Transaction
//	@Failure		400,404			{object}	errorResponse
//	@Failure		500				{object}	errorResponse
//	@Failure		default			{object}	errorResponse
//	@Router			/api/wallet/{walletId}/transaction/{transactionId} [get]
func (h *Handler) getTransactionById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	walletId, err := strconv.Atoi(c.Param("wallet_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid wallet id param")
		return
	}

	transactionId, err := strconv.Atoi(c.Param("transaction_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid transaction id param")
		return
	}

	transaction, err := h.services.GetTransactionById(userId, int64(walletId), int64(transactionId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, transaction)
}

//	@Summary		Update Transaction
//	@Security		ApiKeyAuth
//	@Tags			transactions
//	@Description	update transaction
//	@ID				update-transaction
//	@Accept			json
//	@Produce		json
//	@Param			input			body		domain.UpdateTransactionInput	true	"wallet updated info"
//	@Param			walletId		path		int								true	"Wallet ID"
//	@Param			transactionId	path		int								true	"Transaction ID"
//	@Success		200				{object}	string							"ok"
//	@Failure		400,404			{object}	errorResponse
//	@Failure		500				{object}	errorResponse
//	@Failure		default			{object}	errorResponse
//	@Router			/api/wallet/{walletId}/transaction/{transactionId} [put]
func (h *Handler) updateTransaction(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	walletId, err := strconv.Atoi(c.Param("wallet_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid wallet id param")
		return
	}

	transactionId, err := strconv.Atoi(c.Param("transaction_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid transaction id param")
		return
	}

	var input domain.UpdateTransactionInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.UpdateTransaction(userId, int64(walletId), int64(transactionId), input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{"ok"})
}

//	@Summary		Delete Transaction
//	@Security		ApiKeyAuth
//	@Tags			transactions
//	@Description	delete transaction
//	@ID				delete-transaction
//	@Accept			json
//	@Produce		json
//	@Param			walletId		path		int		true	"Wallet ID"
//	@Param			transactionId	path		int		true	"Transaction ID"
//	@Success		200				{object}	string	"ok"
//	@Failure		400,404			{object}	errorResponse
//	@Failure		500				{object}	errorResponse
//	@Failure		default			{object}	errorResponse
//	@Router			/api/wallet/{walletId}/transaction/{transactionId} [delete]
func (h *Handler) deleteTransaction(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	walletId, err := strconv.Atoi(c.Param("wallet_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid wallet id param")
		return
	}

	transactionId, err := strconv.Atoi(c.Param("transaction_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid transaction id param")
		return
	}

	err = h.services.DeleteTransaction(userId, int64(walletId), int64(transactionId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{"ok"})
}
