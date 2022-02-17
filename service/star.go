package service

import (
	"fmt"
	"static-server/dao"
	"static-server/model"
)

func SelectStarById(id int)(model.Star,error){
	star,err:=dao.SelectStarById(id)
	if err != nil{
		fmt.Println(err)
	}
	return star,err
}
