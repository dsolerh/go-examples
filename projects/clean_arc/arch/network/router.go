package network

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var _ Router = (*router)(nil)

type router struct {
	engine *gin.Engine
}

// GetEngine implements Router.
func (r *router) GetEngine() *gin.Engine {
	return r.engine
}

// LoadControllers implements Router.
func (r *router) LoadControllers(controllers []Controller) {
	for _, c := range controllers {
		g := r.engine.Group(c.Path())
		c.MountRoutes(g)
	}
}

// LoadRootMiddlewares implements Router.
func (r *router) LoadRootMiddlewares(middlewares []RootMiddleware) {
	for _, m := range middlewares {
		m.Attach(r.engine)
	}
}

// RegisterValidationParsers implements Router.
func (r *router) RegisterValidationParsers(tagNameFunc validator.TagNameFunc) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(tagNameFunc)
	}
}

// Start implements Router.
func (r *router) Start(ip string, port uint16) {
	address := fmt.Sprintf("%s:%d", ip, port)
	r.engine.Run(address)
}

func NewRouter(mode string) Router {
	gin.SetMode(mode)
	eng := gin.Default()
	r := router{
		engine: eng,
	}
	return &r
}
