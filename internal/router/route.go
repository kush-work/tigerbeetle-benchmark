package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kushthedude/tigerbeetle-benchmark/internal/handlers"
	"net/http"
)

type Router struct {
	port   int
	server *gin.Engine
}

func NewRouter(port int) *Router {
	route := gin.Default()
	route.POST("/account", handlers.CreateAccount)
	route.GET("/account", handlers.GetAccount)
	route.POST("/transaction", handlers.CreateTransaction)
	route.GET("/transaction", handlers.GetTransaction)

	return &Router{
		server: route,
		port:   port,
	}
}

func (r *Router) RunServer() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", r.port), r.server)
}
