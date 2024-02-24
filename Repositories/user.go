package Repositories

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"webservice-pattern/Entities"
)

type UserRepository interface {
	FindAllUser() ([]Entities.User, error)
	FindUserById(id int) (Entities.User, error)
	Save(user *Entities.User) (Entities.User, error)
	DeleteUserById(id int) error
}

type UserRepositoryImplement struct {
	db *gorm.DB
}

func (u UserRepositoryImplement) FindAllUser() ([]Entities.User, error) {
	var users []Entities.User

	var err = u.db.Preload("Role").Find(&users).Error
	if err != nil {
		log.Error("Got an error finding all couples. Error: ", err)
		return nil, err
	}

	return users, nil
}

func (u UserRepositoryImplement) FindUserById(id int) (Entities.User, error) {
	user := Entities.User{
		ID: id,
	}
	err := u.db.Preload("Role").First(&user).Error
	if err != nil {
		log.Error("Got and error when find user by id. Error: ", err)
		return Entities.User{}, err
	}
	return user, nil
}

func (u UserRepositoryImplement) Save(user *Entities.User) (Entities.User, error) {
	var err = u.db.Save(user).Error
	if err != nil {
		log.Error("Got an error when save user. Error: ", err)
		return Entities.User{}, err
	}
	return *user, nil
}

func (u UserRepositoryImplement) DeleteUserById(id int) error {
	err := u.db.Delete(&Entities.User{}, id).Error
	if err != nil {
		log.Error("Got an error when delete user. Error: ", err)
		return err
	}
	return nil
}

func UserRepositoryInit(db *gorm.DB) *UserRepositoryImplement {
	err := db.AutoMigrate(&Entities.User{})
	if err != nil {
		return nil
	}
	return &UserRepositoryImplement{
		db: db,
	}
}
