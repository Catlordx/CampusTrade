package mysql

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	username string
	password string
	role     string
}

type Commodity struct {
	gorm.Model
	name        string
	description string
	price       float64
	owner       string
}

type User struct {
	gorm.Model
	name     string
	password string
	identity string
}
