package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, r)
	w.Write([]byte{'q'})

}

func RunServer() {
	r := mux.NewRouter()

	r.PathPrefix("/static").Handler(http.StripPrefix("/static",
		http.FileServer(http.Dir("./"+"static/"))))

	r.PathPrefix("/").Handler(http.StripPrefix("/",
		http.FileServer(http.Dir("./"+"static/"))))
	//r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/auth", JWT)
	r.HandleFunc("/welcome", Welcome)

	// http.Handle("/", r)

	srv := &http.Server{

		Handler:      r,
		Addr:         "127.0.0.1:3333",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
