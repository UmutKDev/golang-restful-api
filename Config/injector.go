// go:build wireinject
//go:build wireinject
// +build wireinject

package Config

import (
	"github.com/google/wire"
	"webservice-pattern/Controllers"
	"webservice-pattern/Repositories"
	"webservice-pattern/Services"
)

var db = wire.NewSet(ConnectToDB)

var userServiceSet = wire.NewSet(Services.UserServiceInit,
	wire.Bind(new(Services.UserService), new(*Services.UserServiceImplement)),
)

var userRepoSet = wire.NewSet(Repositories.UserRepositoryInit,
	wire.Bind(new(Repositories.UserRepository), new(*Repositories.UserRepositoryImplement)),
)

var userCtrlSet = wire.NewSet(Controllers.UserControllerInit,
	wire.Bind(new(Controllers.UserController), new(*Controllers.UserControllerImplement)),
)

var roleRepoSet = wire.NewSet(Repositories.RoleRepositoryInit,
	wire.Bind(new(Repositories.RoleRepository), new(*Repositories.RoleRepositoryImplement)),
)

func Init() *Initialization {
	wire.Build(NewInitialization, db, userCtrlSet, userServiceSet, userRepoSet, roleRepoSet)
	return nil
}
