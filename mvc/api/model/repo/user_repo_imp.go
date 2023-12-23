package repo

import (
	"database/sql"

	"github.com/Podre-Henrique/arquitetura-api/mvc/api/model"
)

type MysqlUserRepo struct {
	DB *sql.DB
}

func (ur *MysqlUserRepo) CreateUser(u *model.User) {
	query := "INSERT INTO users(name,password,email) values(?,?,?)"
	ur.DB.Exec(query, u.Name, u.HashedPassword, u.Email)
}

func (ur *MysqlUserRepo) BlockUser(email string) {
	query := "UPDATE users SET block=TRUE WHERE email=?"
	ur.DB.Exec(query, email)
}

func (ur *MysqlUserRepo) GetUser(userId uint64) *model.User {
	var user model.User
	query := "SELECT name,email FROM users WHERE id=? AND block=FALSE"
	ur.DB.QueryRow(query, userId).Scan(&user.Name, &user.Email)
	return &user
}

func (ur *MysqlUserRepo) LoginUser(email string, password string) *model.User {
	var user model.User
	query := "SELECT id,name,email,password FROM users WHERE email=? AND block=FALSE"
	ur.DB.QueryRow(query, email).Scan(&user.Id, &user.Name, &user.Email, &user.HashedPassword)
	if !user.VerifyPassword(password) {
		return nil
	}
	return &user
}
