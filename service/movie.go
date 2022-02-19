package service

import (
	"static-server/dao"
	"static-server/model"
)

func SelectMovie(id int)(model.Movie,error){
	movie,err:=dao.SelectMovieById(id)
	if err != nil {
		return movie, err
	}

	return movie, nil
}

func InsertMovie(movie model.Movie)(error){
	err:=dao.InsertMovie(movie)
	return err
}
func SearchMovieByName(name string)([]model.Movie,error){
	movies,err:=dao.SelectMovieByName(name)
	if err != nil{
		return nil, err
	}
	return movies,nil
}