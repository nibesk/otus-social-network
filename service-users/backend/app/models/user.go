package models

import (
	"database/sql"
	"fmt"
	"github.com/VividCortex/mysqlerr"
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"service-users/app/customErrors"
	"service-users/app/globals"
	"service-users/app/handlers/requests"
	"service-users/app/storage"
	"strings"
)

const SexMale = 1
const SexFemale = 2
const SexOther = 3

type User struct {
	User_id    int            `json:"user_id"`
	Name       string         `json:"name"`
	Email      string         `json:"email"`
	Surname    string         `json:"surname"`
	Age        int            `json:"age"`
	Interests  string         `json:"interests"`
	City       string         `json:"city"`
	Sex        int            `json:"sex"`
	Password   string         `json:"-"`
	Token      sql.NullString `json:"token"`
	Created_at string         `json:"-"`
	Updated_at string         `json:"-"`
}

func (u *User) getSexTitle() string {
	var sexTitle string
	switch u.Sex {
	case SexMale:
		sexTitle = "male"
	case SexFemale:
		sexTitle = "female"
	default:
		sexTitle = "other"
	}

	return sexTitle
}

func (u *User) Public() map[string]interface{} {
	return map[string]interface{}{
		"user_id":   u.User_id,
		"name":      u.Name,
		"surname":   u.Surname,
		"age":       u.Age,
		"sex":       u.Sex,
		"sexTitle":  u.getSexTitle(),
		"interests": u.Interests,
		"city":      u.City,
		"email":     u.Email,
	}
}

func (u *User) UpdateToken(db storage.Executable, token interface{}) error {
	stmt, err := db.Prepare("UPDATE `user` SET token = ? where user_id = ?")
	if err != nil {
		return errors.WithStack(err)
	}

	result, err := stmt.Exec(token, u.User_id)
	if err != nil {
		return errors.WithStack(err)
	}

	stringToken, ok := token.(string)
	if !ok {
		u.Token = sql.NullString{String: ""}
	} else {
		u.Token = sql.NullString{String: stringToken}
	}

	return checkIsExecDidUpdate(result)
}

func UserCreate(db storage.Executable, user *User) (*User, error) {
	insert, err := db.Prepare("INSERT INTO `user`(name, email, password, surname, age, city, interests, sex) VALUES(?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	result, err := insert.Exec(user.Name, user.Email, user.Password, user.Surname, user.Age, user.City, user.Interests, user.Sex)
	if driverErr, ok := err.(*mysql.MySQLError); ok {
		if driverErr.Number == mysqlerr.ER_DUP_ENTRY {
			return nil, errors.Wrap(&customErrors.TypedError{fmt.Sprintf("User with email %s already exist", user.Email)}, driverErr.Message)
		}

		return nil, errors.WithStack(err)
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	user.User_id = int(userId)

	return user, nil
}

func UserFindById(db storage.Queryable, userId int) (*User, error) {
	query := db.QueryRow("SELECT * FROM user WHERE user_id = ?", userId)

	user := &User{}
	err := userQueryScan(query.Scan, user)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return user, nil
}

func UserFindAllExceptUserId(db storage.Queryable, userId int, searchParams requests.AvailableFriendsRequest) ([]*User, error) {
	queryParams := []interface{}{}
	scrollSth := ""
	if 0 != searchParams.LastViewedUserId {
		scrollSth = " AND user_id > ? "
		queryParams = append(queryParams, searchParams.LastViewedUserId)
	}

	likeSearchSth := ""
	if "" != searchParams.Name && "" != searchParams.Surname {
		likeSearchSth = " AND name LIKE ? AND surname LIKE ? "
		queryParams = append(queryParams, searchParams.Name+"%", searchParams.Surname+"%")
	}

	queryParams = append(queryParams, globals.DefaultScrollLimit)

	sth := fmt.Sprintf("SELECT * FROM user WHERE 1=1 %s %s ORDER BY user_id ASC LIMIT ?", scrollSth, likeSearchSth)

	query, err := db.Query(sth, queryParams...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	collection := []*User{}
	for query.Next() {
		user := &User{}
		err := userQueryScan(query.Scan, user)
		if err != nil {
			return nil, err
		}
		if userId == user.User_id {
			continue
		}
		collection = append(collection, user)
	}

	return collection, nil
}

func UserFindByEmail(db storage.Queryable, email string) (*User, error) {
	query := db.QueryRow("SELECT * FROM user WHERE email = ?", email)

	user := &User{}
	err := userQueryScan(query.Scan, user)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return user, nil
}

func UserFindByUserIds(db storage.Queryable, userIds []int) ([]*User, error) {
	if 0 == len(userIds) {
		return []*User{}, nil
	}

	args := make([]interface{}, len(userIds))
	for i, id := range userIds {
		args[i] = id
	}

	query, err := db.Query("SELECT * FROM user WHERE user_id in (? "+strings.Repeat(",?", len(args)-1)+")", args...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	collection := []*User{}
	for query.Next() {
		user := &User{}
		err := userQueryScan(query.Scan, user)
		if err != nil {
			return nil, err
		}
		collection = append(collection, user)
	}

	return collection, nil
}

func UserFindFriendsForUser(db storage.Queryable, userId int) ([]*User, error) {
	query, err := db.Query(
		"SELECT u.* FROM user u "+
			"JOIN user_relation ur ON u.user_id = ur.friend_user_id "+
			"WHERE ur.user_id = ?", userId)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	collection := []*User{}
	for query.Next() {
		user := &User{}
		err := userQueryScan(query.Scan, user)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		collection = append(collection, user)
	}

	return collection, nil
}

func userQueryScan(scan func(dest ...interface{}) error, user *User) error {
	err := scan(
		&user.User_id,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Token,
		&user.Surname,
		&user.Created_at,
		&user.Updated_at,
		&user.Age,
		&user.City,
		&user.Interests,
		&user.Sex)

	if err != nil {
		return err
	}

	return nil
}
