package api

import (
	"net/http"

	"github.com/inflame-ue/godo/internal/database"
)

type API struct {
	db  *database.MongoClient
	mux *http.ServeMux
}

func NewAPI(db *database.MongoClient) *API {
	api := API{
		db:  db,
		mux: http.NewServeMux(),
	}
	api.registerRoutes()
	return &api
}

func (api *API) registerRoutes() {
	api.mux.HandleFunc("GET /todos", api.HandleGetTODOS)
	api.mux.HandleFunc("POST /todos", api.HandlePostTODO)
	api.mux.HandleFunc("GET /todos/{id}", api.HandleGetTODOByID)
	api.mux.HandleFunc("PUT /todos/{id}", api.HandlePutTODOByID)
	api.mux.HandleFunc("DELETE /todos/{id}", api.HandleDeleteTODOByID)
}

func (api *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	api.mux.ServeHTTP(w, r)
}
