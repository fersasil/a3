package routes

import (
	authcontroller "backend_ca/app/controllers/authController"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

/*
	@Funcion:
	@Arguments:
	@Returns:
	@Description:
*/
func AddSignHandler(r *mux.Router) {
	r.HandleFunc("/signup", authcontroller.SignupHandler).Methods("POST")
	r.HandleFunc("/signin", authcontroller.SigninHandler).Methods("GET")
	// r.HandleFunc("/signin", signinGetHandler).Methods("GET")
	// r.HandleFunc("/signin", signinPostHandler).Methods("POST")
	// r.HandleFunc("/signout", signoutGetHandler).Methods("GET")
}

/*
	@Funcion:
	@Arguments:
	@Returns:
	@Description:
*/
func signupGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}
