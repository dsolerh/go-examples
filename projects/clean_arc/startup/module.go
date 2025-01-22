package startup

import (
	"clean_arc/api/auth"
	authMW "clean_arc/api/auth/middleware"
	"clean_arc/api/user"
	coreMW "clean_arc/arch/middleware"
	"clean_arc/arch/mongo"
	"clean_arc/arch/network"
	"clean_arc/arch/redis"
	"clean_arc/config"
	"context"
)

type Module network.Module[module]

var _ Module = (*module)(nil)

type module struct {
	Context     context.Context
	Env         *config.Env
	DB          mongo.Database
	Store       redis.Store
	UserService user.Service
	AuthService auth.Service
	// BlogService blog.Service
}

func (m *module) AuthenticationProvider() network.AuthenticationProvider {
	return authMW.NewAuthenticationProvider(m.AuthService, m.UserService)
}

func (m *module) AuthorizationProvider() network.AuthorizationProvider {
	return authMW.NewAuthorizationProvider()
}

func (m *module) Controllers() []network.Controller {
	return []network.Controller{
		auth.NewController(m.AuthenticationProvider(), m.AuthorizationProvider(), m.AuthService),
		user.NewController(m.AuthenticationProvider(), m.AuthorizationProvider(), m.UserService),
	}
}

func (m *module) RootMiddlewares() []network.RootMiddleware {
	return []network.RootMiddleware{
		coreMW.NewErrorCatcher(), // NOTE: this should be the first handler to be mounted
		authMW.NewKeyProtection(m.AuthService),
		coreMW.NewNotFound(),
	}
}

func (m *module) GetInstance() *module {
	return m
}

func NewModule(context context.Context, env *config.Env, db mongo.Database, store redis.Store) Module {
	userService := user.NewService(db)
	authService := auth.NewService(db, env, userService)
	// blogService := blog.NewService(db, store, userService)

	return &module{
		Context:     context,
		Env:         env,
		DB:          db,
		Store:       store,
		UserService: userService,
		AuthService: authService,
		// BlogService: blogService,
	}
}
