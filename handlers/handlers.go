package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gustavovargas511/sotho/middlew"
	"github.com/gustavovargas511/sotho/routers"
	"github.com/rs/cors"
)

//TheHandlers setteo mi puerto, el handler y pongo a escuchar mi servidor
func TheHandlers() {
	router := mux.NewRouter()

	router.HandleFunc("/signup", middlew.DBCheck(routers.Record)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
