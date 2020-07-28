package models

import (
	"fmt"
	"github.com/VividCortex/mysqlerr"
	"github.com/badThug/otus-social-network/app/customErrors"
	"github.com/badThug/otus-social-network/app/storage"
	"github.com/go-sql-driver/mysql"
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

func (u *User) Public() map[string]interface{} {
	return map[string]interface{}{
		"user_id": u.User_id,
		"name":    u.Name,
		"email":   u.Email,
	}
}

func UserCreate(conn *storage.DbConnection, name, email, password string) (*User, error) {
	db := conn.GetDb()
	insert, err := db.Prepare("INSERT INTO `user`(name, email, password) VALUES(?, ?, ?)")
	if err != nil {
		return nil, err
	}

	result, err := insert.Exec(name, email, password)
	if driverErr, ok := err.(*mysql.MySQLError); ok {
		if driverErr.Number == mysqlerr.ER_DUP_ENTRY {
			return nil, errors.Wrap(&customErrors.TypedError{fmt.Sprintf("User with email %s already exist", email)}, driverErr.Message)
		}

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
		return nil, err
	}

	return user, nil
}

func UserFindByEmail(conn *storage.DbConnection, email string) (*User, error) {
	db := conn.GetDb()

	query := db.QueryRow("SELECT * FROM user WHERE email = ?", email)

	user := &User{}
	err := query.Scan(&user.User_id, &user.Name, &user.Email, &user.Password, &user.Created_at, &user.Updated_at)
	if err != nil {
		return nil, err
	}

	return user, nil
}
