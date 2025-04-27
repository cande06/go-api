package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitRoutes registers all user CRUD endpoints on the given Gin engine.
// It initializes the storage, service, and handler, then binds each HTTP
// method and path to the appropriate handler function.

func RegisterRoutes(e *gin.Engine) {
	userStorage := NewLocalStorage()
	service := NewService(userStorage)

	h := handler{
		userService: service,
	}

	e.POST("/users", h.handleCreate)
	e.GET("/users/:id", h.handleRead)
	e.PATCH("/users/:id", h.handleUpdate)
	e.DELETE("/users/:id", h.handleDelete)

	e.GET("/users/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
