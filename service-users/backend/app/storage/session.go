package storage

import (
	"github.com/gorilla/sessions"
	"net/http"
	"service-users/app/config"
)

const SessionUserIdKey = "user_id"

type SessionStorage struct {
	CookieStore *sessions.CookieStore
}

func InitSession() SessionStorage {
	store := sessions.NewCookieStore([]byte(config.Env.Server.SessionKey))

	store.Options = &sessions.Options{
		Domain:   config.Env.Server.Host,
		Path:     "/",
		MaxAge:   3600 * 24 * 7, // 1 week
		HttpOnly: true,
	}

	session := SessionStorage{
		CookieStore: store,
	}

	return session
}

func (s *SessionStorage) GetSession(r *http.Request) *sessions.Session {
	session, _ := s.CookieStore.Get(r, "_data")

	return session
}
