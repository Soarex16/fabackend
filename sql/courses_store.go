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
	/**
	TODO: почему-то при запросе курсов в массиве упражнений только первое
	[
	    {
	        "label": "Фитнес с гантельками 101",
	        "description": "Базовая тренировка по фитнессу с гантелями для отличного начала дня!",
	        "exercises": [
	            "Выпад 1",
	            "Махи 1",
	            "Я не знаю как это назвать D:",
	            "Учимся качать матрасс",
	            "Танцуем!",
	            "Уклонение от пуль"
	        ]
	    },
	    {
	        "label": "Фитнес с гантельками 228",
	        "description": "Для тех, кому мало боли",
	        "exercises": [
	            "Качау",
	            "Болеем",
	            "Целуйтей",
	            "Я не знаю как это назвать 2",
	            "У меня нет денег на штангу",
	            "Михалыч",
	            "Михалыч 2",
	            "Гантельки"
	        ]
	    },
	    {
	        "label": "Утренняя разминка",
	        "description": "Разминка на утро перед тяжким рабочим днем",
	        "exercises": [
	            "Я не знаю как это назвать D:",
	            "Танцуем!",
	            "Уклонение от пуль",
	            "Болеем"
	        ]
	    }
	]
	*/

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
