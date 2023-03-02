package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yonsina94/go-generics/components/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	if db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}); err == nil {
		route := gin.Default()

		usrController := user.Instance(user.New(db))

		usrController.AssingEndpoints(route.Group("/users"))
	} else {
		panic(err)
	}
}
