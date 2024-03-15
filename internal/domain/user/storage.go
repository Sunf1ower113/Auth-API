package user

type StorageUser interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id uint64) (*User, error)
	CreateUser(u *User) error
	UpdateUser(u *User) error
	GetUserPasswordByEmail(email string) (u *AuthDTO, err error)
	// CreateUser GetUserByToken(token string) (*User, error)
}
