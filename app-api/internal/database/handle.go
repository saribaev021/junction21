package database

import (
	"log"
	"task-api/internal/model"
)

func (pg *PostgresDB) CreateTask(task model.Task) error {
	row := pg.db.QueryRow("INSERT INTO tasks (user_id, name, description, start_date, end_date) VALUES ($1, $2, $3, $4, $5) RETURNING id ",
		task.UserId, task.Name, task.Description, task.StartDate, task.EndDate)

	if err := row.Scan(&task.Id); err != nil {
		return err
	}
	log.Printf("task saved with id: %v", task.Id)

	return nil
}

func (pg *PostgresDB) GetUserIdByName(name string) (int, error) {
	var userId int

	row := pg.db.QueryRow("SELECT id FROM junction21.public.users WHERE name=($1)", name)
	if err := row.Scan(&userId); err != nil {
		return 0, err
	}

	return userId, nil
}

func (pg *PostgresDB) GetUserTasks(userId int) ([]model.Task, error) {
	rows, err := pg.db.Query("SELECT * FROM junction21.public.tasks WHERE user_id=($1)", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]model.Task, 0)

	for rows.Next() {
		var task model.Task

		if err := rows.Scan(&task.Id, &task.UserId, &task.Name, &task.Description, &task.EndDate, &task.StartDate); err != nil {
			return tasks, nil
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}
