package casbin

import "strconv"

func CasbinUserKeyWrap(id int) string {
	return CasbinKeyWrap(id, "user")
}

func CasbinRoleKeyWrap(id int) string {
	return CasbinKeyWrap(id, "role")
}

func CasbinResourceWrap(id int) string {
	return CasbinKeyWrap(id, "resource")
}

func CasbinKeyWrap(id int, t string) string {
	idS := strconv.Itoa(id)
	return t + "_" + idS
}
