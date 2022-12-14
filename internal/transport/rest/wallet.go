package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/skorolevskiy/wallet-backend-go/internal/domain"
	"net/http"
)

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

func (h *Handler) getAllWallets(c *gin.Context) {

}

func (h *Handler) getWalletById(c *gin.Context) {

}

func (h *Handler) updateWallet(c *gin.Context) {

}

func (h *Handler) deleteWallet(c *gin.Context) {

}
