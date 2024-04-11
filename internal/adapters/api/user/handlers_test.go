package user

import (
	"auth-api/internal/domain/user"
	customError "auth-api/internal/error"
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MockUserService struct {
	err error
}

func (m MockUserService) UpdateUser(ctx context.Context, dto *user.UpdateUserDTO) (*user.User, error) {
	if dto.ID != 123 {
		return nil, customError.NotFoundError
	} else if dto.Email == "" && dto.PhoneNumber == "" && dto.BirthDate == "" && dto.Username == "" && dto.Password == "" {
		return nil, customError.NothingToUpdateError
	}
	return nil, nil
}

func (m MockUserService) Login(ctx context.Context, dto *user.CreateUserDTO) (*user.LoginResponseDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockUserService) CreateUser(ctx context.Context, dto *user.CreateUserDTO) error {
	if dto.Email == "existing@example.com" {
		return customError.BusyUpdateEmailError
	} else if dto.Email == "invalid@example.com" {
		return customError.CreateUserBadInputError
	}
	return nil
}

func Test_handler_CreateUser(t *testing.T) {
	type fields struct {
		userService user.ServiceUser
	}
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	type expected struct {
		code  int
		error string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected expected
	}{
		{
			name: "ValidRequest",
			fields: fields{
				userService: MockUserService{},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/register", strings.NewReader(`{"email":"test@example.com","password":"password123"}`)),
			},
			expected: expected{
				code:  201,
				error: "",
			},
		},
		{
			name: "ExistingEmail",
			fields: fields{
				userService: MockUserService{},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/register", strings.NewReader(`{"email":"existing@example.com", "password":"password123"}`)),
			},
			expected: expected{
				code:  400,
				error: "Email is busy",
			},
		},
		{
			name: "InvalidEmail",
			fields: fields{
				userService: MockUserService{},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/register", strings.NewReader(`{"email":"invalid@example.com", "password":"password123"}`)),
			},
			expected: expected{
				code:  400,
				error: "Invalid email or password",
			},
		},
		{
			name: "InvalidJSON",
			fields: fields{
				userService: MockUserService{},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/register", strings.NewReader(`{"email":"test@example.com"`)),
			},
			expected: expected{
				code:  400,
				error: "Invalid JSON syntax",
			},
		},
		{
			name: "InvalidDataType",
			fields: fields{
				userService: MockUserService{},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/register", strings.NewReader(`{"email":123, "password":"password123"}`)),
			},
			expected: expected{
				code:  400,
				error: "Invalid request data type",
			},
		},
		{
			name: "EmptyRequestBody",
			fields: fields{
				userService: MockUserService{},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/register", strings.NewReader(``)),
			},
			expected: expected{
				code:  400,
				error: "Failed to read request body",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			h := &handler{
				userService: tt.fields.userService,
			}
			h.CreateUser(tt.args.w, tt.args.r)
			resp := tt.args.w.Result()

			if resp.StatusCode != tt.expected.code {
				t.Errorf("Expected status code %d, got %d", tt.expected.code, resp.StatusCode)
			}

			if tt.expected.error != "" {
				body := make([]byte, 100)
				n, _ := resp.Body.Read(body)
				body = body[:n]
				if !strings.Contains(string(body), tt.expected.error) {
					t.Errorf("Expected error message '%s' found '%s'", tt.expected.error, string(body))
				}
			}
		})
	}
}

func Test_handler_UpdateUser(t *testing.T) {
	type fields struct {
		userService user.ServiceUser
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	type expected struct {
		code  int
		error string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected expected
	}{
		{
			name: "ValidRequest",
			fields: fields{
				userService: MockUserService{},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("PUT", "/settings", strings.NewReader(`{"user_id":123, "email":"e123@example.com"}`)),
			},
			expected: expected{
				code:  200,
				error: "",
			},
		},
		{
			name: "InvalidId",
			fields: fields{
				userService: MockUserService{},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("PUT", "/settings", strings.NewReader(`{"user_id":124, "email":"e123@example.com"}`)),
			},
			expected: expected{
				code:  404,
				error: "Not found",
			},
		},
		{
			name: "NothingToUpdate",
			fields: fields{
				userService: MockUserService{},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("PUT", "/settings", strings.NewReader(`{"user_id":123}`)),
			},
			expected: expected{
				code:  400,
				error: "No fields have been changed",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &handler{
				userService: tt.fields.userService,
			}
			h.UpdateUser(tt.args.w, tt.args.r)
		})
	}
}
