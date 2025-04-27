package api

import (
	"go-api/internal/sale"
	"go-api/internal/user"

	"github.com/gin-gonic/gin"
)

// inicializa los routes
func InitRoutes(r *gin.Engine) {
	user.RegisterRoutes(r)
	sale.RegisterRoutes(r)
}
