//+build wireinject

package infrastructure

import (
	"github.com/google/wire"

	"go-practice-app/pkg/infrastructure/sqlhandler"
	"go-practice-app/pkg/interfaces/controllers"
)

func InitializeController() *controllers.UserController {
	wire.Build(
		controllers.NewUserController,
		sqlhandler.NewSqlHandler,
	)
	return nil
}
