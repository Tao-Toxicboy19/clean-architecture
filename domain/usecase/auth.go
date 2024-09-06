package usecase

import "clean-architecture/domain/interfaces"

type AuthUsecase struct {
	authRepo interfaces.IAuthRepository
}

func NewAuthUsecase(repo interfaces.IAuthRepository) *AuthUsecase {
	return &AuthUsecase{authRepo: repo}
}

func (uc *AuthUsecase) Signin(username, password string) (interfaces.LoginResponse, error) {
	return uc.authRepo.Signin(username, password)
}
