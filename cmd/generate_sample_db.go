package cmd

import (
	"database/sql"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	// register postgres driver
	_ "github.com/lib/pq"
)

var generateSampleDbCmd = &cobra.Command{
	Use:   "generate_sample_db",
	Short: "Inserts sample data into database",
	RunE: func(cmd *cobra.Command, args []string) error {
		connStr := viper.GetString("dbConnectionString")
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
			INSERT INTO users VALUES
				('abb9cb3c-d405-4a5b-8f68-0a2ffb9e47d3', 'sample', 'sample@mail.com', 'af2bdbe1aa9b6ec1e2ade1d694f41fc71a831d0268e9891562113d8a62add1bf'),
				('5eb94ee3-7d81-4303-9988-29847dd136ad', 'example', 'example@mail.com', '50d858e0985ecc7f60418aaf0cc5ab587f42c2570a884095a9e8ccacd0f6545c')
				ON CONFLICT (id) DO NOTHING;
		`

		_, err = db.Exec(addUsersSQL)

		if err != nil {
			logrus.Errorf("Error while adding users into table: %v", err)
			return err
		}

		const addAchievementsSQL = `
			INSERT INTO achievements values
				('2b91b28e-69d1-41fc-8448-b2fe3aeb21a6', 'abb9cb3c-d405-4a5b-8f68-0a2ffb9e47d3', '2019-01-08 04:05:06', 1, 'dumbbell', '#00c853', 'Тренировка "Пинаем пиналку" завершена', 'Ну ты так себе потренирвоался'),
				('a1ce4a07-d8e9-42d5-b264-7669543e5b11', '5eb94ee3-7d81-4303-9988-29847dd136ad', '2019-01-08 12:25:00', 1, 'dumbbell', '#00c853', 'Тренировка "Пинаем пиналку" завершена', 'Ну ты так себе потренирвоался'),
				('7e4fd93e-b7f2-4f4a-8c6d-45a9badd45d9', 'abb9cb3c-d405-4a5b-8f68-0a2ffb9e47d3', '2019-01-08 16:34:09', 1, 'dumbbell', '#00c853', 'Тренировка "Пинаем пиналку" завершена', 'Ну ты так себе потренирвоался'),
				('fdb11113-03ac-4e39-8df5-c3b536d2e0c3', 'abb9cb3c-d405-4a5b-8f68-0a2ffb9e47d3', '2019-01-08 08:13:46', 1, 'dumbbell', '#00c853', 'Тренировка "Пинаем пиналку" завершена', 'Ну ты так себе потренирвоался'),
				('83277397-13b2-4339-9469-61ccf58f5f53', '5eb94ee3-7d81-4303-9988-29847dd136ad', '2019-01-08 19:01:17', 1, 'dumbbell', '#00c853', 'Тренировка "Пинаем пиналку" завершена', 'Ну ты так себе потренирвоался')
				ON CONFLICT (id) DO NOTHING;
			`

		_, err = db.Exec(addAchievementsSQL)

		if err != nil {
			logrus.Errorf("Error while adding achievements into table: %v", err)
			return err
		}

		const addCoursesSQL = `
			INSERT INTO courses VALUES
				('Фитнес с гантельками 101', 'Базовая тренировка по фитнессу с гантелями для отличного начала дня!', '{"Выпад 1", "Махи 1", "Я не знаю как это назвать D:", "Учимся качать матрасс", "Танцуем!", "Уклонение от пуль"}'),
				('Фитнес с гантельками 228', 'Для тех, кому мало', '{"Качау", "Болеем", "Целуйтей", "Я не знаю как это назвать 2", "У меня нет денег на штангу", "Михалыч", "Михалыч 2", "Гантельки"}'),
				('Утренняя разминка', 'Разминка на утро перед тяжким рабочим днем', '{"Я не знаю как это назвать D:", "Танцуем!", "Уклонение от пуль", "Болеем"}')
				ON CONFLICT (label) DO NOTHING;
			`

		_, err = db.Exec(addCoursesSQL)

		if err != nil {
			logrus.Errorf("Error while adding courses into table: %v", err)
			return err
		}

		logrus.Info("Database successfully filled with test data")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateSampleDbCmd)
}
