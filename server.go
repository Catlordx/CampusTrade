package main

import (
	"fmt"

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
	// v := vipr.New()
	// v.SetConfigFile("./configs/config.dev.toml")
	//if err := v.ReadInConfig(); err != nil {
	//	panic(fmt.Errorf("fatal error reading config file: %w", err))
	//}
	//fmt.Println(v.GetString("title"))
	dsn := "cat:xiaoyuan@123$456@tcp(8.130.120.24:3306)/demo?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
