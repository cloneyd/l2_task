package pattern

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

import (
	"database/sql"
	"log"
	"net/http"
)

// Facade
type ServiceFacade interface {
	Start()
	Stop()
}

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

func (s *Service) Start() {
	s.server.Start()
	s.db.Start()
}

func (s *Service) Stop() {
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
