package server

import (
	"encoding/json"
	"github.com/carinfinin/risk-assessor/internal/config"
	"github.com/carinfinin/risk-assessor/internal/model"
	"github.com/go-chi/chi/v5"
	"net/http"
)

//go:generate
type Service interface {
	CreateUser(clientData *model.ClientData) (*model.User, error)
}

type Router struct {
	Handler *chi.Mux
	Service Service
	Config  *config.Config
}

func NewRouter(cfg *config.Config, service Service) *Router {
	r := Router{
		Handler: chi.NewRouter(),
		Service: service,
		Config:  cfg,
	}
	r.configure()
	return &r
}

func (r *Router) configure() {
	r.Handler.Get("/test", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("tttttttttttttttt"))
	})

	r.Handler.Route("/api", func(router chi.Router) {
		router.Post("/user", r.RouteUser)
	})
}

func (r *Router) RouteUser(writer http.ResponseWriter, request *http.Request) {

	var clientData model.ClientData

	if err := json.NewDecoder(request.Body).Decode(&clientData); err != nil {
		http.Error(writer, "json invalid", http.StatusBadRequest)
		return
	}
	defer request.Body.Close()

	// validate

	writer.Header().Set("Content-Type", "application/json")

	user, err := r.Service.CreateUser(&clientData)
	if err != nil {
		http.Error(writer, "json error write", http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(writer).Encode(user); err != nil {
		http.Error(writer, "json error write", http.StatusInternalServerError)
	}

	/*
		{
		  "full_name": "Иванов Иван Иванович",
		  "phone": "+79991234567",
		  "email": "ivanov@example.com",
		  "income": 75000,
		  "number_passport": 4510123456,
		  "loan_amount": 500000,
		  "loan_term": 24
		}
	*/
}
