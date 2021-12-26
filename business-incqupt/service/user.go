package service

import (
	"business-incqupt/dao"
	"business-incqupt/model"
	"database/sql"
)

func IsRepeatUsername(username string)(bool,error){
	_,err:=dao.SelectUserByUsername(username)
	if err!=nil{
		if err==sql.ErrNoRows{
			return false,nil
		}
		return false,err
	}
	return true,nil
}
func Register(user model.User)error{
	err:=dao.InsertUser(user)
	return err
}
func IsPasswordCorrect(username,password string)(bool,error){
	user,err:=dao.SelectUserByUsername(username)
	if err!=nil{
		if err==sql.ErrNoRows{
			return false,nil
		}
		return false,err
	}
	if user.PassWord!=password{
		return false,nil
	}
	return true,nil
}
func ChangePassword(username,newPassword string)error{
	err:=dao.UpdataPassword(username,newPassword)
	return err

}

