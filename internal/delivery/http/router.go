package http

import (
	"EducationProject/internal/domain"
	"EducationProject/internal/usecase"
	"net/http"
)

//type RouterDeps struct {
//	TaskService usecase.TaskService
//}

func NewRouter(t usecase.TaskService, u usecase.UserService, tokenManager domain.TokenManager) http.Handler {
	mux := http.NewServeMux()
	taskHandler := NewTaskHandler(t)
	userHandler := NewUserHandler(u)
	auth := AuthMiddleware(tokenManager)

	// Задачи
	mux.Handle("POST /api/tasks", Chain(http.HandlerFunc(taskHandler.Create), auth))
	mux.Handle("GET /api/task/show/{id}", Chain(http.HandlerFunc(taskHandler.GetById), auth))
	mux.Handle("POST /api/task/update/{id}", Chain(http.HandlerFunc(taskHandler.Update), auth))
	mux.Handle("DELETE /api/task/delete/{id}", Chain(http.HandlerFunc(taskHandler.Delete), auth))

	// Пользователи
	mux.HandleFunc("POST /api/signup", userHandler.SignUp)
	mux.HandleFunc("POST /api/login", userHandler.Login)
	mux.HandleFunc("POST /api/logout", userHandler.Logout)

	return Chain(mux, Recoverer(), RequestLogger())
}
