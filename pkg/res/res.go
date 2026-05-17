package res

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteJSONResponse(w http.ResponseWriter, code int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println("Ошибка при записи JSON ответа")
	}
}

func WriteJSONError(w http.ResponseWriter, code int, message string) {
	WriteJSONResponse(w, code, map[string]string{"error": message})
}
