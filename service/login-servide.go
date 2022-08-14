package service

// LoginService ...
type LoginService interface {
	Login(username string, password string) bool
}

type loginService struct {
	authorizedUsername string
	authorizedPassword string
}

// NewLoginService ...
func NewLoginService() LoginService {
	return &loginService{
		authorizedUsername: "evilnis",
		authorizedPassword: "reviews",
	}
}

func (service *loginService) Login(username string, password string) bool {
	return service.authorizedUsername == username &&
		service.authorizedPassword == password
}
