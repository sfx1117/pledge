package entity

type UserEntity struct {
	Name     string `form:"name" binding:"required"`
	Password string `form:"password" bindings:"required"`
}
