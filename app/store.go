package app

import (
	"database/sql"
	"github.com/soarex16/fabackend/auth"
	"github.com/soarex16/fabackend/stores"
)

// Store - represents application db
type Store struct {
	// Hack, in if you wants custom request (and for closing db connection when shutting down)
	DB *sql.DB

	Achievements stores.AchievementsStore
	Courses      stores.CoursesStore
	Users        stores.UsersStore

	Sessions *auth.SessionStore
}

// NewStore - initializes new instances of stores
func NewStore(db *sql.DB) *Store {
	s := stores.Store{DB: db}

	return &Store{
		DB:           s.DB,
		Achievements: &stores.PqAchievementsStore{Store: s},
		Courses:      &stores.PqCoursesStore{Store: s},
		Users:        &stores.PqUsersStore{Store: s},

		Sessions: auth.NewSessionStore(),
	}
}
