package sql

import (
	"github.com/google/uuid"
	"github.com/soarex16/fabackend/domain"
)

type AchievementsStore interface {
	GetByUsername(username string) (*[]domain.Achievement, error)
	Add(userId uuid.UUID, ach *domain.Achievement) (bool, error)
	AddByUsername(username string, ach *domain.Achievement) (bool, error)
}

type PqAchievementsStore struct {
	Store
}

func (s *PqAchievementsStore) GetByUsername(username string) (*[]domain.Achievement, error) {
	const query = `
		SELECT achievements.date, achievements.description, achievements.iconcolor, achievements.id, achievements.price, achievements.title
		FROM achievements
		WHERE achievements.userid = (
			SELECT id
			FROM users
			WHERE username = $1
		);
	`

	stmt, _ := s.DB.Prepare(query)
	rows, err := stmt.Query(username)

	if err != nil {
		logDBErr(err, query, "")
		return nil, err
	}

	resultSet := make([]domain.Achievement, 0)
	for rows.Next() {
		ach := domain.Achievement{}
		err := rows.Scan(&ach.Date, &ach.Description, &ach.IconColor, &ach.ID, &ach.Price, &ach.Title)

		if err != nil {
			logDBErr(err, query, "Error, while fetching row from query")
		}

		resultSet = append(resultSet, ach)
	}

	if err = rows.Err(); err != nil {
		logDBErr(err, query, "Error, while querying data from DB")
		return nil, err
	}

	return &resultSet, nil
}

func (s *PqAchievementsStore) Add(userId uuid.UUID, ach *domain.Achievement) (bool, error) {
	const query = `
		INSERT INTO achievements(date, description, iconcolor, price, title, userid) VALUES
			($1, $2, $3, $4, $5);
	`

	stmt, _ := s.DB.Prepare(query)
	res, err := stmt.Exec(ach.Date, ach.Description, ach.IconColor, ach.Price, ach.Title, userId)

	if err != nil {
		logDBErr(err, query, "Error, while inserting data into DB")
		return false, err
	}

	affectedRows, err := res.RowsAffected()

	if err != nil {
		logDBErr(err, query, "")
		return false, err
	}

	return affectedRows == 1, nil
}

func (s *PqAchievementsStore) AddByUsername(username string, ach *domain.Achievement) (bool, error) {
	const query = `
		INSERT INTO achievements(date, description, iconcolor, price, title, userid) VALUES
			($1, $2, $3, $4, $5, SELECT id
				FROM users
				WHERE username = $6);
	`

	stmt, _ := s.DB.Prepare(query)
	res, err := stmt.Exec(ach.Date, ach.Description, ach.IconColor, ach.Price, ach.Title, username)

	affectedRows, err := res.RowsAffected()

	if err != nil {
		logDBErr(err, query, "")
		return false, err
	}

	return affectedRows == 1, nil
}
