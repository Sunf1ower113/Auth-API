package user

import (
	"auth-api/internal/domain/user"
	customError "auth-api/internal/error"
	"database/sql"
	"errors"
	"fmt"
)

type storageUser struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) user.StorageUser {
	return &storageUser{
		db: db,
	}
}

func (su *storageUser) GetUserByEmail(email string) (*user.User, error) {
	u := &user.User{}
	q := `SELECT * FROM users WHERE email = ?`
	row := su.db.QueryRow(q, email)
	if err := row.Scan(&u.ID, &u.Email, &u.Username, &u.HashedPassword, &u.PhoneNumber, &u.BirthDate, &u.Role); err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, customError.NotFoundError
		}
		return nil, err
	}
	return u, nil
}

func (su *storageUser) GetUserById(id uint64) (*user.User, error) {
	u := &user.User{}
	q := `SELECT * FROM users WHERE users.user_id = ?`
	row := su.db.QueryRow(q, id)
	if err := row.Scan(&u.ID, &u.Email, &u.Username, &u.HashedPassword, &u.PhoneNumber, &u.BirthDate, &u.Role); err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, customError.NotFoundError
		}
		return nil, err
	}
	return u, nil
}

func (su *storageUser) GetUserPasswordById(id int64) (*user.AuthDTO, error) {
	u := &user.AuthDTO{}
	q := `SELECT user_id, password FROM users WHERE users.user_id = ?`
	row := su.db.QueryRow(q, id)
	if err := row.Scan(&u.ID, &u.HashedPassword); err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, customError.NotFoundError
		}
		return nil, err
	}
	return u, nil
}

func (su *storageUser) GetUserPasswordByEmail(email string) (*user.AuthDTO, error) {
	u := &user.AuthDTO{}
	q := `SELECT user_id, password FROM users WHERE users.email = ?`
	row := su.db.QueryRow(q, email)
	if err := row.Scan(&u.ID, &u.HashedPassword); err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, customError.NotFoundError
		}
		return nil, err
	}
	return u, nil
}

func (su *storageUser) CreateUser(u *user.User) error {
	q := `INSERT INTO users(email, password) values(?,?)`
	_, err := su.db.Exec(q, u.Email, u.HashedPassword)
	if err != nil {
		return err
	}
	return nil
}

func (su *storageUser) UpdateUser(u *user.User) error {
	existedUser, err := su.GetUserById(u.ID)
	if err != nil {
		return err
	}
	t := &user.User{
		ID:             existedUser.ID,
		Email:          existedUser.Email,
		Username:       existedUser.Username,
		HashedPassword: existedUser.HashedPassword,
		PhoneNumber:    existedUser.PhoneNumber,
		BirthDate:      existedUser.BirthDate,
		Role:           existedUser.Role,
	}
	q, updates := generateUpdateQuery(u, existedUser, t)
	if q == "" {
		return errors.New(customError.NothingToUpdateUserErrorMsg)
	}
	_, err = su.db.Exec(q, append(updates, u.ID)...)
	if err != nil {
		return err
	}
	u.Email = t.Email
	u.PhoneNumber = t.PhoneNumber
	u.Username = t.Username
	u.BirthDate = t.BirthDate
	u.HashedPassword = t.HashedPassword
	return nil
}

func generateUpdateQuery(toUpdateUser, existedUser, t *user.User) (string, []interface{}) {
	var updateQuery string
	var updates []interface{}

	if toUpdateUser.Email != "" && toUpdateUser.Email != existedUser.Email {
		updateQuery += "email=?, "
		updates = append(updates, toUpdateUser.Email)
		t.Email = toUpdateUser.Email
	}

	if toUpdateUser.Username != "" && toUpdateUser.Username != existedUser.Username {
		updateQuery += "username=?, "
		updates = append(updates, toUpdateUser.Username)
		t.Username = toUpdateUser.Username
	}

	if toUpdateUser.HashedPassword != "" && toUpdateUser.HashedPassword != existedUser.HashedPassword {
		updateQuery += "password=?, "
		updates = append(updates, toUpdateUser.HashedPassword)
		t.HashedPassword = toUpdateUser.HashedPassword
	}

	if toUpdateUser.PhoneNumber != "" && toUpdateUser.PhoneNumber != existedUser.PhoneNumber {
		updateQuery += "phone_number=?, "
		updates = append(updates, toUpdateUser.PhoneNumber)
		t.PhoneNumber = toUpdateUser.PhoneNumber
	}

	if toUpdateUser.BirthDate != "" && toUpdateUser.BirthDate != existedUser.BirthDate {
		updateQuery += "birth_date=?, "
		updates = append(updates, toUpdateUser.BirthDate)
		t.BirthDate = toUpdateUser.BirthDate
	}

	if len(updates) == 0 {
		return "", nil
	}

	updateQuery = updateQuery[:len(updateQuery)-2]

	query := fmt.Sprintf("UPDATE users SET %s WHERE user_id=?", updateQuery)
	return query, updates
}

//func (su *storageUser) GetUserByToken(token string) (u *user.User, err error) {
//	q := `SELECT * FROM users WHERE users.session_token = ?`
//	row := su.db.QueryRow(q, token)
//	if err = row.Scan(&u.ID, &u.Username, &u.Email, &u.BirthDate, &u.PhoneNumber, &u.Role, &u.HashedPassword); err != nil {
//		log.Printf("Error query:%s\n", err)
//		return
//	}
//	log.Printf("Succeed query%s\n", u.Username)
//	return
//}

//func (su *storageUser) UpdateUserToken(u user.User, token string) error {
//	u2, err := su.GetUserById(u.ID)
//	if err != nil {
//		return err
//	}
//	q := `UPDATE users SET session_token=? WHERE user_id=?`
//	_, err = su.db.Exec(q, token)
//	if err != nil {
//		log.Printf("Update error:%s\n", err)
//		return err
//	}
//	log.Printf("User %s has been update\n", u2.Email)
//	return nil
//}
