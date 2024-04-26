$goCommand = "go"

# install or update package GROM
$gormPackage = "gorm.io/gorm"
Invoke-Expression -Command "$goCommand get -u $gormPackage"

# install or update package MySQL Driver
$mysqlDriverPackage = "gorm.io/driver/sqlite"
Invoke-Expression -Command "$goCommand get -u $mysqlDriverPackage"

# install or update package Gin
$ginPackage = "github.com/gin-gonic/gin"
Invoke-Expression -Command "$goCommand get -u $ginPackage"

# install or update package Viper
$viperPackage = "github.com/spf13/viper"
Invoke-Expression -Command "$goCommand get -u $viperPackage"

Invoke-Expression -Command "$goCommand mod tidy"