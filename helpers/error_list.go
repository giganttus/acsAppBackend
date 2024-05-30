package helpers

import (
	"errors"
)

var (
	ErrEnvLoad      = errors.New("environment variables load failed")
	ErrEnvConvert   = errors.New("environment variables convert failed")
	ErrDbConn       = errors.New("database connection failed")
	ErrRepo         = errors.New("repo error")
	ErrService      = errors.New("service error")
	ErrForbidden    = errors.New("forbidden")
	ErrUnauthorized = errors.New("unauthorized")
)

type ResponseBody struct {
	Response string `json:"response"`
}
