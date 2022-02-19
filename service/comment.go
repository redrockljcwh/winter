package service

import (
	"static-server/dao"
	"static-server/model"
)

func PostComment(comment model.Comment)error{
	err := dao.InsertComment(comment)
	return err
}

func CheckUserForComment(comment model.Comment)(bool,error){
	check,err:=dao.HadUserPostedAComment(comment.Movieid,comment.Username)
	return check,err
}

func GetCommentsByMovieId(movieid int)([]model.Comment,error){
	comments,err:=dao.SelectCommentsByMovieId(movieid)
	if err != nil{
		return nil, err
	}
	return comments,nil
}