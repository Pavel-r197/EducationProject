package main

import (
	"EducationProject/internal/adapter/postgres/stdsql"
	"EducationProject/internal/config"
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	//Получаем конфигурацию
	cfg := config.New()
	if cfg == nil {
		fmt.Println("Ошибка загрузки конфигурации cfg")
	}
	log.Println("Получена конфигурация cfg")

	// Подключаемся к базе данных
	db := stdsql.OpenDB(cfg.DB.GetDsn())
	if db == nil {
		fmt.Println("Ошибка подключения к базе данных")
	}
	log.Println("Подключение к базе данных установлено")

	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*1)
	defer cancel()

}
