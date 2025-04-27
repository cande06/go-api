package sale

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// rutas para Sale
func RegisterRoutes(e *gin.Engine) {
	saleStorage := NewLocalStorage()
	service := NewService(saleStorage)

	h := handler{
		saleService: service,
	}

	e.POST("/sales", h.handleCreate)
	e.GET("/sales", h.handleRead)
	e.PATCH("/sales/:id", h.handleUpdate)
	// e.DELETE("/users/:id", h.handleDelete)

	e.GET("/sales/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
