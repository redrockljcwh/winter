package dao

import (
	"fmt"
	"static-server/model"
)

func SelectStarById(id int) (model.Star, error) {
	star := model.Star{}

	row := dB.QueryRow("SELECT name,birth,birthplace,work,introduction,picNum FROM stars WHERE id = ? ", id)
	if row.Err() != nil {
		return star, row.Err()
	}
	star.Id = id
	err := row.Scan(&star.Name,&star.Birth,&star.Birthplace,&star.Work,&star.Introduction,&star.PicNum)
	if err != nil {
		fmt.Println(err)
		return star, err
	}

	return star, nil
}