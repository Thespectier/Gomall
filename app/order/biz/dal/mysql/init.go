package mysql

import (
	"fmt"
	"os"

	"github.com/cloudwego/biz-demo/gomall/app/order/biz/model"
	"github.com/cloudwego/biz-demo/gomall/app/order/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN,
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DATABASE"),
	)

	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	type Version struct {
		Version string
	}
	var v Version
	err = DB.Raw("SELECT VERSION() as version").Scan(&v).Error
	if err != nil {
		panic(err)
	}
	fmt.Println("MySQL Version:", v.Version)

	// TODO: add more tables and models here
	if os.Getenv("GO_ENV") != "online" {
		if err = DB.AutoMigrate(
			&model.Order{},
			&model.OrderItem{},
		); err != nil {
			panic(err)
		}
	}
}
