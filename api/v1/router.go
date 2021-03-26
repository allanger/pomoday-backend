package tasks

import (
	"github.com/allanger/pomoday-backend/middleware/auth"
	"github.com/allanger/pomoday-backend/tools/responses"
	"github.com/go-chi/chi"
)



// Router for Tasks API endpoints
func Router() *chi.Mux {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Use(auth.BasicAuth("asdasd", map[string]string{"admin": "admin"}))
		r.Get("/list", List)
		r.Put("/list", Update)
	})
	//TODO: Refactor
	r.Route("/user", func(r chi.Router) {
		r.Post("/", CreateUser)
	})
	return r
}

func newOkResponse() *responses.SuccessResponse {
	resp := &responses.SuccessResponse{
		// StatusText: "Succes",
		// AppCode:    200,
		Data:       "updated",
	}
	return resp

}

func newTasksListResponse(data []*Task) *responses.SuccessResponse {
	resp := &responses.SuccessResponse{
		// StatusText: "Succes",
		// AppCode:    200,
		Data:       data,
	}
	return resp

}
