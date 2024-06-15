package user

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dsypasit/social-clone/server/internal/share/util"
	"github.com/gorilla/mux"
)

type IUserService interface {
	GetUserByUUID(string) (User, error)
	CreateUser(UserCreated) (int64, error)
	GetUserByUsername(string) (User, error)
}

type UserHandler struct {
	userSrv IUserService
}

func NewUserHandler(userSrv IUserService) *UserHandler {
	return &UserHandler{userSrv}
}

func (h *UserHandler) GetUserByUUID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid, ok := vars["uuid"]
	if !ok || uuid == "" {
		util.SendJson(w, map[string]string{"message": "invalid uuid"}, http.StatusBadRequest)
		return
	}

	userQuery, err := h.userSrv.GetUserByUUID(uuid)
	if err != nil {
		if err == ErrUserNotFound {
			util.SendJson(w, map[string]string{"message": "user not found"}, http.StatusNotFound)
			return
		}
		util.SendJson(w, map[string]string{"message": "failed to get user by uuid", "error": err.Error()},
			http.StatusInternalServerError)
		return
	}

	util.SendJson(w, userQuery, http.StatusOK)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser UserCreated
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		util.SendJson(w, map[string]string{"message": "invalid body"}, http.StatusBadRequest)
		return
	}

	errResponse := util.BuildErrResponse("Failed to create user")
	if !util.IsValidEmail(newUser.Email) {
		util.SendJson(w, errResponse(errors.New("invalid email format")), http.StatusBadRequest)
		return
	}
	_, err := h.userSrv.CreateUser(newUser)
	if err != nil {
		util.SendJson(w, errResponse(err), http.StatusInternalServerError)
		return
	}

	util.SendJson(w, util.BuildResponse("User created successfully!"), http.StatusCreated)
}

func (h *UserHandler) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	errInvalidRes := util.BuildErrResponse("invalid request")
	if username == "" {
		util.SendJson(w, errInvalidRes(errors.New("invalid username")), http.StatusBadRequest)
		return
	}
	user, err := h.userSrv.GetUserByUsername(username)
	errServiceRes := util.BuildErrResponse("service failure")
	if err != nil {
		if err == ErrUserNotFound {
			util.SendJson(w, util.BuildResponse("user not found"), http.StatusNotFound)
			return
		}
		util.SendJson(w, errServiceRes(err), http.StatusInternalServerError)
		return
	}

	userRes := UserResponse{
		user.UUID, user.Username, user.Email,
	}

	util.SendJson(w, userRes, http.StatusOK)
}
