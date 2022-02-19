package dao

import (
	"database/sql"
	"static-server/model"
)

func InsertComment(comment model.Comment) error {
	_, err := dB.Exec("INSERT INTO comment(movieid, username,content,date,score) "+"values(?, ?, ?, ?, ?);", comment.Movieid,comment.Username,comment.Content,comment.Date,comment.Score)
	if err != nil {
	}
	return err
}
func HadUserPostedAComment(movieid int,username string)(bool,error){
	row := dB.QueryRow("SELECT id,content,dat,score FROM comment WHERE username = ? ", username)
    if row.Err()!=nil{

		return false,row.Err()
	}
	if row.Err() == sql.ErrNoRows{
		return false,nil
	}
	return true,nil
}
func SelectCommentsByMovieId(movieid int) ([]model.Comment,error){
		comment := model.Comment{}
        comments := make([]model.Comment,0)
		rows,err := dB.Query("SELECT id,username,content,dat,score FROM comment WHERE movieid = ? ", movieid)
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