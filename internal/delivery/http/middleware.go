package http

import (
	"EducationProject/internal/domain"
	"EducationProject/pkg/res"
	"context"
	"log"
	"net/http"
	"strings"
	"time"
)

type Middleware func(http.Handler) http.Handler

//type ContextKey string
//
//const UserIdContextKey ContextKey = "user_id"

func Chain(handler http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

func RequestLogger() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			log.Printf("Метод: %s , URL: %s , RemoteAddr: %s , Время выполнения: %s", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
		})
	}
}

func Recoverer() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if recovered := recover(); recovered != nil {
					log.Printf("Panic recovered: %v", recovered)
					res.WriteJSONError(w, http.StatusInternalServerError, "Внутренняя ошибка сервера")
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func AuthMiddleware(tokenManager domain.TokenManager) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			log.Println(authHeader)
			if authHeader == "" {
				res.WriteJSONError(w, http.StatusUnauthorized, "Токен авторизации не передан")
				return
			}
			tokenString := strings.TrimPrefix(authHeader, "Bearer")
			log.Println(tokenString)
			if tokenString == "" {
				res.WriteJSONError(w, http.StatusUnauthorized, "Токен пустой")
				return
			}
			claims, err := tokenManager.ParseToken(tokenString)
			log.Println(claims)
			if err != nil {
				res.WriteJSONError(w, http.StatusUnauthorized, "Токен не валиден или истек")
				return
			}
			ctx := context.WithValue(r.Context(), "userID", claims.UserId)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
