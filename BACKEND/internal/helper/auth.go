package helper

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/siddheshRajendraNimbalkar/Go-Ecommerce/BACKEND/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Secret string
}

func NewAuth(secret string) Auth {
	if secret == "" {
		return Auth{}
	}
	return Auth{
		Secret: secret,
	}
}

func (a Auth) CreateHashPassword(password string) (string, error) {

	if len(password) < 6 {
		return "", errors.New("password must be at least 6 characters long")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", errors.New("failed to hash password")
	}

	return string(hashedPassword), nil
}

func (a Auth) VerifyPassword(password string, hashedPassword string) (bool, error) {

	if len(password) < 6 {
		return false, errors.New("password must be at least 6 characters long")
	}

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		return false, errors.New("password does not match")
	}

	return true, nil
}

func (a Auth) GenerateToken(id uint, email string, role string) (string, error) {

	if id == 0 || email == "" || role == "" {
		return "", errors.New("invalid user details for token generation")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"email":   email,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(a.Secret))

	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}

	return tokenString, nil
}

func (a Auth) VerifyToken(token string) (domain.User, error) {
	tokenArr := strings.Split(token, " ")

	if len(tokenArr) != 2 || strings.ToLower(tokenArr[0]) != "bearer" {
		return domain.User{}, errors.New("invalid token format")
	}

	tokenStr := tokenArr[1]

	parsedToken, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(a.Secret), nil
	})

	if err != nil || !parsedToken.Valid {
		return domain.User{}, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if ok || parsedToken.Valid {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return domain.User{}, errors.New("token has expired")
		}

		user := domain.User{}
		user.ID = uint(claims["user_id"].(float64))
		user.Email = claims["email"].(string)
		user.UserType = claims["role"].(string)

		return user, nil
	}

	return domain.User{}, errors.New("token verification failed")
}

func (a Auth) Authorized(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")

	if authHeader == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
			"error":   "Authorization header is missing",
		})
	}

	user, err := a.VerifyToken(authHeader)

	if err == nil && user.ID > 0 {
		ctx.Locals("user", user)
		return ctx.Next()
	}

	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "Unauthorized",
		"error":   err.Error(),
	})
}

func (a Auth) GetCurrentUser(ctx *fiber.Ctx) domain.User {

	user := ctx.Locals("user")

	return user.(domain.User)

}

func (a Auth) GetCode() (int, error) {

	return RandomNumbers(6)
}
