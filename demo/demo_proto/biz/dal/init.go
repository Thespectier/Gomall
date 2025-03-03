package dal

import (
	"github.com/qingz2/Gomall/demo/demo_proto/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
