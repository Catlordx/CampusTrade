/*
Package mysql

Define mysql configuration
*/
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

// LoadDbConfigFromViper return a connect string
func (db *DbConfig) LoadDbConfigFromViper(v *viper.Viper) (string, error) {

	if err := v.ReadInConfig(); err != nil {
		return "", err
	}
	db.dbname = v.GetString("database.dbname")
	db.host = v.GetString("database.host")
	db.password = v.GetString("database.password")
	db.username = v.GetString("database.username")
	db.port = v.GetInt("database.port")
	return db.String(), nil
}

// convert mysql config struct to string
func (db *DbConfig) String() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		db.username,
		db.password,
		db.host,
		db.port,
		db.dbname)
}

func (db *DbConfig) WithName(name string) *DbConfig {
	db.username = name
	return db
}

// DbBuilder 建造者模式创建DbConfig
func DbBuilder() *DbConfig {
	return new(DbConfig)
}
func (db *DbConfig) WithPassword(password string) *DbConfig {
	db.password = password
	return db
}

func (db *DbConfig) WithPort(port int) *DbConfig {
	db.port = port
	return db
}

func (db *DbConfig) WithHost(host string) *DbConfig {
	db.host = host
	return db
}

func (db *DbConfig) WithDBName(dbname string) *DbConfig {
	db.dbname = dbname
	return db
}
func (db *DbConfig) Build() DbConfig {
	return *db
}
