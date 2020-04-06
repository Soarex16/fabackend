package stores

import (
	"github.com/soarex16/fabackend/domain"
)

type UsersStore interface {
	FindByName(username string) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	Add(u *domain.User) (bool, error)
	Update(u *domain.User) (bool, error)
	Delete(username string) (bool, error)
}

type PqUsersStore struct {
	Store
}

// Add - creates new user
func (s *PqUsersStore) Add(u *domain.User) (bool, error) {
	query := `
		INSERT INTO users(username, email, password) VALUES
			($1, $2, $3);
	`

	stmt, err := s.DB.Prepare(query)

	res, err := stmt.Exec(u.Username, u.Email, u.Password)

	if err != nil {
		logDBErr(err, query, "Error, while inserting user into DB")
		return false, err
	}

	affectedRows, err := res.RowsAffected()

	if err != nil {
		logDBErr(err, query, "")
		return false, err
	}

	return affectedRows == 1, nil
}

// FindByName - finds user with specified username
// NOTE: Also returns password!
func (s *PqUsersStore) FindByName(username string) (*domain.User, error) {
	const query = `
		SELECT id, username, email, password
		FROM users
		WHERE users.username = $1;
	`

	u := domain.User{}
	err := s.DB.QueryRow(query, username).Scan(&u.ID, &u.Username, &u.Email, &u.Password)

	if err != nil {
		logDBErr(err, query, "Error, while querying user from DB")
		return nil, err
	}

	return &u, nil
}

// FindByName - finds user with specified email
// NOTE: Also returns password!
func (s *PqUsersStore) FindByEmail(email string) (*domain.User, error) {
	const query = `
		SELECT id, username, email, password
		FROM users
		WHERE users.email = $1;
	`

	u := domain.User{}
	err := s.DB.QueryRow(query, email).Scan(&u.ID, &u.Username, &u.Email, &u.Password)

	if err != nil {
		logDBErr(err, query, "Error, while querying user from DB")
		return nil, err
	}

	return &u, nil
}

// Update - updates user information
func (s *PqUsersStore) Update(u *domain.User) (bool, error) {
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
		logDBErr(err, query, "Error, while updating user in DB")
		return false, err
	}

	affectedRows, err := res.RowsAffected()

	if err != nil {
		logDBErr(err, query, "")
		return false, err
	}

	return affectedRows == 1, nil
}

// Delete - removes user from db, returns success even if it doesn't exists
func (s *PqUsersStore) Delete(username string) (bool, error) {
	query := `
		DELETE FROM users
		WHERE username = $1;
	`

	stmt, err := s.DB.Prepare(query)

	res, err := stmt.Exec(username)

	if err != nil {
		logDBErr(err, query, "Error, while deleting user from DB")
		return false, err
	}

	affectedRows, err := res.RowsAffected()

	if err != nil {
		logDBErr(err, query, "")
		return false, err
	}

	return affectedRows == 1, nil
}
