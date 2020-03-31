package context

import (
	"database/sql"
)

// Context - global application context
type Context struct {
	DB *sql.DB
}

// RequestContextKey - type definition for key of injected data into request.Context
type RequestContextKey string
