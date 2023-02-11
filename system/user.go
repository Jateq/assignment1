package system

type User struct {
	Username string
	Password string
}

func (u *User) Register(username, password string) {
	u.Username = username
	u.Password = password
}

func (u *User) Authorize(username, password string) bool {
	return u.Username == username && u.Password == password
}

func (u *User) ChangePassword(oldPassword, newPassword string) bool {
	if u.Password == oldPassword {
		u.Password = newPassword
		return true
	}
	return false
}
