package models

import (
	"gin_curd/database"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var Users []User

func (user *User) GetUsers() (users []User, err error) {
	if err = database.Eloquent.Find(&users).Error; err != nil {
		return
	}
	return
}

func (user *User) Insert() (id int64, err error) {
	result := database.Eloquent.Create(&user)
	id = user.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

func (user *User) Update(id int64) (updateUser User, err error) {
	if err = database.Eloquent.First(&updateUser, id).Error; err != nil {
		return
	}
	if err = database.Eloquent.Model(&updateUser).Updates(&user).Error; err != nil {
		return
	}
	return
}

func (user *User) DestroyUser(id int64) (deleteUser User, err error) {
	if err = database.Eloquent.First(&deleteUser, id).Error; err != nil {
		return
	}
	if err = database.Eloquent.Delete(&deleteUser).Error; err != nil {
		return
	}
	return
}
