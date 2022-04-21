package repository

import (
	"hello/pkg/repository/api/baidu"
	"hello/pkg/repository/api/taobao"
	"hello/pkg/repository/dao/pod"
	"hello/pkg/repository/dao/tblservicetree"
	"hello/pkg/repository/dao/user"
)

type Repository struct {
	Baidu  baidu.BaiduService
	Taobao taobao.TaobaoService
	User   user.UserService
	Pod pod.PodService
	ServiceTree tblservicetree.TblservicetreeService
}
