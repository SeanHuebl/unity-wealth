package config

import (
	"database/sql"

	"github.com/seanhuebl/unity-wealth/internal/interfaces"
)

type ApiConfig struct {
	Port        string
	Queries     interfaces.Quierier
	Database    *sql.DB
	TokenSecret string
	Auth        interfaces.AuthInterface
}
