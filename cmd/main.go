package main

import (
	"EducationProject/internal/config"
	myhttp "EducationProject/internal/delivery/http"
	"EducationProject/internal/repository/cryptography"
	"EducationProject/internal/repository/postgres"
	"EducationProject/internal/usecase"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

//type MySting string
//
//func (m MySting) TestFunc () {
//	fmt.Println("TestFunc ok")
//}
//
//type MyFunc func (int, int)
//
//func (m MyFunc) TestFunc () {
//	fmt.Println("TestFunc ok 2")
//}
//
//func MyNewFunc (a, b int) {
//	fmt.Println("MyNewFunc")
//}

func main() {
	// Получаем конфигурацию
	cfg := config.New()
	if cfg == nil {
		fmt.Println("Ошибка загрузки конфигурации cfg")
	}
	log.Println("Получена конфигурация cfg")

	// Подключаемся к базе данных
	db := config.OpenDb(cfg.DB.GetDsn())
	if db == nil {
		fmt.Println("Ошибка подключения к базе данных")
	}
	log.Println("Подключение к базе данных установлено")

	// Выполняем автомиграцию
	config.AutoMigrate(db)
	log.Println("Автомиграция выполнена")

	bycryptCost, err := strconv.Atoi(cfg.BYCRYPTCOST)
	if err != nil {
		log.Println("Некорректный bycryptCost, используем 10")
		bycryptCost = 10
	}

	// Создаем слой репозитория
	hasher := cryptography.NewHasher(bycryptCost)
	tokenManager := cryptography.NewJWTManager("dev-secret", 24*time.Hour)
	taskRepo := postgres.NewTaskRepository(db)
	userRepo := postgres.NewUserRepository(db)

	// Создаем слой бизнес-логики
	taskUseCase := usecase.NewTaskUseCase(taskRepo)
	userUseCase := usecase.NewAuthUseCase(userRepo, hasher, tokenManager)

	// Создаем транспортный слой
	mux := myhttp.NewRouter(taskUseCase, userUseCase)

	log.Println("Сервер запущен на порту: " + cfg.SRV.Port)
	err = http.ListenAndServe(":"+cfg.SRV.Port, mux)
	if err != nil {
		log.Println(err)
	}
}
