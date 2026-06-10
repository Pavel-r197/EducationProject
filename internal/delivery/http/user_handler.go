package http

import (
	"EducationProject/internal/usecase"
	"EducationProject/pkg/res"
	"encoding/json"
	"net/http"
	"time"
)

// TODO: добавить остальные поля
type signUpRequest struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BirthDate time.Time `json:"birth_date,omitempty"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserHandler struct {
	u usecase.UserService
}

func NewUserHandler(u usecase.UserService) *UserHandler {
	return &UserHandler{u: u}
}

func (h *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var request signUpRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		res.WriteJSONError(w, http.StatusBadRequest, "Не валидный JSON")
		return
	}
	requestInput := usecase.SignUpInput{FirstName: request.FirstName, LastName: request.LastName, BirthDate: request.BirthDate, Email: request.Email, Password: request.Password}
	if err := h.u.SignUp(r.Context(), requestInput); err != nil {
		res.WriteJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	m := map[string]string{"msg": "Пользователь успешно зарегистрирован"}
	res.WriteJSONResponse(w, http.StatusCreated, m)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var request loginRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		res.WriteJSONError(w, http.StatusBadRequest, "Не валидный JSON")
		return
	}
	loginInput := usecase.LoginInput{Email: request.Email, Password: request.Password}
	accessToken, err := h.u.Login(r.Context(), loginInput)
	if err != nil {
		res.WriteJSONError(w, http.StatusBadRequest, err.Error())
		return

	}
	res.WriteJSONResponse(w, http.StatusOK, accessToken)

}

func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {

}
