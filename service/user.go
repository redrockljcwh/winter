package service

import (
	"database/sql"
	"fmt"
	"static-server/dao"
	"static-server/model"
)

func ChangePassword(username, newPassword string) error {
	err := dao.UpdatePassword(username, newPassword)
	return err
}

func IsPasswordCorrect(username, password string) (bool, error) {
	user, err := dao.SelectUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	if user.Password != password {
		return false, nil
	}

	return true, nil
}

//判断用户名是否重复
func IsRepeatUsername(username string) (bool, error) {
	_, err := dao.SelectUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func Register(user model.User) error {
	err := dao.InsertUser(user)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
func UpdateSelfInfo(username,selfInfo string)(error){
	err:=dao.UpdateSelfInfo(username,selfInfo)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func UserInfo(username string)(model.User,error){
	user,err:=dao.SelectUserByUsername(username)
	if err != nil{
		fmt.Println(err)
		return model.User{}, err
	}
	return user,err
}