package dal

import (
	"github.com/qingz2/gomall/app/product/biz/dal/mysql"
	//"github.com/qingz2/gomall/app/product/biz/dal/redis"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
