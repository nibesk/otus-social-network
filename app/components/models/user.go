package models

import (
	"github.com/badThug/otus-social-network/app/components/storage"
	"github.com/pkg/errors"
)

type User struct {
	User_id    int
	Name       string
	Email      string
	Password   string
	Created_at string
	Updated_at string
}

func UserCreate(conn *storage.DbConnection, name, email, password string) (*User, error) {
	db := conn.GetDb()
	insert, err := db.Prepare("INSERT INTO `user`(name, email, password) VALUES(?, ?, ?)")
	if err != nil {
		return nil, err
	}

	result, err := insert.Exec(name, email, password)
	if err != nil {
		return nil, err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user := &User{
		User_id:  int(userId),
		Name:     name,
		Email:    email,
		Password: password,
	}

	return user, nil
}

func UserFindById(conn *storage.DbConnection, userId int) (*User, error) {
	db := conn.GetDb()

	query := db.QueryRow("SELECT * FROM user WHERE user_id = ?", userId)

	user := &User{}
	err := query.Scan(&user.User_id, &user.Name, &user.Email, &user.Password, &user.Created_at, &user.Updated_at)
	if err != nil {
		return nil, errors.Wrapf(err, "User id %d is not found", userId)
	}

	return user, nil
}
