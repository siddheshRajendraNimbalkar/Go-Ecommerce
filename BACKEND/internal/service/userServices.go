package service

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/siddheshRajendraNimbalkar/Go-Ecommerce/BACKEND/configs"
	"github.com/siddheshRajendraNimbalkar/Go-Ecommerce/BACKEND/internal/domain"
	"github.com/siddheshRajendraNimbalkar/Go-Ecommerce/BACKEND/internal/dto"
	"github.com/siddheshRajendraNimbalkar/Go-Ecommerce/BACKEND/internal/helper"
	"github.com/siddheshRajendraNimbalkar/Go-Ecommerce/BACKEND/internal/repository"
)

type UserService struct {
	Repo   repository.UserRepository
	Auth   helper.Auth
	Config configs.Config
}

func (s UserService) findUserByEmail(email string) (*domain.User, error) {

	user, err := s.Repo.FindUser(email)
	if err != nil {
		log.Println("[ERROR IN SERVICE] Error while finding user by email:", err)
		return nil, fmt.Errorf("error while finding user by email: %w", err)
	}

	return &user, nil
}

func (s UserService) Signup(input dto.UserSignup) (string, error) {

	hPassword, err := s.Auth.CreateHashPassword(input.Password)
	if err != nil {
		return "", fmt.Errorf("error while hashing password: %w", err)
	}

	user, err := s.Repo.CreateUser(domain.User{
		Email:    input.Email,
		Password: hPassword,
		Phone:    input.Phone,
	})

	if err != nil {
		return "", fmt.Errorf("error while creating user: %w", err)
	}

	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
}

func (s UserService) Login(email string, password string) (string, error) {

	user, err := s.findUserByEmail(email)

	if err != nil {
		log.Println("[ERROR IN SERVICE] Error while finding user by email:", err)
		return "", fmt.Errorf("error while finding user by email: %w", err)
	}

	isVeriy, err := s.Auth.VerifyPassword(password, user.Password)

	if err != nil || !isVeriy {
		return "", fmt.Errorf("error while verifying password: %w", err)
	}

	// Generate Token

	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
}

func (s UserService) isVeriyfied(id uint) bool {

	user, err := s.Repo.FindUserById(id)

	return err == nil && user.Verified
}

func (s UserService) GetVerificationCode(e domain.User) (int, error) {

	if s.isVeriyfied(e.ID) {
		return 0, errors.New("user already verified")
	}

	code, err := s.Auth.GetCode()

	if err != nil {
		return 0, fmt.Errorf("%v", err)
	}

	user := domain.User{
		Code:   fmt.Sprintf("%d", code),
		Expire: time.Now().Add(30 * time.Minute),
	}

	_, err = s.Repo.UpdateUser(e.ID, user)

	if err != nil {
		return 0, errors.New("unable to update verification code")
	}

	// mes := fmt.Sprintf("Verification code: %v", code)
	// notification := notification.NewNotificationClient(s.Config)
	// notification.SendSMS(udatedUser.Phone, mes)

	return code, nil
}

func (s UserService) VerifyCode(id uint, code int) error {

	if s.isVeriyfied(id) {
		return errors.New("user already verified")
	}

	user, err := s.Repo.FindUserById(id)

	if err != nil {
		return errors.New("error while finding user")
	}
	// Convert user.Code (string) to int for comparison
	userCodeInt, err := strconv.Atoi(user.Code)
	if err != nil {
		return errors.New("invalid code format")
	}
	if code != userCodeInt {
		return errors.New("verification code does not match")
	}

	if !time.Now().Before(user.Expire) {
		return errors.New("verification code expire")
	}

	updateUser := domain.User{
		Verified: true,
	}

	_, err = s.Repo.UpdateUser(id, updateUser)
	if err != nil {
		return errors.New("unable to verify user")
	}

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

func (s UserService) BecomeSeller(id uint, input dto.SellerInput) (string, error) {

	// Check if user is already a seller

	user, err := s.Repo.FindUserById(id)
	if err != nil {
		log.Println("[ERROR IN SERVICE] Error while finding user by ID:", err)
		return "", fmt.Errorf("error while finding user by ID: %w", err)
	}
	if user.UserType == domain.SELLER {
		return "", errors.New("user is already a seller")
	}

	// Update user type to seller
	updateUser := domain.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		UserType:  domain.SELLER,
	}

	updateUser, err = s.Repo.UpdateUser(id, updateUser)
	if err != nil {
		log.Println("[ERROR IN SERVICE] Error while updating user to seller:", err)
		return "", fmt.Errorf("error while updating user to seller: %w", err)
	}

	// generate token for seller
	token, err := s.Auth.GenerateToken(updateUser.ID, updateUser.Email, updateUser.UserType)
	if err != nil {
		log.Println("[ERROR IN SERVICE] Error while generating token for seller:", err)
		return "", fmt.Errorf("error while generating token for seller: %w", err)
	}

	// Create bank account for seller
	bankAccount := domain.BankAccount{
		BankAccount: uint(input.BankAccountNumber),
		SwiftCode:   input.SwiftCode,
		PaymentType: input.PaymetType,
		UserID:      updateUser.ID,
	}

	err = s.Repo.CreateBankAccount(bankAccount)
	if err != nil {
		log.Println("[ERROR IN SERVICE] Error while creating bank account for seller:", err)
		return "", fmt.Errorf("error while creating bank account for seller: %w", err)
	}

	return token, nil
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
