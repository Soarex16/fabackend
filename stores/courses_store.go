package stores

import (
	"github.com/lib/pq"
	"github.com/soarex16/fabackend/domain"
)

type CoursesStore interface {
	GetAll() (*[]domain.Course, error)
}

type PqCoursesStore struct {
	Store
}

func (s *PqCoursesStore) GetAll() (*[]domain.Course, error) {
	const query = `
		SELECT * 
		FROM courses;
	`

	rows, err := s.DB.Query(query)

	if err != nil {
		logDBErr(err, query, "Error, while querying courses from DB")
		return nil, err
	}

	defer rows.Close()

	resultSet := make([]domain.Course, 0)
	for rows.Next() {
		course := domain.Course{}

		err := rows.Scan(&course.Label, &course.Description, pq.Array(&course.Exercises))
		if err != nil {
			logDBErr(err, query, "Error, while fetching row from query")
		}

		resultSet = append(resultSet, course)
	}

	if err = rows.Err(); err != nil {
		logDBErr(err, query, "Error, while querying data from DB")
		return nil, err
	}

	return &resultSet, nil
}
