package main

import (
	"github.com/gorilla/mux"
	"database/sql"
	_ "github.com/lib/pq"
	"net/http"
	"strconv"
	"encoding/json"
	"fmt"
	"log"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type LogReturn struct {
	Name string `json:"name"`
	Email string `json:"email"`
	adminAccess bool `json:"access"`
	Stamp	time.Time `json:"time"`
}

type App struct {
	Router *mux.Router
	DB	   *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.MarshalIndent(payload, "", "    ")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) getIssue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	i := issue{ID: id}
	if err := i.getIssue(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Product not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, i)
}


func (a *App) updateIssue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var i issue
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&i); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	i.ID = id

	if err := i.updateIssue(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, i)
}

func (a *App) createIssue(w http.ResponseWriter, r *http.Request) {
	var i issue
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&i); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := i.createIssue(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, i)
}

func (a *App) deleteIssue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	i := issue{ID: id}
	if err := i.deleteIssue(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (a *App) getIssues(w http.ResponseWriter, r *http.Request) {
	issues, err := getIssues(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, issues)
}

func (a *App) isUserEnabled(w http.ResponseWriter, req *http.Request) {

	email := req.FormValue("email")

	var enabled bool

	err := a.DB.QueryRow(`SELECT hasEnabledAccount FROM users WHERE email=$1`, email).Scan(&enabled)

	// Handle if the query fails
	if err != nil {
		w.Write([]byte("Invalid User"))
		return
	}

	if enabled {
		w.Write([]byte("Enabled"))
	} else {
		w.Write([]byte("Not Enabled"))
	}

}

func (a *App) performLogin(w http.ResponseWriter, req *http.Request){

	// User input credentials
	email := req.FormValue("email")

	password := req.FormValue("password")

	//fmt.Sprintf("user=%s password=%s", email, password)

	var dbName string
	var dbEmail	string
	var dbPass string
	var dbAdmin	bool

	err := a.DB.QueryRow(`SELECT name, email, password, isadmin FROM users WHERE email=$1`, email).Scan(&dbName, &dbEmail, &dbPass, &dbAdmin)

	// Handle if the query fails
	if err != nil {
		w.Write([]byte("Failed Query Execution within Login"))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(password))

	if err != nil {
		w.Write([]byte("Invalid Password"))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	returnResult := LogReturn{dbName, dbEmail, dbAdmin, time.Now() }


	json.NewEncoder(w).Encode(returnResult)

}

func (a *App) performRegister(w http.ResponseWriter, req *http.Request){

	// User input credentials
	email := req.FormValue("email")
	name := req.FormValue("name")
	password := req.FormValue("password")

	stmt, err := a.DB.Prepare(`UPDATE users SET name=$1, password=$2, hasEnabledAccount=true, isadmin=false WHERE email=$3`)

	if err != nil {
		log.Fatal(err)
		w.Write([]byte("Failed to Register"))
		return
	}

	_, err = stmt.Exec(name, password, email)

	if err != nil {
		w.Write([]byte("Failed to register"))
		return
	}

	w.Write([]byte("Registered"))

}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/issue/{id:[0-9]+}", a.getIssue).Methods("GET")
	a.Router.HandleFunc("/issue", a.createIssue).Methods("POST")
	a.Router.HandleFunc("/issue/{id:[0-9]+}", a.updateIssue).Methods("PUT")
	a.Router.HandleFunc("/issue/{id:[0-9]+}", a.deleteIssue).Methods("DELETE")
	a.Router.HandleFunc("/issues", a.getIssues).Methods("GET")

	a.Router.HandleFunc("/getUser", a.isUserEnabled).Methods("POST")

	//These will be GET for testing purposes... Should be POST
	a.Router.HandleFunc("/login", a.performLogin).Methods("POST")
	a.Router.HandleFunc("/register", a.performRegister).Methods("POST")
}