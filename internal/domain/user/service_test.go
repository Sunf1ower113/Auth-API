package user

import (
	"context"
	"reflect"
	"testing"
)

func TestNewUserService(t *testing.T) {
	type args struct {
		storage StorageUser
	}
	tests := []struct {
		name string
		args args
		want ServiceUser
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserService(tt.args.storage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkPassword(t *testing.T) {
	type args struct {
		hashedPassword []byte
		password       []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkPassword(tt.args.hashedPassword, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("checkPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_dtoCreateValidator(t *testing.T) {
	type args struct {
		dto *CreateUserDTO
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dtoCreateValidator(tt.args.dto)
			if (err != nil) != tt.wantErr {
				t.Errorf("dtoCreateValidator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dtoCreateValidator() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateToken(t *testing.T) {
	type args struct {
		id uint64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateToken(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("generateToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serviceUser_CreateUser(t *testing.T) {
	type fields struct {
		storage StorageUser
	}
	type args struct {
		ctx context.Context
		dto *CreateUserDTO
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
			s := &serviceUser{
				storage: tt.fields.storage,
			}
			if err := s.CreateUser(tt.args.ctx, tt.args.dto); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_serviceUser_GetUserByEmail(t *testing.T) {
	type fields struct {
		storage StorageUser
	}
	type args struct {
		ctx   context.Context
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &serviceUser{
				storage: tt.fields.storage,
			}
			got, err := s.GetUserByEmail(tt.args.ctx, tt.args.email)
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

func Test_serviceUser_GetUserById(t *testing.T) {
	type fields struct {
		storage StorageUser
	}
	type args struct {
		ctx context.Context
		id  uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &serviceUser{
				storage: tt.fields.storage,
			}
			got, err := s.GetUserById(tt.args.ctx, tt.args.id)
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

func Test_serviceUser_Login(t *testing.T) {
	type fields struct {
		storage StorageUser
	}
	type args struct {
		ctx context.Context
		dto *CreateUserDTO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *LoginResponseDTO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &serviceUser{
				storage: tt.fields.storage,
			}
			got, err := s.Login(tt.args.ctx, tt.args.dto)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Login() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serviceUser_UpdateUser(t *testing.T) {
	type fields struct {
		storage StorageUser
	}
	type args struct {
		ctx context.Context
		dto *UpdateUserDTO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &serviceUser{
				storage: tt.fields.storage,
			}
			got, err := s.UpdateUser(tt.args.ctx, tt.args.dto)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serviceUser_getUserPasswordByEmail(t *testing.T) {
	type fields struct {
		storage StorageUser
	}
	type args struct {
		ctx   context.Context
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *AuthDTO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &serviceUser{
				storage: tt.fields.storage,
			}
			got, err := s.getUserPasswordByEmail(tt.args.ctx, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("getUserPasswordByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getUserPasswordByEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUpdateValidator(t *testing.T) {
	type args struct {
		dto *UpdateUserDTO
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := userUpdateValidator(tt.args.dto); (err != nil) != tt.wantErr {
				t.Errorf("userUpdateValidator() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userUpdater(t *testing.T) {
	type args struct {
		u   *User
		dto *UpdateUserDTO
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := userUpdater(tt.args.u, tt.args.dto); gotCount != tt.wantCount {
				t.Errorf("userUpdater() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
