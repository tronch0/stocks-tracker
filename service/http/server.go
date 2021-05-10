package http

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
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
	s.router.HandleFunc("/similar", s.timeTracker(s.getSimilarWords)).Methods("GET").Queries("word", "{word}")}
}

func (s *StocksTrackerHttpServer) Start() error {
	srv := &http.Server{
		Handler:      s.router,
		Addr:         ADDRESS,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}

func (s *StocksTrackerHttpServer) timeTracker(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		h.ServeHTTP(w, r)
		d := time.Now().Sub(startTime).Nanoseconds()

		dInt := int(d)
		s.analytics.AddRequest(dInt)
	}
}

func (s *StocksTrackerHttpServer) getStats(w http.ResponseWriter, r *http.Request) {
	resp := &contract.StatsHttpResponse{}
	resp.TotalWords, resp.TotalRequests, resp.AvgProcessingTimeNs = s.analytics.GetStats()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}


func validateRequestParameter(w string) error {
	if len(w) == 0 {
		return fmt.Errorf("request parameter \"type/symbol\" is empty")
	}

	return nil
}
