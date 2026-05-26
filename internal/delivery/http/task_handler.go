package http

import (
	"EducationProject/internal/usecase"
	"EducationProject/pkg/res"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type TaskHandler struct {
	u usecase.TaskService
}

type TaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

//type TaskResponse struct {
//	Id          int64     `json:"id"`
//	Title       string    `json:"title"`
//	Description string    `json:"description"`
//	UserId      int64     `json:"user_id"`
//	CreatedAt   time.Time `json:"created_at"`
//	UpdatedAt   time.Time `json:"updated_at"`
//}

func NewTaskHandler(u usecase.TaskService) *TaskHandler {
	return &TaskHandler{u: u}
}

func (t *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		res.WriteJSONError(w, http.StatusBadRequest, "Не валидный JSON")
		return
	}
	task := usecase.TaskInput{Title: request.Title, Description: request.Description}
	newTask, err := t.u.Create(r.Context(), task)
	//responseTask := TaskResponse{Id: newTask.Id, Title: newTask.Title, Description: newTask.Description, UserId: newTask.UserId, CreatedAt: newTask.CreatedAt, UpdatedAt: newTask.UpdatedAt}
	if err != nil {
		res.WriteJSONError(w, http.StatusInternalServerError, "Не получилось создать запись")
		return
	}
	res.WriteJSONResponse(w, http.StatusCreated, newTask)
}

func (t *TaskHandler) GetById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		res.WriteJSONError(w, http.StatusBadRequest, "ID должно быть числом")
		return
	}
	task, err := t.u.GetById(r.Context(), int64(id))
	if err != nil {
		res.WriteJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	res.WriteJSONResponse(w, http.StatusOK, task)
}

func (t *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	// TO-DO: переписать структуру, чтобы получали id из параметра пути как в 51 строке
	//id, err := strconv.Atoi(r.PathValue("id"))
	//if err != nil {
	//	res.WriteJSONError(w, http.StatusBadRequest, "ID должно быть числом")
	//	return
	//}

	request := struct {
		Id          int64  `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		res.WriteJSONError(w, http.StatusBadRequest, "Не валидный JSON")
		return
	}
	task := usecase.UpdateTask{Id: request.Id, Title: request.Title, Description: request.Description}
	if err := t.u.Update(r.Context(), task); err != nil {
		res.WriteJSONError(w, http.StatusInternalServerError, "Не получилось обновить запись")
		return
	}
	msg := fmt.Sprintf("Задача %d обновлена", request.Id)
	res.WriteJSONResponse(w, http.StatusOK, map[string]string{"msg": msg})
}

func (t *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		res.WriteJSONError(w, http.StatusBadRequest, "Id должно быть числом")
		return
	}
	err = t.u.Delete(r.Context(), int64(id))
	if err != nil {
		res.WriteJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	msg := fmt.Sprintf("Задача %d удалена", id)
	res.WriteJSONResponse(w, http.StatusOK, map[string]string{"msg": msg})
}
