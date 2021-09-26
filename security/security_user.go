package security

type LoginUser struct {
	Id       int64
	Phone    string
	Username string
}

func NewLoginUser(id int64, phone string, username string) *LoginUser {
	return &LoginUser{Id: id, Phone: phone, Username: username}
}
