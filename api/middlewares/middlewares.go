package middlewares

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewCorsMiddleware),
	fx.Provide(NewDatabaseTransactionMiddleware),
	fx.Provide(NewJWTAuthMiddleware),
	fx.Provide(NewMiddlewares),
)

type IMiddleware interface {
	SetUp()
}

type Middlewares []IMiddleware

func NewMiddlewares(
	corsMiddleware CorsMiddleware,
	databaseTransactionMiddleware DatabaseTransactionMiddleware,
) Middlewares {
	return Middlewares{
		corsMiddleware,
		databaseTransactionMiddleware,
	}
}

func (m Middlewares) SetUp() {
	for _, middleware := range m {
		middleware.SetUp()
	}
}
