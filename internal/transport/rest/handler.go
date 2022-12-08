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
	}

	api := router.Group("/api")
	{

		wallet := api.Group("/wallet")
		{
			wallet.POST("/", h.createWallet)
			wallet.GET("/", h.getAllWallets)
			wallet.GET("/:id", h.getWalletById)
			wallet.PUT("/:id", h.updateWallet)
			wallet.DELETE("/:id", h.deleteWallet)
		}

		transaction := api.Group("/transaction")
		{
			transaction.POST("/", h.createTransaction)
			transaction.GET("/", h.getAllTransactions)
			transaction.GET("/:id", h.getTransactionById)
			transaction.PUT("/:id", h.updateTransaction)
			transaction.DELETE("/:id", h.deleteTransaction)
		}
	}

	return router
}
