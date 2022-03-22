package repository

import (
	"hello/pkg/repository/api/baidu"
	"hello/pkg/repository/api/taobao"
	"hello/pkg/repository/dao/user"
)

type Repository struct {
	Baidu  baidu.BaiduApi
	Taobao taobao.TaobaoApi
	User   user.UserService
}
