package http

import (
	"EducationProject/internal/usecase"
	"net/http"
)

//type RouterDeps struct {
//	TaskService usecase.TaskService
//}

func NewRouter(t usecase.TaskService, u usecase.UserService) http.Handler {
	mux := http.NewServeMux()
	taskHandler := NewTaskHandler(t)
	userHandler := NewUserHandler(u)

	// Задачи
	mux.HandleFunc("POST /api/tasks", taskHandler.Create)
	mux.HandleFunc("GET /api/task/show/{id}", taskHandler.GetById)
	mux.HandleFunc("POST /api/task/update/{id}", taskHandler.Update)
	mux.HandleFunc("DELETE /api/task/delete/{id}", taskHandler.Delete)

	// Пользователи
	mux.HandleFunc("POST /api/register", userHandler.Register)
	mux.HandleFunc("POST /api/login", userHandler.Login)

	return mux
}
