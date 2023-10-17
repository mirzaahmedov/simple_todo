package router

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mirzaahmedov/simple_todo/internal/store"
)

type HTTPRouter struct {
	router *gin.Engine
	store  store.Store
	logger *slog.Logger
}

func NewHTTPRouter(store store.Store, logger *slog.Logger) *HTTPRouter {
	return &HTTPRouter{
		store:  store,
		router: gin.Default(),
	}
}

func (r *HTTPRouter) Listen(httpAddress string) error {

	r.router.POST("/todos", r.handleTodoCreate)
	r.router.GET("/todos", r.handleTodoGetAll)
	r.router.GET("/todos/:id", r.handleTodoGetByID)
	r.router.PUT("/todos/:id", r.handleTodoUpdate)
	r.router.DELETE("/todos/:id", r.handleTodoDelete)

	return r.router.Run(httpAddress)
}
func (r *HTTPRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}
