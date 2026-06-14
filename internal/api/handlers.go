package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/inflame-ue/godo/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (api *API) HandleGetTODOS(w http.ResponseWriter, r *http.Request) {
	todos, err := api.db.GetAllTODO(r.Context())
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(todos)
	if err != nil {
		log.Printf("marhsalling todos into json: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	_, err = w.Write(data)
	if err != nil {
		log.Printf("writing the json body: %v", err)
	}
}

func (api *API) HandlePostTODO(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		log.Printf("invalid content type received: %s", r.Header.Get("Content-Type"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var todo models.TODO
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		log.Printf("decoding the request body into todo struct: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	createdTODO, err := api.db.InsertTODO(r.Context(), todo.Title, todo.Completed)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(createdTODO)
	if err != nil {
		log.Printf("marshalling the created todo: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(data)
	if err != nil {
		log.Printf("writing the todo to json: %v", err)
	}
}

func (api *API) HandleGetTODOByID(w http.ResponseWriter, r *http.Request) {
	todoID := r.PathValue("id")
	todoObjectID, err := bson.ObjectIDFromHex(todoID)
	if err != nil {
		log.Printf("converting to object id from string: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	todo, err := api.db.GetTODOByID(r.Context(), todoObjectID)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(todo)
	if err != nil {
		log.Printf("marshalling the todo into json: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	_, err = w.Write(data)
	if err != nil {
		log.Printf("writing json to response body: %v", err)
	}
}

func (api *API) HandlePutTODOByID(w http.ResponseWriter, r *http.Request) {
	todoID := r.PathValue("id")
	todoObjectID, err := bson.ObjectIDFromHex(todoID)
	if err != nil {
		log.Printf("converting to object id from string: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var todo models.TODO
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		log.Printf("decoding the request body into todo struct: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updatedTODO, err := api.db.UpdateTODOByID(r.Context(), todoObjectID, todo.Title, todo.Completed)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(updatedTODO)
	if err != nil {
		log.Printf("marshalling the updated todo into json: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	_, err = w.Write(data)
	if err != nil {
		log.Printf("writing json to response body: %v", err)
	}
}

func (api *API) HandleDeleteTODOByID(w http.ResponseWriter, r *http.Request) {
	todoID := r.PathValue("id")
	todoObjectID, err := bson.ObjectIDFromHex(todoID)
	if err != nil {
		log.Printf("converting to object id from string: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = api.db.DeleteTODOByID(r.Context(), todoObjectID)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	msg := struct {
		Msg string `json:"message"`
	}{
		Msg: fmt.Sprintf("TODO with id %s was deleted successfully", todoID),
	}
	data, err := json.Marshal(msg)
	if err != nil {
		log.Printf("marshalling the success message into json: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	_, err = w.Write(data)
	if err != nil {
		log.Printf("writing json to response body: %v", err)
	}
}
