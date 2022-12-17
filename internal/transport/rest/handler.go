package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/skorolevskiy/wallet-backend-go/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	user := router.Group("/auth")
	{
		user.POST("/sign-up", h.signUp)
		user.POST("/sign-in", h.signIn)
		user.GET("/refresh", h.refresh)
	}

	api := router.Group("/api", h.userIdentity)
	{

		wallet := api.Group("/wallet")
		{
			wallet.POST("/", h.createWallet)
			wallet.GET("/", h.getAllWallets)
			wallet.GET("/:wallet_id", h.getWalletById)
			wallet.PUT("/:wallet_id", h.updateWallet)
			wallet.DELETE("/:wallet_id", h.deleteWallet)

			transaction := wallet.Group(":wallet_id/transaction")
			{
				transaction.POST("/", h.createTransaction)
				transaction.GET("/", h.getAllTransactions)
				transaction.GET("/:transaction_id", h.getTransactionById)
				transaction.PUT("/:transaction_id", h.updateTransaction)
				transaction.DELETE("/:transaction_id", h.deleteTransaction)
			}
		}

	}

	return router
}
