package router

import (
	"lisani/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.Use(enableCORS)
	router.HandleFunc("/lisaani/students", controller.GetAllstudents).Methods("GET", "OPTIONS")
	// router.HandleFunc("/lisaani/students/id", GetByID).Methods("GET", "OPTIONS")
	router.HandleFunc("/lisaani/students/edit/{id}", controller.Markedaspresent).Methods("PUT", "OPTIONS")
	router.HandleFunc("/lisaani/students/add", controller.Creatstudent).Methods("POST", "OPTIONS")
	router.HandleFunc("/lisaani/students/deleteAll", controller.DeleteAllStudents).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/lisaani/students/delete/{id}", controller.DeleteStudent).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/register", controller.RegisterHandler).Methods("POST")
	router.HandleFunc("/login", controller.LoginHandler).Methods("GET")
	return router
}
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Log the request if needed
		log.Printf("[%s] %s %s", r.Method, r.URL.Path, r.RemoteAddr)

		next.ServeHTTP(w, r)
	})
}
