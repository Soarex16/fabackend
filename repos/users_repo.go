package repos

import (
	"database/sql"

	"github.com/sirupsen/logrus"
	"github.com/soarex16/fabackend/domain"
)

// AddUser - creates new user
func AddUser(db *sql.DB, u *domain.User) (sql.Result, error) {
	query := `
		INSERT INTO users(username, email, password) VALUES
			($1, $2, $3);
	`

	stmt, err := db.Prepare(query)

	res, err := stmt.Exec(u.Username, u.Email, u.Password)

	if err != nil {
		logrus.
			WithField("query", query).
			Errorf("Error, while inserting data into DB: %v", err)
		return nil, err
	}

	return res, nil
}

// FindUserByName - finds user with specified username
func FindUserByName(db *sql.DB, username string) (*domain.User, error) {
	return nil, nil
}

// UpdateUser - updates user information
func UpdateUser(db *sql.DB, u *domain.User) (sql.Result, error) {
	query := `
		UPDATE users SET
			email = $1,
			password = $2
		WHERE 
			username = $3;
	`

	stmt, err := db.Prepare(query)

	res, err := stmt.Exec(u.Email, u.Password, u.Username)

	if err != nil {
		logrus.
			WithField("query", query).
			Errorf("Error, while inserting data into DB: %v", err)
		return nil, err
	}

	return res, nil
}

// DeleteUser - removes user from db, returns success even if it doesn't exists
func DeleteUser(db *sql.DB, username string) (sql.Result, error) {
	query := `
		DELETE FROM users
		WHERE username = $1;
	`

	stmt, err := db.Prepare(query)

	res, err := stmt.Exec(username)

	if err != nil {
		logrus.
			WithField("query", query).
			Errorf("Error, while inserting data into DB: %v", err)
		return nil, err
	}

	return res, nil
}
