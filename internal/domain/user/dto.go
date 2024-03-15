package user

type CreateUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UpdateUserDTO struct {
	ID          uint64 `json:"user_id"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	BirthDate   string `json:"birth_date"`
}

type LoginResponseDTO struct {
	Token string `json:"token"`
}

type AuthDTO struct {
	ID             uint64 `json:"user_id"`
	HashedPassword string `json:"hashed_password"`
}
