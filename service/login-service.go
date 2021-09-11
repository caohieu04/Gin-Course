package service

type LoginService interface {
	Login(username string, password string) bool
}

type loginService struct {
	authorizedUsername string
	authorizedPassword string
}

func NewLoginService() LoginService {
	return &loginService{
		authorizedUsername: "azura",
		authorizedPassword: "123",
	}
}

func (service *loginService) Login(username string, password string) bool {
	return service.authorizedPassword == username &&
		service.authorizedUsername == password
}
