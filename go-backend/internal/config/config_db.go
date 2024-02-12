package config

import (
	"distributed-calculator/internal/database"
	"errors"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/joho/godotenv"
)

func GetDBParams() (database.DBParams, error) {
	// Загружаем виртуальное окружение
	projectName := regexp.MustCompile(`^(.*` + `go-backend` + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	if err := godotenv.Load(string(rootPath) + `/.env`); err != nil {
		log.Print("No .env file found")
		return database.DBParams{}, err
	}

	// Получаем переменные для бд
	db_host, found := os.LookupEnv("DB_HOST")
	if !found {
		env_var_not_found("DB_HOST")
		return database.DBParams{}, errors.New("Param not found")
	}

	port, found := os.LookupEnv("DB_PORT")
	if !found {
		env_var_not_found("DB_PORT")
		return database.DBParams{}, errors.New("Param not found")
	}
	db_port, err := strconv.Atoi(port)
	if err != nil {
		log.Printf("Error: Invalid type for DB_PORT. Must be integer")
		return database.DBParams{}, errors.New("Param not found")
	}

	db_name, found := os.LookupEnv("DB_NAME")
	if !found {
		env_var_not_found("DB_NAME")
		return database.DBParams{}, errors.New("Param not found")
	}
	db_user, found := os.LookupEnv("DB_USER")
	if !found {
		env_var_not_found("DB_USER")
		return database.DBParams{}, errors.New("Param not found")
	}
	db_pass, found := os.LookupEnv("DB_PASS")
	if !found {
		env_var_not_found("DB_PASS")
		return database.DBParams{}, errors.New("Param not found")
	}
	return database.DBParams{Host: db_host, Port: db_port,
		Username: db_user, Password: db_pass,
		DBName: db_name}, nil

}

func Сreate_all_tables_in_db() {
	DBParams, err := GetDBParams()

	if err != nil {
		log.Println("Error with db params")
	}

	database.New(DBParams)

}

func env_var_not_found(val string) {
	log.Printf("Error: Variable %s not found in .env", val)
}
