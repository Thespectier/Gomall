package dal

import (
	"github.com/qingz2/gomall/app/checkout/biz/dal/mysql"
	"github.com/qingz2/gomall/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
