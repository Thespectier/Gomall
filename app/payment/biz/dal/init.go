package dal

import (
	"github.com/qingz2/gomall/app/payment/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
