package response

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type ErrorMsg struct {
    Err string `json:"error"`
}

func getStatus(status int, code ...any) int {
    if len(code) == 1 {
        st, ok := code[0].(int)

        if ok {
            status = st
        }
    }

    return status
}

func Error(ctx *gin.Context, err ErrorMsg, code ...any) {
    ctx.AbortWithStatusJSON(getStatus(http.StatusInternalServerError, code), err)
}

func Success(ctx *gin.Context, data map[string]string, code ...any) {
    data["success"] = "true"

    ctx.JSON(getStatus(http.StatusOK, code), data)
}
