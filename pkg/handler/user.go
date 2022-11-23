package handler

import (
    "fmt"
    "github.com/afv1/sims-tree/pkg/dto"
    "github.com/afv1/sims-tree/pkg/response"
    "github.com/gin-gonic/gin"
)

func (h *Handler) Register(ctx *gin.Context) {
    var data dto.AuthUser
    if err := ctx.BindJSON(&data); err != nil {
        response.Error(ctx, response.ErrorMsg{Err: err.Error()})
        return
    }

    user, err := h.services.Auth.Register(data)
    if err != nil {
        response.Error(ctx, response.ErrorMsg{Err: err.Error()})
        return
    }

    msg := map[string]string{
        "success": "true",
        "login": user.Login,
    }

    response.Success(ctx, msg)
}

func (h *Handler) Login(ctx *gin.Context) {
    var data dto.AuthUser
    if err := ctx.BindJSON(&data); err != nil {
        response.Error(ctx, response.ErrorMsg{Err: err.Error()})
        return
    }

    user, err := h.services.Auth.Login(data)
    if err != nil {
        response.Error(ctx, response.ErrorMsg{Err: err.Error()})
        return
    }

    msg := map[string]string{
        "success": "true",
        "id": fmt.Sprint(user.ID),
    }

    response.Success(ctx, msg)
}