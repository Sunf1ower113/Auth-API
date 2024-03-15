package user

import (
	customError "auth-api/internal/error"
	"auth-api/internal/midlleware"
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

type ServiceUser interface {
	GetUserById(ctx context.Context, id uint64) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	CreateUser(ctx context.Context, dto *CreateUserDTO) error
	UpdateUser(ctx context.Context, dto *UpdateUserDTO) (*User, error)
	Login(ctx context.Context, dto *CreateUserDTO) (*LoginResponseDTO, error)
}

type serviceUser struct {
	storage StorageUser
}

func NewUserService(storage StorageUser) ServiceUser {
	return &serviceUser{
		storage: storage,
	}
}

func (s *serviceUser) GetUserById(ctx context.Context, id uint64) (*User, error) {
	return s.storage.GetUserById(id)
}

func (s *serviceUser) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	return s.storage.GetUserByEmail(email)
}

func (s *serviceUser) CreateUser(ctx context.Context, dto *CreateUserDTO) error {
	newUser, err := dtoCreateValidator(dto)
	if err != nil {
		return err
	}
	_, err = s.GetUserByEmail(ctx, newUser.Email)
	if err == nil {
		return errors.New("user already exist")
	}
	p, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u := &User{Email: dto.Email, HashedPassword: string(p)}
	if s.storage.CreateUser(u) != nil {
		return err
	}
	return nil
}

func (s *serviceUser) UpdateUser(ctx context.Context, dto *UpdateUserDTO) (*User, error) {
	existingUser, err := s.GetUserByEmail(ctx, dto.Email)
	if err == nil {
		if existingUser.ID != dto.ID {
			return nil, customError.BusyUpdateEmailError
		}
	}
	if err := userUpdateValidator(dto); err != nil {
		return nil, err
	}
	b := dto.Password
	if b != "" {
		p, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		} else {
			dto.Password = string(p)
		}
	}
	u := &User{
		ID:             dto.ID,
		Email:          dto.Email,
		Username:       dto.Username,
		HashedPassword: dto.Password,
		PhoneNumber:    dto.PhoneNumber,
		BirthDate:      dto.BirthDate,
		Role:           "",
	}
	err = s.storage.UpdateUser(u)
	if err != nil {
		return nil, err
	}
	//userUpdater(u, dto)
	return u, nil
}

func (s *serviceUser) Login(ctx context.Context, dto *CreateUserDTO) (*LoginResponseDTO, error) {
	u, err := s.getUserPasswordByEmail(ctx, dto.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}
	if checkPassword([]byte(u.HashedPassword), []byte(dto.Password)) != nil {
		return nil, errors.New("invalid email or password")
	}
	token, err := generateToken(u.ID)
	if err != nil {
		return nil, errors.New("internal server error")
	}
	return &LoginResponseDTO{Token: token}, nil
}

func (s *serviceUser) getUserPasswordByEmail(ctx context.Context, email string) (*AuthDTO, error) {
	u, err := s.storage.GetUserPasswordByEmail(email)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func dtoCreateValidator(dto *CreateUserDTO) (*User, error) {
	u := &User{}
	if dto.Email == "" || dto.Password == "" {
		return nil, errors.New("invalid registration data")
	}
	u.Email = dto.Email
	u.HashedPassword = dto.Password
	return u, nil
}

func userUpdateValidator(dto *UpdateUserDTO) error {
	if dto.Username == "" && dto.PhoneNumber == "" && dto.BirthDate == "" && dto.Password == "" && dto.Email == "" {
		return errors.New("nothing to update")
	}
	return nil
}

func userUpdater(u *User, dto *UpdateUserDTO) (count int) {
	if dto.Email != "" && u.Email != dto.Email {
		u.Email = dto.Email
		count++
	}
	if dto.Password != "" && u.HashedPassword != dto.Password {
		u.HashedPassword = dto.Password
		count++
	}
	if dto.BirthDate != "" && u.BirthDate != dto.BirthDate {
		u.BirthDate = dto.BirthDate
		count++
	}
	if dto.PhoneNumber != "" && u.PhoneNumber != dto.PhoneNumber {
		u.PhoneNumber = dto.PhoneNumber
		count++
	}
	if dto.Username != "" && u.Username != dto.Username {
		u.Username = dto.Username
		count++
	}
	return
}

func checkPassword(hashedPassword, password []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		return err
	}
	return nil
}

func generateToken(id uint64) (string, error) {
	claims := &midlleware.Claims{
		UserID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Устанавливаем срок действия токена на 24 часа
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//func withContext(ctx context.Context, fn func(context.Context) (*User, error)) (*User, error) {
//	ch := make(chan *User)
//	errCh := make(chan error)
//	go func() {
//		u, err := fn(ctx)
//		if err != nil {
//			errCh <- err
//			return
//		}
//		ch <- u
//	}()
//	select {
//	case <-ctx.Done():
//		return nil, ctx.Err()
//	case err := <-errCh:
//		return nil, err
//	case u := <-ch:
//		return u, nil
//	}
//}

//func (s *serviceUser) GetUserByToken(ctx context.Context, token string) (*User, error) {
//	return withContext(ctx, func(ctx context.Context) (*User, error) {
//		return s.storage.GetUserByToken(token)
//	})
//}

/*
func (s *serviceUser) GetUserById(ctx context.Context, id uint64) (*User, error) {
	return withContext(ctx, func(ctx context.Context) (*User, error) {
		return s.storage.GetUserById(id)
	})
}

func (s *serviceUser) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	return withContext(ctx, func(ctx context.Context) (*User, error) {
		return s.storage.GetUserByEmail(email)
	})
}

func (s *serviceUser) CreateUser(ctx context.Context, dto *CreateUserDTO) error {
	log.Println(dto)
	log.Println(*dto)
	log.Println(dto.Email, dto.Password)
	newUser, err := dtoCreateValidator(dto)
	if err != nil {
		return err
	}
	_, err = withContext(ctx, func(ctx context.Context) (*User, error) {
		_, err := s.GetUserByEmail(ctx, newUser.Email)
		if err == nil {
			return nil, errors.New("user already exist")
		}
		return newUser, nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceUser) UpdateUser(ctx context.Context, dto *UpdateUserDTO) (*User, error) {
	return withContext(ctx, func(ctx context.Context) (*User, error) {
		existingUser, err := s.GetUserByEmail(ctx, dto.Email)
		if err != nil {
			return nil, err
		}
		if err := userUpdateValidator(dto); err != nil {
			return nil, err
		}
		err = s.storage.UpdateUser(existingUser)
		if err != nil {
			return nil, err
		}
		userUpdater(existingUser, dto)
		return existingUser, nil
	})
}

*/
