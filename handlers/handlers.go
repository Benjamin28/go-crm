package handlers

import (
	"encoding/json"
	"net/http"

	_ "github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Healthcheck .
func (c *Controller) Healthcheck(w http.ResponseWriter, req *http.Request) {
	if err := c.PingDB(); err != nil {
		// Return error.
		logrus.WithError(err).Debugf("Failed to ping DB.")
		w.WriteHeader(http.StatusTeapot)
	}
	// Success.
	w.WriteHeader(http.StatusNoContent)
}

// GetPeople .
func (c *Controller) GetClients(w http.ResponseWriter, r *http.Request) {
	// json.NewEncoder(w).Encode(&people)
	w.WriteHeader(404)
}

// GetPerson .
func (c *Controller) GetClient(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	// for _, item := range people {
	// 	if item.ID == params["id"] {
	// 		json.NewEncoder(w).Encode(item)
	// 		return
	// 	}
	// }
	w.WriteHeader(404)
	// json.NewEncoder(w).Encode(&Person{})
}

// GetStatus .
func (c *Controller) GetStatus(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	// for _, item := range people {
	// 	if item.ID == params["id"] {
	// 		json.NewEncoder(w).Encode(item.Status)
	// 		return
	// 	}
	// }
	//json.NewEncoder(w).Encode("")
	w.WriteHeader(404)
}

// SetStatus .
func (c *Controller) SetStatus(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	// for _, item := range people {
	// 	if item.ID == params["id"] {
	// 		item.Status = params["newstatus"]
	// 		json.NewEncoder(w).Encode(people)
	// 		return
	// 	}
	//}
	w.WriteHeader(404)
}

// CreateClient .
func (c *Controller) CreateClient(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	var client Client
	_ = json.NewDecoder(r.Body).Decode(&client)
	//people = append(people, person)
	//insert client into DB
	w.WriteHeader(200)
}

// DeleteClient .
func (c *Controller) DeleteClient(w http.ResponseWriter, r *http.Request) {
}