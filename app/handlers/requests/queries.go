package requests

import (
	"net/http"
	"strconv"
)

type QueryParse interface {
	Parse(r *http.Request)
}

type AvailableFriendsRequest struct {
	LastViewedUserId int    `validate:"numeric"`
	Name             string `validate:"required_with=Surname"`
	Surname          string `validate:"required_with=Database"`
}

func (a *AvailableFriendsRequest) Parse(r *http.Request) {
	lastViewedUserId, err := strconv.Atoi(r.URL.Query().Get("lastViewedUserId"))
	if nil != err {
		lastViewedUserId = 0
	}
	a.LastViewedUserId = lastViewedUserId

	a.Name = r.URL.Query().Get("name")
	a.Surname = r.URL.Query().Get("surname")
}
