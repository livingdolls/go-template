package port

import (
	"io"

	"github.com/jmoiron/sqlx"
)

type DatabasePort interface {
	io.Closer
	GetDatabase() *sqlx.DB
}
