package sql

import "database/sql"

type Store struct {
	DB *sql.DB
}

//TODO: DRY DRY DRY!!! вынести повторяющийся код сюда в виде хелперов
