package handler

import (
    "github.com/afv1/sims-tree/pkg/response"
    "github.com/gin-gonic/gin"
)

// appMiddleware apply APP middleware for request
func (h *Handler) appMiddleware(ctx *gin.Context) {
}

// authMiddleware allow requests for authenticated users only
func (h *Handler) authMiddleware(ctx *gin.Context) {
    block, exists := ctx.Get("block")

    if exists && block == "true" {
        response.Error(ctx, response.ErrorMsg{Err: "Oops"})
        return
    }
}