package handler

import (
    "github.com/afv1/sims-tree/pkg/response"
    "github.com/gin-gonic/gin"
)

func (h *Handler) getApiInfo(ctx *gin.Context) {
    info := "Api description..."

    response.Success(ctx, map[string]string{"description": info})
}