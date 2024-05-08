package core

import (
	"github.com/Catlordx/CampusTrade/internal/core/config"
	"github.com/Catlordx/CampusTrade/internal/db/mysql"
)

func NewAppContext() (*config.AppContext, error) {
	conf := mysql.DbConfig{}
	db, err := mysql.Connect(&conf)
	if err != nil {
		return nil, err
	}
	return &config.AppContext{DB: db}, nil
}
