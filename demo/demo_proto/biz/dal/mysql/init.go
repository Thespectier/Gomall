package mysql

import (
	"fmt"
	"os"

	"github.com/qingz2/Gomall/demo/demo_proto/biz/model"
	"github.com/qingz2/Gomall/demo/demo_proto/conf"

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
	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
}
