package sql

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
)

type Store struct {
	DB *sql.DB
}

func logDBErr(err error, query string, errMsg string) {
	if errMsg != "" {
		errMsg = fmt.Sprintf("Error, while executing query: %v", err)
	}

	logrus.
		WithField("query", query).
		WithField("err", err).
		Error(errMsg)
}
