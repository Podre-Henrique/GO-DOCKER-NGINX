package repo

import "github.com/Podre-Henrique/arquitetura-api/mvc/api/model"

type UserRepo interface {
	CreateUser(u *model.User)
	BlockUser(email string)
	GetUser(userId uint64) *model.User
	LoginUser(email string, password string) *model.User
}
