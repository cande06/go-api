package sale

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// handler holds the sale service and implements HTTP handlers for sale CRUD.
type handler struct {
	saleService *Service
}

// handleCreate maneja POST /sales
func (h *handler) handleCreate(ctx *gin.Context) {
	// request payload = solicitar datos
	var req struct {
		User_id string  `json:"user_id"`
		Amount  float32 `json:"amount"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s := &Sale{
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
func (h *handler) handleRead(ctx *gin.Context) {
	id := ctx.Query("user_id")
	status := ctx.Query("status")

	s, err := h.saleService.Get(id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": s.User_id, "status requested": status})
}

// handleUpdate maneja PUT /sales/:id
func (h *handler) handleUpdate(ctx *gin.Context) {
	id := ctx.Param("id")

	// bind partial update fields
	var fields *UpdateFields
	if err := ctx.ShouldBindJSON(&fields); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s, err := h.saleService.Update(id, fields)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, s)
}

// handleDelete maneja DELETE /sales/:id
// func (h *handler) handleDelete(ctx *gin.Context) {
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
