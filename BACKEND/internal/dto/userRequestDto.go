package dto

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignup struct {
	UserLogin
	Phone string `json:"phone"`
}

type VerifyCodeInput struct {
	Code int `json:"code"`
}

type SellerInput struct {
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	BankAccountNumber int    `json:"bank_account"`
	SwiftCode         string `json:"swift_code"`
	PaymetType        string `json:"payment_type"`
}
