package interfaces

import (
	"encoding/json"
	"fmt"
	router "github.com/takashabe/go-router"
	"log"
	"net/http"
	"sample-go-ddd/application"
	"sample-go-ddd/domain/entity"
	"sample-go-ddd/domain/repository"
)

type Handler struct {
	repo repository.UserRepository
}

func jsonResponse(response http.ResponseWriter, code int, obj interface{}) {
	jsObj, err := json.Marshal(obj)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(code)
	response.Write(jsObj)

}

func (h Handler) getUser(response http.ResponseWriter, request *http.Request, id int) {
	ctx := request.Context()

	app := application.UserRepo{
		Repo: h.repo,
	}
	usr, err := app.GetUser(ctx, id)
	if err != nil {
		http.Error(response, "invalid user id", http.StatusNotFound)
		return
	}
	jsonResponse(response, http.StatusOK, usr)
}

func (h Handler) getUsers(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	app := application.UserRepo{
		Repo: h.repo,
	}
	users, err := app.Users(ctx)
	if err != nil {
		http.Error(response, "invalid user list request", http.StatusBadRequest)
		return
	}
	type payload struct {
		Users []*entity.User `json:"users"`
	}
	jsonResponse(response, http.StatusOK, payload{Users: users})
}

func (h Handler) createUser(response http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	type payload struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	var pld payload
	if err := json.NewDecoder(r.Body).Decode(&pld); err != nil {
		http.Error(response, "invalid request", http.StatusBadRequest)
		return
	}
	app := application.UserRepo{
		Repo: h.repo,
	}
	if err := app.Save(ctx, pld.Name, pld.Email); err != nil {
		http.Error(response, "invalid request", http.StatusBadRequest)
		return
	}
	jsonResponse(response, http.StatusCreated, nil)
}

func (h Handler) Routes() *router.Router {
	r := router.NewRouter()
	r.Get("/user/:id", h.getUser)
	r.Get("/users", h.getUsers)
	r.Post("/user", h.createUser)
	return r
}

func NewHandler(repo repository.UserRepository) Handler{
	return Handler{
		repo: repo,
	}
}

func (h Handler) Run(port int) error {
	log.Printf("Started: http://localhost:%d/", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), h.Routes())
}
