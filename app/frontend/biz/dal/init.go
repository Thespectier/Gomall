package dal

import (
	"github.com/qinz2/Gomall/app/frontend/biz/dal/mysql"
	"github.com/qinz2/Gomall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
