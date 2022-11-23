package handler

import (
    "github.com/afv1/sims-tree/pkg/service"
    "github.com/gin-gonic/gin"
)

type Handler struct {
    services *service.Provider
}

func NewHandler(services *service.Provider) *Handler {
    return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
    router := gin.New()

    router.Group("/api")
    {
        auth := router.Group("/auth")
        {
            auth.POST("/register", h.Register)
            auth.POST("/login", h.Login)
        }

        user := router.Group("/user", h.authMiddleware)
        {
            user.POST("/logout")
            user.PUT("/update")
            user.DELETE("/delete")

            // TODO: Add handlers
        }

//        tree := router.Group("/tree")
//        {
//            tree.GET("/get/:id")
//            tree.POST("/create")
//            tree.PUT("/update/:id")
//            tree.DELETE("/delete/:id")
//
//            // TODO: Add handlers
//        }
    }

    return router
}
