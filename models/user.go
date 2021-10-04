package models

import (
	"go-rest-crud/datasources/mysql"
	"log"
)

func CreateUser(user *User) (err error) {
	if err = mysql.Connect.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *User) (err error) {
	log.Println(user)
	mysql.Connect.Updates(user)
	return nil
}

func GetUserByID(user *User, id uint32) (err error) {
	if err = mysql.Connect.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func GetAllUsers(user *[]User) (err error) {
	if err = mysql.Connect.Find(user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(user *User, id uint32) (err error) {
	mysql.Connect.Where("id = ?", id).Delete(user)
	return nil
}
