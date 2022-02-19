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
func SelectMovieByName(name string) ([]model.Movie, error) {
	var movies = make([]model.Movie,0)
	var movie model.Movie

	rows,err := dB.Query("SELECT name,director,main_performer,type,country,langua,ondate,length,stuff FROM movie WHERE name = ? ","%"+name+"%" )
	defer rows.Close()
	if rows.Err() != nil {
		return movies, rows.Err()
	}
	for rows.Next() {
		err = rows.Scan(&movie.Name, &movie.Director, &movie.Main_performer, &movie.Type, &movie.Country, &movie.Language, &movie.Date, &movie.Length, &movie.Stuff, &movie.Writer, &movie.PicNum)
		if err != nil {
			return movies, err
		}
		movies = append(movies,movie)
	}
	return movies, nil
}
func InsertMovie(movie model.Movie) error {
	_, err := dB.Exec("INSERT INTO movie(name, director,main_performer,type,country,langua,ondate) "+"values(?, ?, ?, ?, ?, ?, ?);",movie.Name,movie.Director,movie.Main_performer,movie.Type,movie.Country,movie.Language,movie.Date )
	return err
}