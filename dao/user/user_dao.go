package user

import (
	"exam/core/db"
	"exam/models"
)

func SelectByPhone(phone string) (models.User, error) {
	sql := "select id, username, password, phone from user where phone = ?"
	var user models.User
	err := db.SelectOne(sql, &user, phone)
	return user, err
}
