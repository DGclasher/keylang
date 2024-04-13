package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/DGclasher/keylang/service/license"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db: db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api").Subrouter()
	licenseService := license.NewLicense(s.db)
	licenseHandler := license.NewHandler(licenseService)
	licenseHandler.LicenseRoutes(subrouter)

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}