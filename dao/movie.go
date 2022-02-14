package dao

import "static-server/model"

func SelectMovieById(id int) (model.Movie, error) {
	movie := model.Movie{}

	row := dB.QueryRow("SELECT name,director,main_performer,type,country,langua,ondate,length,stuff FROM movie WHERE id = ? ", id)
	if row.Err() != nil {
		return movie, row.Err()
	}
	movie.Id = id
	err := row.Scan(&movie.Name,&movie.Director,&movie.Main_performer,&movie.Type,&movie.Country,&movie.Language,&movie.Date,&movie.Length,&movie.Stuff,&movie.Writer,&movie.PicNum)
	if err != nil {
		return movie, err
	}

	return movie, nil
}