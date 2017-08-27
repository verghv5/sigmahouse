package main
//
//import (
//		"database/sql"
//		"net/http"
//		"time"
//		"golang.org/x/crypto/bcrypt"
//
//	"encoding/json"
//)
//
//var con *sql.DB
//var err error
//
//
//type LoginReturn struct {
//	Name	string 'json:"name"'
//	Email	string 'json:"email"'
//	adminAccess	bool 'json:"access"'
//	Stamp	time.Time
//}
//
//func (a *App) initializeRoutesForLogin() {
//
//
//    a.Router.HandleFunc("/getUser", a.getUser).Methods("GET")
//
//	//These will be GET for testing purposes... Should be POST
//	a.Router.HandleFunc("/login", a.login).Methods("GET")
//	a.Router.HandleFunc("/register", a.register).Methods("GET")
//
//}
//
//type Error struct {
//
//}
//
//
//func main(){
//
//}
//
//func getUser(w http.ResponseWriter, req *http.Request) {
//
//	username := req.FormValue("user")
//
//	var enabled bool
//
//	err := db.QueryRow("SELECT hasEnabledAccount FROM users WHERE username=?", username).Scan(&enabled)
//
//	// Handle if the query fails
//	if err != nil {
//		res.Write([]byte("Failed Query Execution within GET USER"))
//        return
//    }
//
//	if enabled {
//		res.Write([]byte("Found"))
//	} else {
//		res.Write([]byte("None"))
//	}
//
//}
//
//func login(w http.ResponseWriter, req *http.Request){
//
//	// User input credentials
//	email := req.FormValue("email")
//	password := req.FormValue("pass")
//
//	var dbName string
//	var dbEmail	string
//	var dbPass string
//	var dbAdmin	bool
//
//	err := db.QueryRow("SELECT Name, Email, Password, Admin FROM users WHERE username=?", username).Scan(&dbName, &dbEmail, &dbPass, &dbAdmin)
//
//	// Handle if the query fails
//	if err != nil {
//		res.Write([]byte("Failed Query Exection within Login"))
//        return
//    }
//
//	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
//
//	if err != nil {
//		res.Write([]byte("Invalid Password"))
//		return
//	}
//
//	w.Header.Set("Content-Type", "application/json")
//	w.WriteHeader(code)
//
//	returnResult := LoginReturns{
//		LoginReturn{dbName, dbEmail, dbAdmin},
//	}
//
//	json.NewEncoder(w).Encode(LoginReturns)
//
//}
//
//func register(w http.ResponseWriter, req *http.Request){
//
//	// User input credentials
//	email := req.FormValue("email")
//	password := req.FormValue("pass")
//
//	err := db.QueryRow("UPDATE users SET Name=?, Password=?, hasEnabledAccount=?, admin=?, session_id=? WHERE Email=?", email).Scan(&dbName, &dbEmail, &dbPass, &dbAdmin)
//
//
//	return nil
//}
//
///* May be implemented in a later version....
//
//func createAndReturnSessionId(){
//
//} */