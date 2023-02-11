package system

// User represents a user in the system.
type User struct {
	Username string
	Password string
}

// UserRepository represents the user repository.
type UserRepository struct {
	users map[string]*User
}

// NewUserRepository returns a new instance of the user repository.
func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[string]*User),
	}
}

// AddUser adds a user to the repository.
func (ur *UserRepository) AddUser(username, password string) {
	ur.users[username] = &User{Username: username, Password: password}
}

// AuthorizeUser authorizes a user in the repository.
func (ur *UserRepository) AuthorizeUser(username, password string) bool {
	user, ok := ur.users[username]
	if !ok {
		return false
	}
	return user.Password == password
}
