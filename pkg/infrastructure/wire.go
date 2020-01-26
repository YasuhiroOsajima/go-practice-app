//+build wireinject

// The build tag makes sure the stub is not built in the final build.
package infrastructure

import (
	"github.com/google/wire"

	"go-practice-app/pkg/infrastructure/sqlhandler"
	"go-practice-app/pkg/interfaces/controllers"
)

// InitializeEvent creates an Event. It will error if the Event is staffed with
// a grumpy greeter.
func InitializeEvent() *controllers.UserController {
	wire.Build(
		controllers.NewUserController,
		sqlhandler.NewSqlHandler,
	)
	return nil
}
