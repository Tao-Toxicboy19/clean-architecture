package repositories

import (
	"clean-architecture/domain/interfaces"
	"errors"
)

type AuthRepository struct{}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{}
}

func (repo *AuthRepository) Signin(username, password string) (interfaces.LoginResponse, error) {
	if username == "admin" && password == "password" {
		// หากสำเร็จ จะคืนค่า LoginResponse กลับมา
		return interfaces.LoginResponse{
			AccessToken:  "access_token_example",
			RefreshToken: "refresh_token_example",
		}, nil
	}

	// หาก username หรือ password ไม่ถูกต้อง ให้คืน error
	return interfaces.LoginResponse{}, errors.New("invalid username or password")
}
