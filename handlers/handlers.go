package handlers

import (
	"encoding/json"
	"net/http"
	"sql"
	"github.com/gorilla/mux"
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

// CreateClient .
func (c *Controller) CreateClient(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	var client Client
	_ = json.NewDecoder(r.Body).Decode(&client)
	//people = append(people, person)
	//insert client into DB
	sqlStatement := `
	INSERT INTO client (name, id, contact, stage, description)
	VALUES ($1, nextval('unique_id_sequence'), $3, $4, $5);
	`
	_, err := c.Db.Exec(sqlStatement, client.Name, client.Contact, client.Stage, client.Description)
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(200)
	}
}

// GetClient .
func (c *Controller) GetClient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sqlStatement := `
	SELECT name, id, contact, stage, description from client
	WHERE ID = $1;
	`

	row := c.Db.QueryRow(sqlStatement, params["id"])
	err := row.Scan(&col)
	if err != nil {
		if err == sql.ErrNoRows {
			w.Write([]byte("Zero rows found"))
			w.WriteHeader(404)
		} else {
			w.WriteHeader(500)
		}
	}else{
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(row)
	}
}

// DeleteClient .
func (c *Controller) DeleteClient(w http.ResponseWriter, r *http.Request) {
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