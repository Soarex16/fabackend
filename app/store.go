package app

import (
	"database/sql"
	stores "github.com/soarex16/fabackend/sql"
)

// Store - represents application db
type Store struct {
	// Hack, in if you wants custom request (and for closing db connection when shutting down)
	DB *sql.DB

	Achievements *stores.AchievementsStore
	Courses      *stores.CoursesStore
	Users        *stores.UsersStore
}

// NewStore - initializes new instances of stores
func NewStore(db *sql.DB) *Store {
	s := stores.Store{DB: db}

	return &Store{
		DB:           s.DB,
		Achievements: &stores.AchievementsStore{Store: s},
		Courses:      &stores.CoursesStore{Store: s},
		Users:        &stores.UsersStore{Store: s},
	}
}
