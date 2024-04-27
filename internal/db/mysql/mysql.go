package mysql

import (
	"fmt"
	"github.com/spf13/viper"
)

type DbConfig struct {
	host     string
	username string
	password string
	port     int
	dbname   string
}

func (db *DbConfig) LoadDbConfigFromViper(v *viper.Viper) (string, error) {

	if err := v.ReadInConfig(); err != nil {
		return "", err
	}
	db.dbname = v.GetString("database.dbname")
	db.host = v.GetString("database.host")
	db.password = v.GetString("database.password")
	db.username = v.GetString("database.username")
	db.port = v.GetInt("database.port")
	return db.ToString(), nil
}

func (db *DbConfig) ToString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		db.username,
		db.password,
		db.host,
		db.port,
		db.dbname)
}
