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
