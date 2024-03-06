package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id             uint64
	Name           string
	HashedPassword []byte
	Email          string
	Block          bool
}

func (user *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.HashedPassword = hashedPassword
	return nil
}

func (user *User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
	return err != nil //true se for igual
}
