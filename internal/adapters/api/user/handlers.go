package user

import (
	"auth-api/internal/adapters/api"
	userDomain "auth-api/internal/domain/user"
	customError "auth-api/internal/error"
	"auth-api/internal/midlleware"
	"auth-api/internal/utils"
	"encoding/json"
	"net/http"
)

const (
	createUserURL   = "/register"
	loginUserURL    = "/login"
	userProfileURL  = "/profile"
	userSettingsURL = "/settings"
	POST            = "POST "
	GET             = "GET "
	PUT             = "PUT "
	PATCH           = "PATCH "
	DELETE          = "DELETE "
)

type handler struct {
	userService userDomain.ServiceUser
}

func (h *handler) Register(router *http.ServeMux) {
	router.Handle(POST+createUserURL, midlleware.TimeoutMiddleware(http.HandlerFunc(h.CreateUser)))
	router.Handle(PUT+userSettingsURL, midlleware.TimeoutMiddleware(midlleware.AuthMiddleware(http.HandlerFunc(h.UpdateUser))))
	router.Handle(POST+loginUserURL, midlleware.TimeoutMiddleware(http.HandlerFunc(h.LoginUser)))
}

func NewHandler(service userDomain.ServiceUser) api.Handler {
	return &handler{userService: service}
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var dtoUser = &userDomain.CreateUserDTO{}
	if json.NewDecoder(r.Body).Decode(dtoUser) != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}
	if err := h.userService.CreateUser(r.Context(), dtoUser); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	utils.RenderJSON(w, http.StatusCreated, "user has been created")
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var dtoUser = &userDomain.UpdateUserDTO{}
	if err := json.NewDecoder(r.Body).Decode(dtoUser); err != nil {
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}
	u, err := h.userService.UpdateUser(r.Context(), dtoUser)
	if err != nil {
		if err.Error() == customError.NothingToUpdateError.Error() {
			w.Header().Set("X-Error-Message", "Resource not modified")
			http.Error(w, err.Error(), http.StatusNotModified)
			return
		} else if err.Error() == customError.NotFoundError.Error() {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		return
	}
	utils.RenderJSON(w, http.StatusOK, u)
}

func (h *handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var dtoUser = &userDomain.CreateUserDTO{}
	if json.NewDecoder(r.Body).Decode(dtoUser) != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}
	token, err := h.userService.Login(r.Context(), dtoUser)
	if err != nil {
		if err.Error() == customError.LoginError.Error() {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		} else {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
	}
	utils.RenderJSON(w, http.StatusOK, token)
	utils.SetCookie(w, token.Token)
}
