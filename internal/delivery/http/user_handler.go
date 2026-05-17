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
	BirthDate time.Time `json:"birth_date,omitempty"` // TODO: загуглить и сделать необязательным полем
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
	//TODO: обработать ошибку, воспользоваться res.go
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		res.WriteJSONError(w, http.StatusBadRequest, "Не валидный JSON")
		return
	}
	requestInput := usecase.SignUpInput{FirstName: request.FirstName, LastName: request.LastName, BirthDate: request.BirthDate, Email: request.Email, Password: request.Password}
	h.u.SignUp(r.Context(), requestInput) //TODO: обработать возвращаемую ошибку
	m := map[string]string{"msg": "Пользователь успешно зарегистрирован"}
	res.WriteJSONResponse(w, http.StatusCreated, m)
}

//TODO: Сделать Login
// Создать структуру loginRequest
// по аналогии с 28-30 строкой получить json и распарсить
// вызвать метод Login из бизнес логики (поле u структуры UserHandler)
// вернуть токен, обработать ошибку

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var request loginRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		res.WriteJSONError(w, http.StatusBadRequest, "Не валидный JSON")
		return
	}
	//loginInput := usecase.LoginInput{Email: request.Email, Password: request.Password}
	//accessToken, err := h.u.Login(r.Context(), loginInput)
	//if err != nil {
	//	return domain.AuthToken{}, err
	//
	//}
	//return domain.AuthToken{AccessToken: accessToken}, nil

}

func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {

}
