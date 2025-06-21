package service

import (
	"log"

	"github.com/siddheshRajendraNimbalkar/Go-Ecommerce/BACKEND/internal/domain"
)

type UserService struct {
}

func (s UserService) findUserByEmail(email string) (*domain.User, error) {

	return nil, nil
}

func (s UserService) Signup(input any) (string, error) {
	log.Println(input)

	return "this is my token", nil
}

func (s UserService) Login(input any) (string, error) {

	return "", nil
}

func (s UserService) GetVerificationCode(e domain.User) (int, error) {

	return 0, nil
}

func (s UserService) VerifyCode(id uint, code int) error {

	return nil
}

func (s UserService) CreateProfile(id uint, input any) error {

	return nil
}

func (s UserService) GetProfile(id uint) (*domain.User, error) {

	return nil, nil
}

func (s UserService) UpdateProfile(id uint, input any) error {

	return nil
}

func (s UserService) BecomeSeller(id uint, input any) (string, error) {

	return "", nil
}

func (s UserService) FindCart(id uint) ([]interface{}, error) {

	return nil, nil
}

func (s UserService) CreateCart(id uint, input any) ([]interface{}, error) {

	return nil, nil
}

func (s UserService) CreateOrder(u domain.User) (int, error) {

	return 0, nil
}

func (s UserService) GetOrder(u domain.User) ([]interface{}, error) {

	return nil, nil
}

func (s UserService) GetOrderById(id uint, uId uint) (interface{}, error) {

	return nil, nil
}
