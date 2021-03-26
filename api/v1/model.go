package tasks

import (
	"errors"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Logs struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}

// Task is a data struct for chain json
type Task struct {
	ID         int    `json:"id"`
	UUID       string `json:"uuid"`
	Archived   bool   `json:"archived"`
	Tag        string `json:"tag"`
	Title      string `json:"title"`
	Status     int    `json:"status"`
	Lastaction int64  `json:"lastaction"`
	Logs       []struct {
		Start int64 `json:"start"`
		End   int64 `json:"end"`
	} `json:"logs"`
	// Logs []string `json:"logs"`
	// Logs []map[string]int64 `json:"logs"`
}

type User struct {
	UserID   string `json:"userID" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

var (
	errTaskValidation = errors.New("Good: json validation failed")
)

// Validate is to validate request params for goods
func (in *Task) Validate() error {
	return validation.ValidateStruct(in,
		validation.Field(&in.ID, validation.Required),

		validation.Field(&in.UUID, validation.Required),
		validation.Field(&in.UUID, is.UUID),

		// validation.Field(&in.Archived, validation.Required),

		validation.Field(&in.Tag, validation.Required),

		validation.Field(&in.Title, validation.Required),

		validation.Field(&in.Status, validation.Required),

		validation.Field(&in.Lastaction, validation.Required),

		// validation.Field(&in.Logs, validation.Required),
	)
}

// Bind implemintation for Tasks model
func (in *Task) Bind(r *http.Request) error {
	return nil
}

// Bind implemintation for Tasks model
func (u *User) Bind(r *http.Request) error {
	return nil
}
