package tasks

import (
	"encoding/json"
	"net/http"

	"github.com/allanger/pomoday-backend/tools/errors"
	"github.com/allanger/pomoday-backend/tools/hasher"
	"github.com/allanger/pomoday-backend/tools/logger"
	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
)

var log = logger.NewLogger("pomoday-service").Entry

//List: Get a list of tasks
func List(w http.ResponseWriter, r *http.Request) {
	log.Info("List endpoint hit")

	ctx := r.Context()
	// Validate data model
	log.Info("Fetching tasks")
	tasks, err := list(ctx)
	if err != nil {
		log.Error(err)
		render.Render(w, r, errors.ErrServer(err))
		return
	}
	render.Respond(w, r, newTasksListResponse(tasks))
	return
}

//TODO: Refactor
type HttpReq struct {
	Tasks []Task `json:"tasks"`
}

func Update(w http.ResponseWriter, r *http.Request) {
	log.Info("Update endpoint hit")
	ctx := r.Context()

	var httpReq HttpReq
	err := json.NewDecoder(r.Body).Decode(&httpReq)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, v := range httpReq.Tasks {
		t := v
		log.Info(v.Archived)
		log.Info("Validating tasks")
		if err := t.Validate(); err != nil {
			log.Error(err)
			render.Render(w, r, errors.ErrValidation(errTaskValidation, err.(validation.Errors)))
			return
		}

		log.Info("Updating tasks")
		err := update(ctx, t)
		if err != nil {
			log.Error(err)
			render.Render(w, r, errors.ErrServer(err))
			return
		}
	}

	render.Respond(w, r, newOkResponse())
	return
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Info("CreateUser endpoint hit")
	ctx := r.Context()
	u := &User{}
	if err := render.Bind(r, u); err != nil {
		log.Error(err)
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}
	u.UserID = uuid.New().String()
	u.Password = hasher.Encrypt(u.Password)
	err := createUser(ctx, u)
	if err != nil {
		log.Error(err)
		render.Render(w, r, errors.ErrServer(err))
		return
	}
	render.Respond(w, r, newOkResponse())
	return
}
