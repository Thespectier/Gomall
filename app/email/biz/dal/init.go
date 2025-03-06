package dal

import (
	"github.com/qingz2/gomall/app/email/biz/dal/mysql"
	"github.com/qingz2/gomall/app/email/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
