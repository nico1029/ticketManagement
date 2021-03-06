package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// The new router function creates the router and
// returns it to us. We can now use this function
// to instantiate and test the router outside of the main function
func newRouter() *mux.Router {
	r := mux.NewRouter()

	// Declare the static file directory and point it to the directory we just made
	staticFileDirectory := http.Dir("./assets/")
	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/assets/" prefix when looking for files.
	// For example, if we type "/assets/index.html" in our browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for "./assets/assets/index.html", and yield an error
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/assets/", instead of the absolute route itself
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/ticket", getTicketsHandler).Methods("GET")
	r.HandleFunc("/ticket", createTicketHandler).Methods("POST")
	r.HandleFunc("/ticket/delete", deleteTicketHandler).Methods("POST")
	r.HandleFunc("/ticket/update", updateTicketHandler).Methods("POST")
	r.HandleFunc("/ticket/get/deleted", getDeletedTicketsHandler).Methods("GET")
	r.HandleFunc("/ticket/recover", recoverTicketHandler).Methods("POST")
	r.HandleFunc("/ticket/recoverall", recoverAllTicketsHandler).Methods("POST")
	r.HandleFunc("/ticket/filter", filterTicketsHandler).Methods("POST")
	r.HandleFunc("/ticket/filter/delete", deleteFilterTicketsHandler).Methods("POST")
	return r
}

func main() {
	fmt.Println("Starting server...")
	connString := "host=localhost port=5432 password=$YOUR_PASSWORD$ user=postgres dbname=ticket_database sslmode=disable"
	db, err := sql.Open("postgres", connString)

	if err != nil {
		panic(err)
	}
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	InitStore(&dbStore{db: db})

	// The router is now formed by calling the `newRouter` constructor function
	// that we defined above. The rest of the code stays the same
	r := newRouter()
	fmt.Println("Serving on port 8082")
	http.ListenAndServe(":8082", r)
}
