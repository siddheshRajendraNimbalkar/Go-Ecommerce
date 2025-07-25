package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/siddheshRajendraNimbalkar/Go-Ecommerce/BACKEND/internal/api/rest"
	"github.com/siddheshRajendraNimbalkar/Go-Ecommerce/BACKEND/internal/dto"
	"github.com/siddheshRajendraNimbalkar/Go-Ecommerce/BACKEND/internal/repository"
	"github.com/siddheshRajendraNimbalkar/Go-Ecommerce/BACKEND/internal/service"
)

type UserRoute struct {
	svc service.UserService
}

func SetupUserRoutes(r *rest.RestRoutes) {
	app := r.App

	svc := service.UserService{
		Repo:   repository.NewUserRepository(r.DB),
		Auth:   r.Auth,
		Config: r.Config,
	}

	handler := &UserRoute{
		svc: svc,
	}

	pubRoutes := app.Group("/users")
	// unAuthenticated routes
	pubRoutes.Post("/register", handler.Register)
	pubRoutes.Post("/login", handler.Login)

	pvtRoutes := pubRoutes.Group("/", r.Auth.Authorized)

	// Authenticated routes
	pvtRoutes.Get("/verify", handler.getVerificationCode)
	pvtRoutes.Post("/verify", handler.verification)
	pvtRoutes.Post("/profile", handler.CreateProfile)
	pvtRoutes.Get("/profile", handler.GetProfile)

	pvtRoutes.Post("/cart", handler.AddToCart)
	pvtRoutes.Get("/cart", handler.GetCart)
	pvtRoutes.Get("/order", handler.GetOrders)
	pvtRoutes.Get("/orde/:id", handler.GetOrder)

	pvtRoutes.Post("/become-seller", handler.BecomeSeller)

}

func (u *UserRoute) Register(ctx *fiber.Ctx) error {

	user := dto.UserSignup{}
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input data",
		})
	}

	token, err := u.svc.Signup(user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"token":   token,
	})
}

func (u *UserRoute) Login(ctx *fiber.Ctx) error {

	LoginInput := dto.UserLogin{}
	err := ctx.BodyParser(&LoginInput)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input data",
		})
	}

	token, err := u.svc.Login(LoginInput.Email, LoginInput.Password)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error on login",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"token":   token,
	})
}

func (u *UserRoute) getVerificationCode(ctx *fiber.Ctx) error {

	user := u.svc.Auth.GetCurrentUser(ctx)

	// create a verification code to update the user profile in DB
	code, err := u.svc.GetVerificationCode(user)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while getting code",
			"error":   err,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "getVerificationCode",
		"code":    code,
	})
}

func (u *UserRoute) verification(ctx *fiber.Ctx) error {

	user := u.svc.Auth.GetCurrentUser(ctx)

	var req dto.VerifyCodeInput

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "please provide verification code",
		})
	}

	err := u.svc.VerifyCode(user.ID, req.Code)

	if err != nil {
		mess := fmt.Sprintf("%v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": mess,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "verification done successfuly",
	})
}

func (u *UserRoute) CreateProfile(ctx *fiber.Ctx) error {

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "CreateProfile",
	})
}

func (u *UserRoute) GetProfile(ctx *fiber.Ctx) error {

	user := u.svc.Auth.GetCurrentUser(ctx)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "GetProfile",
		"user":    user,
	})
}

func (u *UserRoute) AddToCart(ctx *fiber.Ctx) error {

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "AddToCart",
	})
}

func (u *UserRoute) GetCart(ctx *fiber.Ctx) error {

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "GetCart",
	})
}

func (u *UserRoute) CreateOrder(ctx *fiber.Ctx) error {

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "CreateOrder",
	})
}

func (u *UserRoute) GetOrders(ctx *fiber.Ctx) error {

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "GetOrders",
	})
}

func (u *UserRoute) GetOrder(ctx *fiber.Ctx) error {

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "GetOrder",
	})
}

func (u *UserRoute) BecomeSeller(ctx *fiber.Ctx) error {

	user := u.svc.Auth.GetCurrentUser(ctx)

	req := dto.SellerInput{}

	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(400).JSON(&fiber.Map{
			"message": "Invalid input data",
		})
	}

	token, err := u.svc.BecomeSeller(user.ID, req)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Be The Seller",
		"token":   token,
	})
}
