package sql

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"github.com/soarex16/fabackend/domain"
)

type AchievementsStore struct {
	Store
}

func (s *AchievementsStore) GetByUsername(username string) (*[]domain.Achievement, error) {
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
		logrus.
			WithField("query", query).
			Errorf("Error, while querying data from DB: %v", err)
		return nil, err
	}

	resultSet := make([]domain.Achievement, 0)
	for rows.Next() {
		ach := domain.Achievement{}
		err := rows.Scan(&ach.Date, &ach.Description, &ach.IconColor, &ach.ID, &ach.Price, &ach.Title)

		if err != nil {
			logrus.
				WithField("query", query).
				Errorf("Error, while fetching row from query: %v", err)
		}

		resultSet = append(resultSet, ach)
	}

	if err = rows.Err(); err != nil {
		logrus.
			WithField("query", query).
			Errorf("Error, while querying data from DB: %v", err)
		return nil, err
	}

	return &resultSet, nil
}

func (s *AchievementsStore) Add(username string, ach *domain.Achievement) (sql.Result, error) {
	const query = `
		INSERT INTO achievements(date, description, iconcolor, price, title, userid) VALUES
			($1, $2, $3, $4, $5, SELECT id
				FROM users
				WHERE username = $6);
	`

	stmt, _ := s.DB.Prepare(query)
	res, err := stmt.Exec(ach.Date, ach.Description, ach.IconColor, ach.Price, ach.Title, username)

	if err != nil {
		logrus.
			WithField("query", query).
			Errorf("Error, while inserting data into DB: %v", err)
		return nil, err
	}

	return res, nil
}
