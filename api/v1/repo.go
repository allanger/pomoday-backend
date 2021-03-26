package tasks

import (
	"context"
	"encoding/json"

	"github.com/allanger/pomoday-backend/middleware/auth"
	d "github.com/allanger/pomoday-backend/third_party/postgres"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	createTaskRequest  = "INSERT INTO tasks (userid, id, uuid, archived, tag, title, status, lastaction, logs) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	createUserRequest  = "INSERT INTO users (id, username, password) VALUES ($1, $2, $3)"
	getTasksRequest    = "SELECT * FROM tasks WHERE userid=$1"
	getTaskByIDRequest = "SELECT title FROM tasks WHERE uuid = $1"
	updateTaskRequest  = "UPDATE tasks SET archived=$1, tag=$2, title=$3, status=$4, lastaction=$5, logs=$6 WHERE uuid = $1;"
)

var (
	database *pgxpool.Pool
)

func update(ctx context.Context, t Task) (err error) {
	log.Info("Adding to DB")
	database := d.Pool()
	// logs, err := json.Marshal(t.Logs)
	logs, err := json.MarshalIndent(t.Logs, "", "\t")
	if err != nil {
		log.Error(err)
		return err
	}
	log.Infof("%v", t.Logs)
	var title string
	if err := database.QueryRow(ctx, getTaskByIDRequest, t.UUID).Scan(&title); err != nil {
		_, err = database.Exec(ctx, createTaskRequest, auth.UserID, t.ID, t.UUID, t.Archived, t.Tag, t.Title, t.Status, t.Lastaction, logs)
	} else {
		_, err = database.Exec(ctx, updateTaskRequest, t.UUID, t.Archived, t.Tag, t.Title, t.Status, t.Lastaction, logs)
	}

	if err != nil {
		log.Error(err)
		return err
	}
	return
}

func list(ctx context.Context) ([]*Task, error) {
	database := d.Pool()
	var tasks []*Task
	if err := pgxscan.Select(ctx, database, &tasks, getTasksRequest, auth.UserID); err != nil {
		return nil, err
	}
	log.Info(tasks)
	return tasks, nil
}

func createUser(ctx context.Context, u *User) (err error) {
	database := d.Pool()
	_, err = database.Exec(ctx, createUserRequest, u.UserID, u.Username, u.Password)
	if err != nil {
		return err
	}
	return
}
