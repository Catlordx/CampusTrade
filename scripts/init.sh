#!/bin/bash

# define command `go`
go_command="go"

# install or update package gorm
gorm_package="gorm.io/gorm"
$go_command get -u $gorm_package

# install or update package MySQL Driver
mysql_driver_package="gorm.io/driver/mysql"
$go_command get -u $mysql_driver_package

# install or update package Gin
gin_package="github.com/gin-gonic/gin"
$go_command get -u $gin_package

# install or update package Viper
viper_package="github.com/spf13/viper"
$go_command get -u $viper_package

$go_command mod tidy