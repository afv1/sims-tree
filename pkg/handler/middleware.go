package handler

import (
    "github.com/afv1/sims-tree/pkg/response"
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
)

// appMiddleware apply APP middleware for request
func (h *Handler) appMiddleware(ctx *gin.Context) {
    h.testMiddleware(ctx)
}

func (h *Handler) testMiddleware(ctx *gin.Context) {
    user, exists := ctx.Get("user")

    if exists {
        logrus.Info("Testing " + user.(string) + " for fun")
    } else {
        logrus.Info("Testing without user")
    }
}

// authMiddleware allow requests for authorized users only
func (h *Handler) authMiddleware(ctx *gin.Context) {
    block, exists := ctx.Get("block")

    if exists && block == "true" {
        response.Error(ctx, response.ErrorMsg{Err: "Oops"})
        return
    }
}