package postgres

import (
	"os"

	"github.com/spf13/viper"
)

type DBParams struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func GetDBParams() DBParams {
	return DBParams{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
	}
}
