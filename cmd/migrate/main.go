package main

import (
	"EducationProject/internal/adapter/postgres/migrations"
	"EducationProject/internal/adapter/postgres/stdsql"
	"EducationProject/internal/config"
	"context"
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	direction := flag.String("direction", "up", "up or down")
	flag.Parse()
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
	switch *direction {
	case "up":
		if err := migrations.UP(ctx, db); err != nil {
			log.Fatal(err)
		}
		log.Println("Миграции up успешно применены")
	case "down":
		if err := migrations.DOWN(ctx, db); err != nil {
			log.Fatal(err)
		}
		log.Println("Миграции down успешно применены")

	default:
		log.Fatal("Непонятное направление миграции")
	}
}
