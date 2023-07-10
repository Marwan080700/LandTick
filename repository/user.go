package repository

import (
	"Backend/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository{
	return &repository{db}
}

// find user--------------------------------------------------
func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error

	return users, err
}

// get a user--------------------------------------------------
func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	err := r.db.First(&user,ID).Error

	return user, err
}