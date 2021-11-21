package database

import (
	"log"
	"task-api/internal/model"
)

func (pg *PostgresDB) CreateTask(task model.Task) error {
	row := pg.db.QueryRow("INSERT INTO tasks (user_id, name, description, start_date, end_date, xp) "+
		"VALUES ($1, $2, $3, $4, $5, $6) RETURNING id ",
		task.UserId, task.Name, task.Description, task.StartDate, task.EndDate, task.Xp)

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

		if err := rows.Scan(&task.Id, &task.UserId, &task.Name, &task.Xp, &task.Description, &task.StartDate, &task.EndDate); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (pg *PostgresDB) UpdateUserXP(userId int, xp int) error {
	_, err := pg.db.Exec("UPDATE junction21.public.users SET xp=($1) WHERE id=($2)", xp, userId)

	return err
}

func (pg *PostgresDB) GetTaskByUser(userId int, taskName string) (model.Task, error) {
	row := pg.db.QueryRow("SELECT * FROM junction21.public.tasks WHERE name=($1) AND user_id=($2)", taskName, userId)

	var task model.Task
	if err := row.Scan(&task.Id, &task.UserId, &task.Name, &task.Xp, &task.Description, &task.StartDate, &task.EndDate); err != nil {
		return model.Task{}, nil
	}

	return task, nil
}

func (pg *PostgresDB) DeleteTaskByUser(taskName string, userId int) error {
	_, err := pg.db.Exec("DELETE FROM junction21.public.tasks WHERE user_id=($1) AND name=($2)", userId, taskName)

	return err
}

func (pg *PostgresDB) GetUserByName(name string) (model.User, error) {
	row := pg.db.QueryRow("SELECT * FROM junction21.public.users WHERE name=($1)", name)

	var user model.User
	if err := row.Scan(&user.Id, &user.Name, &user.Xp); err != nil {
		return model.User{}, err
	}

	return user, nil
}
