package infra

import "github.com/gofiber/fiber/v2/middleware/session"

type SessionStore struct {
	*session.Store
}

var store = &SessionStore{}

func InitSessionStorage() {
	store.Store = session.New()
}

func GetSessionStore() *SessionStore {
	return store
}
