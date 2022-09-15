package facade

import (
	"database/sql"
	"log"
	"net/http"
)

// Facade
type Service struct {
	server *HTTPServer
	db     *DB
}

func NewService() *Service {
	s := new(Service)
	s.server = new(HTTPServer)
	s.db = new(DB)

	return s
}

func (s *Service) StartService() {
	s.server.Start()
	s.db.Start()
}

func (s *Service) StopService() {
	s.server.Stop()
	s.db.Stop()
}

// Subsystem classes
type HTTPServer struct {
	server *http.Server
}

func (hs *HTTPServer) Start() {
	log.Fatalln(hs.server.ListenAndServe())
}

func (hs *HTTPServer) Stop() {
	log.Fatalln(hs.server.Close())
}

type DB struct {
	conn *sql.DB
}

func (db *DB) Start() {
	conn, err := sql.Open("test", "test")

	if err != nil {
		log.Fatalln(err)
	}

	db.conn = conn
}

func (db *DB) Stop() {
	log.Fatalln(db.conn.Close())
}
