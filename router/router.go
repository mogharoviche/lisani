package router

import (
	"lisani/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/lisaani/students", controller.GetAllstudents).Methods("GET")
	// router.HandleFunc("/lisaani/students/id", GetByID).Methods("GET")
	router.HandleFunc("/lisaani/students/edit/{id}", controller.Markedaspresent).Methods("PUT")
	router.HandleFunc("/lisaani/students/add", controller.Creatstudent).Methods("POST")
	router.HandleFunc("/lisaani/students/deleteAll", controller.DeleteAllStudents).Methods("DELETE")
	router.HandleFunc("/lisaani/students/delete/{id}", controller.DeleteStudent).Methods("DELETE")

	return router
}
