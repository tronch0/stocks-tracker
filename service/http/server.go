package http

import (
	"github.com/gorilla/mux"
)

const ADDRESS = "127.0.0.1:8000"

type StocksTrackerHttpServer struct {
	router *mux.Router
}

func NewHttpServer() *StocksTrackerHttpServer {
	res := &StocksTrackerHttpServer{
		router:    mux.NewRouter(),
	}

	// set prefix
	res.router = res.router.PathPrefix("/api/v1/").Subrouter()
	res.setApiPrefix()
	res.registerRoutes()

	return res
}

func (s *StocksTrackerHttpServer) setApiPrefix() {
	s.router = s.router.PathPrefix("/api/v1/").Subrouter()
}

func (s *StocksTrackerHttpServer) registerRoutes() {
	//s.router.HandleFunc("/stats", s.getStats).Methods("GET")
	//s.router.HandleFunc("/similar", s.timeTracker(s.getSimilarWords)).Methods("GET").Queries("word", "{word}")}

}