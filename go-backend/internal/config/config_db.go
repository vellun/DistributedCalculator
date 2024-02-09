package config

import (
	"distributed-calculator/internal/database"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/joho/godotenv"
)

func Сreate_all_tables_in_db() {
	// Загружаем виртуальное окружение
	projectName := regexp.MustCompile(`^(.*` + `go-backend` + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	if err := godotenv.Load(string(rootPath) + `/.env`); err != nil {
		log.Print("No .env file found")
		return
	}

	// Получаем переменные для бд
	db_host, found := os.LookupEnv("DB_HOST")
	if !found {
		env_var_not_found("DB_HOST")
		return
	}

	port, found := os.LookupEnv("DB_PORT")
	if !found {
		env_var_not_found("DB_PORT")
		return
	}
	db_port, err := strconv.Atoi(port)
	if err != nil {
		log.Printf("Error: Invalid type for DB_PORT. Must be integer")
		return
	}

	db_name, found := os.LookupEnv("DB_NAME")
	if !found {
		env_var_not_found("DB_NAME")
		return
	}
	db_user, found := os.LookupEnv("DB_USER")
	if !found {
		env_var_not_found("DB_USER")
		return
	}
	db_pass, found := os.LookupEnv("DB_PASS")
	if !found {
		env_var_not_found("DB_PASS")
		return
	}

	database.New(database.DBParams{Host: db_host, Port: db_port,
		Username: db_user, Password: db_pass,
		DBName: db_name})

}

func env_var_not_found(val string) {
	log.Printf("Error: Variable %s not found in .env", val)
}
