package dao

import "static-server/model"

func UpdatePassword(username, newPassword string) error {
	_, err := dB.Exec("UPDATE users SET password = ? WHERE username = ?", newPassword, username)
	return err
}

func SelectUserByUsername(username string) (model.User, error) {
	user := model.User{}

	row := dB.QueryRow("SELECT id,password,name,selfInfo FROM users WHERE username = ? ", username)
	if row.Err() != nil {
		return user, row.Err()
	}
user.Username = username
	err := row.Scan(&user.Id, &user.Password,&user.Name,&user.SelfInfo)
	if err != nil {
		return user, err
	}

	return user, nil
}

func InsertUser(user model.User) error {
	_, err := dB.Exec("INSERT INTO user(username, password) "+"values(?, ?);", user.Username, user.Password)
	return err
}

