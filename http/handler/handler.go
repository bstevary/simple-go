package handler

import (
	"net/http"
	"simple_go/database/db"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	db db.Store
}

func NewHandler(store db.Store) *Handler {
	return &Handler{db: store}
}

func (h *Handler) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
