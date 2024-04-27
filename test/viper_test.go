package test

import (
	"github.com/Catlordx/CampusTrade/internal/db/mysql"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadDbConfigFromViper(t *testing.T) {
	v := viper.New()
	v.SetConfigFile("../configs/config.dev.toml")
	//conf, err := mysql.DbConfig{}.LoadDbConfigFromViper(v)
	var conf mysql.DbConfig
	conn, err := conf.LoadDbConfigFromViper(v)
	if err != nil {
		t.Fatalf("Failed to load DB config from viper: %v", err)
	}
	assert.Equal(t, "cat:jyxt$098@123@tcp(8.130.120.24:3306)/demo?charset=utf8mb4&parseTime=True", conn)
}
