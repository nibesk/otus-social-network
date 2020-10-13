package services

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"net/http"
	"service-chat/app/config"
	"service-chat/app/customErrors"
)

type ServiceUsers struct {
	Token string
}

const (
	getUserRoute = "/api/users/getUser"
)

type UserModel struct {
	User_id int    `json:"user_id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
}

type serverResponse struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
	Errors interface{} `json:"errors"`
}

var ErrUserNotFound = customErrors.TypedError{Msg: "User not found"}

func (s *ServiceUsers) GetUser() (*UserModel, error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", config.Env.Services.UsersUrl+getUserRoute, nil)
	if nil != err {
		return nil, errors.WithStack(err)
	}
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.Token))

	resp, err := client.Do(req)
	if nil != err {
		return nil, err
	}

	var result serverResponse
	json.NewDecoder(resp.Body).Decode(&result)

	if !result.Status {
		errs, _ := json.Marshal(result.Errors)

		return nil, customErrors.TypedError{string(errs)}
	}

	user := &UserModel{}
	err = mapstructure.Decode(result.Data.(map[string]interface{})["user"], user)
	if nil != err {
		return nil, errors.WithStack(err)
	}

	if 0 == user.User_id {
		return nil, ErrUserNotFound
	}

	return user, nil
}
