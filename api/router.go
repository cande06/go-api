package api

import (
	"go-api/internal/sale"
	"go-api/internal/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitRoutes(e *gin.Engine) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	userStorage := user.NewLocalStorage()
	saleStorage := sale.NewLocalStorage()

	//sale service recibe dos localStorage para comprobar que el una compra le pertenece a un usuario
	userService := user.NewService(userStorage, logger)
	saleService := sale.NewService(saleStorage, userStorage, logger)

	h := handler{
		userService: userService,
		saleService: saleService,
		logger: logger,
	}

	//****** RUTAS PARA USER ********
	e.POST("/users", h.handleCreateUser)
	e.GET("/users/:id", h.handleReadUser)
	e.PATCH("/users/:id", h.handleUpdateUser)
	e.DELETE("/users/:id", h.handleDeleteUser)

	//****** RUTAS PARA SALE ********
	e.POST("/sales", h.handleCreateSale)
	e.GET("/sales", h.handleReadSale)
	e.PATCH("/sales/:id", h.handleUpdateSale)
	// e.DELETE("/users/:id", h.handleDeleteSale)

	e.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
