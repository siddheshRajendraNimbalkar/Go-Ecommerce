package repository

import (
	"errors"
	"log"

	"github.com/siddheshRajendraNimbalkar/Go-Ecommerce/BACKEND/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	CreateUser(user domain.User) (domain.User, error)
	FindUser(email string) (domain.User, error)
	FindUserById(id uint) (domain.User, error)
	UpdateUser(id uint, user domain.User) (domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r userRepository) CreateUser(usere domain.User) (domain.User, error) {

	err := r.db.Create(&usere).Error

	if err != nil {
		log.Println("[ERROR IN REPOSITORY] Error while creating user:", err)
		return domain.User{}, errors.New("error while creating user")
	}

	return usere, nil
}

func (r userRepository) FindUser(email string) (domain.User, error) {

	var user domain.User

	err := r.db.Find(&user, "email=?", email).Error
	if err != nil {
		log.Println("[ERROR IN REPOSITORY] Error while finding user:", err)
		return domain.User{}, errors.New("error while finding user")
	}

	return user, nil
}

func (r userRepository) FindUserById(id uint) (domain.User, error) {

	var user domain.User
	err := r.db.First(&user, id).Error
	if err != nil {
		log.Println("[ERROR IN REPOSITORY] Error while finding user by ID:", err)
		return domain.User{}, errors.New("error while finding user by ID")
	}

	return user, nil
}

func (r userRepository) UpdateUser(id uint, u domain.User) (domain.User, error) {

	var user domain.User

	err := r.db.Model(&user).Clauses(clause.Returning{}).Where("id=?", id).Updates(u).Error
	if err != nil {
		log.Println("[ERROR IN REPOSITORY] Error while updating user:", err)
		return domain.User{}, errors.New("error while updating user")
	}

	// Fetch the updated user
	err = r.db.First(&user, id).Error
	if err != nil {
		log.Println("[ERROR IN REPOSITORY] Error while fetching updated user:", err)
		return domain.User{}, errors.New("error while fetching updated user")
	}

	return user, nil
}
