package dal

import (
	"github.com/qingz2/gomall/app/user/biz/dal/mysql"
	"github.com/qingz2/gomall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
