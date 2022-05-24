package services

import (
	"hello/pkg/services/audit"
	"hello/pkg/services/casbin"
	"hello/pkg/services/project"
	"hello/pkg/services/user"
)

type HttpHandler struct {
	User    user.HttpHandler
	Audit   audit.HttpHandler
	Project project.HttpHandler
}

type Services struct {
	User    user.UserService
	Audit   audit.AuditService
	Project project.ProjectService
	Casbin  casbin.CasbinService
}
