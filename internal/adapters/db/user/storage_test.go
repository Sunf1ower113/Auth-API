package user

import (
	"auth-api/internal/domain/user"
	"database/sql"
	"reflect"
	"testing"
)

func TestNewUserStorage(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
		want user.StorageUser
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserStorage(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateUpdateQuery(t *testing.T) {
	type args struct {
		toUpdateUser *user.User
		existedUser  *user.User
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 []interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := generateUpdateQuery(tt.args.toUpdateUser, tt.args.existedUser)
			if got != tt.want {
				t.Errorf("generateUpdateQuery() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("generateUpdateQuery() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_storageUser_CreateUser(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		u *user.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			su := &storageUser{
				db: tt.fields.db,
			}
			if err := su.CreateUser(tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_storageUser_GetUserByEmail(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *user.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			su := &storageUser{
				db: tt.fields.db,
			}
			got, err := su.GetUserByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_storageUser_GetUserById(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		id uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *user.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			su := &storageUser{
				db: tt.fields.db,
			}
			got, err := su.GetUserById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_storageUser_GetUserPasswordByEmail(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *user.AuthDTO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			su := &storageUser{
				db: tt.fields.db,
			}
			got, err := su.GetUserPasswordByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserPasswordByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserPasswordByEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_storageUser_GetUserPasswordById(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *user.AuthDTO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			su := &storageUser{
				db: tt.fields.db,
			}
			got, err := su.GetUserPasswordById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserPasswordById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserPasswordById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_storageUser_UpdateUser(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		u *user.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			su := &storageUser{
				db: tt.fields.db,
			}
			if err := su.UpdateUser(tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
