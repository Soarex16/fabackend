package sql

import (
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/soarex16/fabackend/domain"
)

type CoursesStore struct {
	Store
}

func (s *CoursesStore) GetAll() (*[]domain.Course, error) {
	const query = `
		SELECT * 
		FROM courses;
	`

	rows, err := s.DB.Query(query)

	if err != nil {
		logrus.
			WithField("query", query).
			Errorf("Error, while querying data from DB: %v", err)
		return nil, err
	}

	defer rows.Close()

	resultSet := make([]domain.Course, 0)
	for rows.Next() {
		course := domain.Course{}

		err := rows.Scan(&course.Label, &course.Description, pq.Array(&course.Exercises))
		if err != nil {
			logrus.
				WithField("query", query).
				Errorf("Error, while fetching row from query: %v", err)
		}

		resultSet = append(resultSet, course)
	}

	if err = rows.Err(); err != nil {
		logrus.
			WithField("query", query).
			Errorf("Error, while querying data from DB: %v", err)
		return nil, err
	}

	return &resultSet, nil
}
