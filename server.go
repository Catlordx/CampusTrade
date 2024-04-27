package main

import (
	"fmt"
	"github.com/spf13/viper"

	alimysql "github.com/Catlordx/CampusTrade/internal/db/mysql"

	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	v := viper.New()
	v.SetConfigFile("./configs/config.dev.toml")
	var conf = alimysql.DbConfig{}
	conn, err := conf.LoadDbConfigFromViper(v)
	if err != nil {
		panic("Failed to load config")
	}

	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&Product{})
	if err != nil {
		return
	}
	db.Create(&Product{Code: "D42", Price: 100})

	var product Product
	db.First(&product, 1)
	db.First(&product, "code = ?", "D42")
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	err = r.Run()
	if err != nil {
		return
	}
	fmt.Println("Hello!World")
}
