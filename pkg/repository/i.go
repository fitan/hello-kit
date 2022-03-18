package repository

import (
	"hello/pkg/repository/api/baidu"
	"hello/pkg/repository/api/taobao"
)

type Repository struct {
	Baidu  baidu.BaiduApi
	Taobao taobao.TaobaoApi
}
