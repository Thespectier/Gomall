package dal

import (
	"github.com/qingz2/gomall/app/cart/biz/dal/mysql"
	"github.com/qingz2/gomall/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
