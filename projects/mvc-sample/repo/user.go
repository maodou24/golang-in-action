package repo

import (
	"mvc-sample/model"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

var userRepo *UserRepo

func InitUserRepo(db *gorm.DB) {
	userRepo = &UserRepo{db: db}
}

func GetUserRepo() *UserRepo {
	return userRepo
}

func (r *UserRepo) FindByName(name string) (*model.User, error) {
	var user model.User
	err := r.db.Where("name = ?", name).First(&user).Error
	return &user, err
}

func (r *UserRepo) FindAll() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	return users, err
}
