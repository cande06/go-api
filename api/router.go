package api

import (
	"go-api/internal/sale"
	"go-api/internal/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(e *gin.Engine) {
	saleStorage := sale.NewLocalStorage()
	saleService := sale.NewService(saleStorage)

	userStorage := user.NewLocalStorage()
	userService := user.NewService(userStorage)

	h := handler{
		saleService: saleService,
		userService: userService,
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
