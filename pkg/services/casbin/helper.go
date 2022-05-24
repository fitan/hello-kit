package casbin

import "strconv"

const (
	UserPre = "user_"
	RolePre = "role_"
)

type CheckPermission struct {
	User       int
	Domain     string
	ResourceId int
}

func (c CheckPermission) CasbinObj() []interface{} {
	return []interface{}{c.UserKey(), c.DomainKey(), c.ResourceKey()}
}

func (c CheckPermission) UserKey() string {
	return UserPre + strconv.Itoa(c.User)
}

func (c CheckPermission) DomainKey() string {
	return c.Domain
}

func (c CheckPermission) ResourceKey() string {
	return strconv.Itoa(c.ResourceId)
}

type Permission struct {
	UserId int
	RoleId int
	Domain string
}

func (p Permission) CasbinObj() []interface{} {
	return []interface{}{p.UserKey(), p.RoleKey(), p.DomainKey()}
}

func (p Permission) UserKey() string {
	return UserPre + strconv.Itoa(p.UserId)
}

func (p Permission) RoleKey() string {
	return RolePre + strconv.Itoa(p.RoleId)
}

func (p Permission) DomainKey() string {
	return p.Domain
}

type Role struct {
	RoleId     int
	ResourceId int
}

func (r Role) CasbinObj() []interface{} {
	return []interface{}{r.NameKey(), r.ResourceKey()}
}

func (r Role) NameKey() string {
	return RolePre + strconv.Itoa(r.RoleId)
}

func (r Role) ResourceKey() string {
	return strconv.Itoa(r.ResourceId)
}
