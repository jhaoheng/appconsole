package module

type ILogin interface {
	Check() bool
}

type Login struct {
	Username string
	Password string
}

func NewLogin(username string, password string) ILogin {
	return &Login{
		Username: username,
		Password: password,
	}
}

func (l *Login) Check() bool {
	return true
}
