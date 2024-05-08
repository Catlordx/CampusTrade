package core

import (
	"github.com/Catlordx/CampusTrade/internal/db/mysql"
	"gorm.io/gorm"
)

type AppContext struct {
	DB *gorm.DB
}

func NewAppContext() (*AppContext, error) {
	conf := mysql.DbConfig{}
	db, err := mysql.Connect(&conf)
	if err != nil {
		return nil, err
	}
	return &AppContext{DB: db}, nil
}
