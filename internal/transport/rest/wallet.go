package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/skorolevskiy/wallet-backend-go/internal/domain"
	"net/http"
	"strconv"
)

// @Summary		Create wallet
// @Security		ApiKeyAuth
// @Tags			wallets
// @Description	create wallet
// @ID				create-wallet
// @Accept			json
// @Produce		json
// @Param			input	body		domain.Wallet	true	"wallet info"
// @Success		200		{integer}	integer			1
// @Failure		400,404	{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/api/wallet [post]
func (h *Handler) createWallet(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input domain.Wallet
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateWallet(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllWalletsResponse struct {
	Data []domain.Wallet `json:"data"`
}

// @Summary		Get All Wallets
// @Security		ApiKeyAuth
// @Tags			wallets
// @Description	get all wallets
// @ID				get-all-wallets
// @Accept			json
// @Produce		json
// @Success		200		{object}	getAllWalletsResponse
// @Failure		400,404	{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/api/wallet [get]
func (h *Handler) getAllWallets(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	wallets, err := h.services.GetAllWallets(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllWalletsResponse{
		Data: wallets,
	})
}

// @Summary		Get Wallet By ID
// @Security		ApiKeyAuth
// @Tags			wallets
// @Description	get wallet by id
// @ID				get-wallet-by-id
// @Accept			json
// @Produce		json
// @Param        id   path      int  true  "Wallet ID"
// @Success		200		{object}	domain.Wallet
// @Failure		400,404	{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/api/wallet/{id} [get]
func (h *Handler) getWalletById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	walletId, err := strconv.Atoi(c.Param("wallet_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	wallet, err := h.services.GetWalletById(userId, int64(walletId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, wallet)
}

// @Summary		Update Wallet
// @Security		ApiKeyAuth
// @Tags			wallets
// @Description	update wallet
// @ID				update-wallet
// @Accept			json
// @Produce		json
// @Param			input	body		domain.UpdateWalletInput	true	"wallet updated info"
// @Param        id   path      int  true  "Wallet ID"
// @Success		200		{object}	string "ok"
// @Failure		400,404	{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/api/wallet/{id} [put]
func (h *Handler) updateWallet(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	walletId, err := strconv.Atoi(c.Param("wallet_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input domain.UpdateWalletInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.UpdateWallet(userId, int64(walletId), input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{"ok"})
}

// @Summary		Delete Wallet
// @Security		ApiKeyAuth
// @Tags			wallets
// @Description	delete wallet
// @ID				delete-wallet
// @Accept			json
// @Produce		json
// @Param        id   path      int  true  "Wallet ID"
// @Success		200		{object}	string "ok"
// @Failure		400,404	{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/api/wallet/{id} [delete]
func (h *Handler) deleteWallet(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	walletId, err := strconv.Atoi(c.Param("wallet_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.DeleteWallet(userId, int64(walletId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}
