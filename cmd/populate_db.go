package cmd

import (
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	// register postgres driver
	_ "github.com/lib/pq"
)

var populateDbCmd = &cobra.Command{
	Use:   "populate_db",
	Short: "Populates tables in database",
	RunE: func(cmd *cobra.Command, args []string) error {
		connStr := viper.GetString("dbConnectionString")

		if connStr == "" {
			connStr = viper.GetString("DATABASE_URL")

			if connStr == "" {
				return fmt.Errorf("'dbConnectionString' must be specified in configuration file")
			}
		}
		db, err := sql.Open("postgres", connStr)

		if err != nil {
			logrus.Error("Cannot open connection to database")
			return err
		}

		if err := db.Ping(); err != nil {
			logrus.Error("Cannot ping database")
			return err
		}

		defer db.Close()

		const addUsersSQL = `
			CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

			CREATE TABLE users (
				id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
				username varchar(50) UNIQUE,
				email varchar(320),
				-- SHA-1 hashed
				password char(64) NOT NULL
			);
		`

		_, err = db.Exec(addUsersSQL)

		if err != nil {
			logrus.Errorf("Error while creating users table: %v", err)
			return err
		}

		const addAchievementsSQL = `
			CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

			CREATE TABLE achievements (
				id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
				-- Пользователь, который получил это достижение
				userId uuid REFERENCES users(id),
				-- Дата в формате unix timestamp
				date timestamp,
				-- Вес достижения (нужно для календаря)
				price int,
				-- Имя иконки для отображения в карточке
				iconName varchar(100),
				iconColor varchar(20),
				description varchar(300),
				title varchar(100)
			);
		`

		_, err = db.Exec(addAchievementsSQL)

		if err != nil {
			logrus.Errorf("Error while creating achievements table: %v", err)
			return err
		}

		const addCoursesSQL = `
			CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

			CREATE TABLE courses (
				label varchar(100) PRIMARY KEY NOT NULL,
				description varchar(300),
				exercises varchar(100)[]
			);
		`

		_, err = db.Exec(addCoursesSQL)

		if err != nil {
			logrus.Errorf("Error while creating courses table: %v", err)
			return err
		}

		logrus.Info("Database successfully created")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(populateDbCmd)
}
