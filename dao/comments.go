package dao

import (
	"database/sql"
	"static-server/model"
)

func InsertComment(comment model.Comment) error {
	_, err := dB.Exec("INSERT INTO comment(movieid, username,content,date,score) "+"values(?, ?, ?, ?, ?);", comment.Movieid,comment.Username,comment.Score,comment.Date,comment.Score)
	return err
}
func HadUserPostedAComment(movieid int,username string)(bool,error){
	row := dB.QueryRow("SELECT id,content,date,score FROM comment WHERE movieid = ? and username = ? ", movieid,username)
    if row.Err()!=nil{
	if row.Err() == sql.ErrNoRows{
		return false,nil
	}else{
		return false,row.Err()
	}
	}
	return true,nil
}
func SelectCommentsByMovieId(movieid int) ([]model.Comment,error){
		comment := model.Comment{}
        comments := make([]model.Comment,0)
		rows,err := dB.Query("SELECT id,username,content,date,score FROM comment WHERE movieid = ? ", movieid)
		defer rows.Close()
		if err != nil{
			return comments, err
		}
		comment.Movieid=movieid
		for rows.Next() {
			err := rows.Scan(&comment.Id, &comment.Username, &comment.Content, &comment.Date, &comment.Score)
			if err != nil {
				return comments, err
			}
			comments = append(comments,comment)
		}
		return comments, nil
}