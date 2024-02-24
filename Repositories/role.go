package Repositories

import (
	"gorm.io/gorm"
	"webservice-pattern/Entities"
)

type RoleRepository interface {
	FindAllRole()
}

type RoleRepositoryImplement struct {
	db *gorm.DB
}

func (r RoleRepositoryImplement) FindAllRole() {
	panic("implement me")
}

func RoleRepositoryInit(db *gorm.DB) *RoleRepositoryImplement {
	err := db.AutoMigrate(&Entities.Role{}, &Entities.User{})
	if err != nil {
		return nil
	}
	return &RoleRepositoryImplement{
		db: db,
	}
}
