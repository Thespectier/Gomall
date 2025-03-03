package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/qingz2/Gomall/demo/demo_proto/biz/dal"
	"github.com/qingz2/Gomall/demo/demo_proto/biz/dal/mysql"
	"github.com/qingz2/Gomall/demo/demo_proto/biz/model"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dal.Init()
	// CURD
	//mysql.DB.Create(&model.User{Email: "example1.com", Password: "123456"})
	//mysql.DB.Model(&model.User{}).Where("email = ?", "example1.com").Update("password", "654321")
	var row model.User
	mysql.DB.Model(&model.User{}).Where("email = ?", "example1.com").First(&row)

	fmt.Printf("row: %+v\n", row)

	mysql.DB.Unscoped().Where("email = ?", "example1.com").Delete(&model.User{})
}
