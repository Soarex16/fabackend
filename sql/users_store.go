package sql

import (
	"github.com/sirupsen/logrus"
	"github.com/soarex16/fabackend/domain"
)

type UsersStore struct {
	Store
}

// Add - creates new user
func (s *UsersStore) Add(u *domain.User) (bool, error) {
	query := `
		INSERT INTO users(username, email, password) VALUES
			($1, $2, $3);
	`

	stmt, err := s.DB.Prepare(query)

	res, err := stmt.Exec(u.Username, u.Email, u.Password)

	if err != nil {
		logrus.
			WithField("query", query).
			Errorf("Error, while inserting data into DB: %v", err)
		return false, err
	}

	affectedRows, _ := res.RowsAffected()

	return affectedRows == 1, nil
}

// FindByName - finds user with specified username
func (s *UsersStore) FindByName(username string) (*domain.User, error) {
	const query = `
		SELECT id, username, email
		FROM users
		WHERE users.username = $1;
	`

	u := domain.User{}
	err := s.DB.QueryRow(query, username).Scan(&u.ID, &u.Username, &u.Email)

	if err != nil {
		logrus.
			WithField("query", query).
			Errorf("Error, while querying data from DB: %v", err)
		return nil, err
	}

	return &u, nil
}

// Update - updates user information
func (s *UsersStore) Update(u *domain.User) (bool, error) {
	query := `
		UPDATE users SET
			email = $1,
			password = $2
		WHERE 
			username = $3;
	`

	stmt, err := s.DB.Prepare(query)

	res, err := stmt.Exec(u.Email, u.Password, u.Username)

	if err != nil {
		logrus.
			WithField("query", query).
			Errorf("Error, while inserting data into DB: %v", err)
		return false, err
	}

	affectedRows, _ := res.RowsAffected()

	return affectedRows == 1, nil
}

// Delete - removes user from db, returns success even if it doesn't exists
func (s *UsersStore) Delete(username string) (bool, error) {
	query := `
		DELETE FROM users
		WHERE username = $1;
	`

	stmt, err := s.DB.Prepare(query)

	res, err := stmt.Exec(username)

	if err != nil {
		logrus.
			WithField("query", query).
			Errorf("Error, while inserting data into DB: %v", err)
		return false, err
	}

	affectedRows, _ := res.RowsAffected()

	return affectedRows == 1, nil
}
