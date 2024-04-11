package midlleware

import (
	"net/http"
	"reflect"
	"testing"
)

func TestAuthMiddleware(t *testing.T) {
	type args struct {
		next http.Handler
	}
	tests := []struct {
		name string
		args args
		want http.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AuthMiddleware(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoggerRequestMiddleware(t *testing.T) {
	type args struct {
		next http.Handler
	}
	tests := []struct {
		name string
		args args
		want http.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoggerRequestMiddleware(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoggerRequestMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeoutMiddleware(t *testing.T) {
	type args struct {
		next http.Handler
	}
	tests := []struct {
		name string
		args args
		want http.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeoutMiddleware(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeoutMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}
