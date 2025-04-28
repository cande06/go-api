package api

import (
	"errors"
	"go-api/internal/sale"
	"go-api/internal/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	saleService *sale.Service
	userService *user.Service
}

// **************   USERS   *******************

// handleCreate maneja POST /users
func (h *handler) handleCreateUser(ctx *gin.Context) {
	// request payload = datos
	var req struct {
		Name     string `json:"name"`
		Address  string `json:"address"`
		NickName string `json:"nickname"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := &user.User{
		Name:     req.Name,
		Address:  req.Address,
		NickName: req.NickName,
	}
	if err := h.userService.Create(u); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, u)
}

// handleRead maneja GET /users/:id
func (h *handler) handleReadUser(ctx *gin.Context) {
	id := ctx.Param("id")

	u, err := h.userService.Get(id)
	if err != nil {
		if errors.Is(err, user.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, u)
}

// handleUpdate maneja PUT /users/:id
func (h *handler) handleUpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")

	// bind partial update fields
	var fields *user.UpdateFields
	if err := ctx.ShouldBindJSON(&fields); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := h.userService.Update(id, fields)
	if err != nil {
		if errors.Is(err, user.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, u)
}

// handleDelete maneja DELETE /users/:id
func (h *handler) handleDeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.userService.Delete(id); err != nil {
		if errors.Is(err, user.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// **************   SALES   *******************

// handleCreate maneja POST /sales
func (h *handler) handleCreateSale(ctx *gin.Context) {
	// request payload = solicitar datos
	var req struct {
		User_id string  `json:"user_id"`
		Amount  float32 `json:"amount"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s := &sale.Sale{
		User_id: req.User_id,
		Amount:  req.Amount,
	}
	if err := h.saleService.Create(s); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, s)
}

// handleRead maneja GET /sales?querystring
func (h *handler) handleReadSale(ctx *gin.Context) {
	id := ctx.Query("user_id")
	status := ctx.Query("status")

	u, err := h.userService.Get(id)
	if err != nil {
		if errors.Is(err, sale.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": u.ID, "status requested": status})
}

// handleUpdate maneja PUT /sales/:id
func (h *handler) handleUpdateSale(ctx *gin.Context) {
	id := ctx.Param("id")

	// bind partial update fields
	var fields *sale.UpdateFields
	if err := ctx.ShouldBindJSON(&fields); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s, err := h.saleService.Update(id, fields)
	if err != nil {
		if errors.Is(err, sale.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, s)
}

// handleDelete maneja DELETE /sales/:id
// func (h *handler) handleDeleteSale(ctx *gin.Context) {
// 	id := ctx.Param("id")

// 	if err := h.saleService.Delete(id); err != nil {
// 		if errors.Is(err, ErrNotFound) {
// 			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
// 			return
// 		}

// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	ctx.Status(http.StatusNoContent)
// }
