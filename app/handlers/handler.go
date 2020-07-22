package handlers

import (
	"github.com/badThug/otus-social-network/app/database"
)

type Handler struct {
	Db *database.Connection
}
