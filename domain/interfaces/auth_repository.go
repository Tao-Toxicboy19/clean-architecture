package interfaces

type IAuthRepository interface {
	Signin(username, password string) (LoginResponse, error)
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
