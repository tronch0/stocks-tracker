package http

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"stocks_tracker/service/analytics"
	"stocks_tracker/service/dataproviders"
	"stocks_tracker/service/http/contract"
	"time"
)

const ADDRESS = "127.0.0.1:8000"

type StocksTrackerHttpServer struct {
	providers map[string]dataproviders.Provider
	stats *analytics.Analytics
	router *mux.Router
}

func NewHttpServer(providers map[string]dataproviders.Provider, stats *analytics.Analytics) *StocksTrackerHttpServer {
	res := &StocksTrackerHttpServer{
		router:    mux.NewRouter(),
		providers: providers,
		stats: stats,
	}

	res.setApiPrefix()
	res.registerRoutes()

	return res
}

func (s *StocksTrackerHttpServer) setApiPrefix() {
	s.router = s.router.PathPrefix("/api/v1/").Subrouter()
}

func (s *StocksTrackerHttpServer) registerRoutes() {
	s.router.HandleFunc("/stats", s.getStats).Methods("GET")
	s.router.HandleFunc("/{assetType}", s.timeTracker(s.getQuote)).Methods("GET").Queries("symbol", "{symbol}","date", "{date}")
}

func (s *StocksTrackerHttpServer) getQuote (w http.ResponseWriter, r *http.Request) {

	assetType, symbol, date, err := s.parseRequestParams(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errResp := &contract.ErrorHttpResponse{Error: err.Error()}
		json.NewEncoder(w).Encode(errResp)
		return
	}

	var res float64

	if len(date) > 0 {
		res, err = s.providers[assetType].GetQuoteByDate(symbol,date)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errResp := &contract.ErrorHttpResponse{Error: err.Error()}
			json.NewEncoder(w).Encode(errResp)
			return
		}
	} else {
		res, err = s.providers[assetType].GetQuote(symbol)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errResp := &contract.ErrorHttpResponse{Error: err.Error()}
			json.NewEncoder(w).Encode(errResp)
			return
		}
	}

	httpRes := &contract.GetQuoteResponse{
		Price: res,
		Symbol: symbol,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(httpRes)
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
		s.stats.AddRequest(dInt)
	}
}

func (s *StocksTrackerHttpServer) getStats(w http.ResponseWriter, r *http.Request) {
	resp := &contract.StatsResponse{}
	resp.TotalRequests, resp.AvgProcessingTimeNs = s.stats.GetStats()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}


func (s *StocksTrackerHttpServer) parseRequestParams(r *http.Request) (assetType, symbol, date string, err error) {
	param := mux.Vars(r)
	//word := r.FormValue("word")

	assetType = param["assetType"]
	symbol = param["symbol"]
	date = param["date"]

	return assetType, symbol, date, s.validateRequestParameter(assetType,symbol,date)
}
func (s *StocksTrackerHttpServer) validateRequestParameter(assetType, symbol, date string) error {
	if _, isExist := s.providers["crypto"]; isExist == false {
		return fmt.Errorf("request parameter \"assetType\" is invalid")
	}

	if len(symbol) == 0 {
		return fmt.Errorf("request parameter \"symbol\" is invalid")
	}

	if len(date) != 0  {
		_, err := time.Parse("YYYY-MM-DD", date)
		if err != nil {
			return fmt.Errorf("error: date is not a valid")
		}

		return fmt.Errorf("request parameter \"type/symbol\" is empty")
	}

	return nil
}
