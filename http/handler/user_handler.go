package handler

import (
	"fmt"
	"net/http"
	"simple_go/database/model"
	"simple_go/services/gin_pgx_err"
	"simple_go/services/password"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email,min=10,max=60"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}

func (h *Handler) CreateUser(ctx *gin.Context) {
	var req CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin_pgx_err.ErrorResponse(err))
		return
	}
	hashedPassword, err := password.Hash(req.Password)
	if err != nil {
		log.Error().Err(err).Msg("password hashing failed")
		err = fmt.Errorf("password hashing failed")
		ctx.JSON(http.StatusInternalServerError, gin_pgx_err.ErrorResponse(err))
		return
	}

	arg := model.CreateUserParams{
		Email:          req.Email,
		HashedPassword: hashedPassword,
	}
	user, err := h.db.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin_pgx_err.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, user)

}
