package Config

import (
	"webservice-pattern/Controllers"
	"webservice-pattern/Repositories"
	"webservice-pattern/Services"
)

type Initialization struct {
	userRepo Repositories.UserRepository
	userSvc  Services.UserService
	UserCtrl Controllers.UserController
	RoleRepo Repositories.RoleRepository
}

func NewInitialization(userRepo Repositories.UserRepository,
	userService Services.UserService,
	userCtrl Controllers.UserController,
	roleRepo Repositories.RoleRepository) *Initialization {
	return &Initialization{
		userRepo: userRepo,
		userSvc:  userService,
		UserCtrl: userCtrl,
		RoleRepo: roleRepo,
	}
}
