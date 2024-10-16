package http

import (
	"blixenkrone/spirii/internal/chargers"
	"blixenkrone/spirii/storage"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// type chargersQuerier interface {
// 	Getchargers(ctx context.Context, id uuid.UUID) (example.chargers, error)
// 	Writechargers(ctx context.Context, arg example.WritechargersParams) (example.chargers, error)
// }

type Server struct {
	logger logrus.FieldLogger
	srv    *http.Server
	memDB  storage.TableProjectionReadWriter[chargers.MeterReading]
	// fq     chargersQuerier
}

func NewServer(l logrus.FieldLogger, addr string, db storage.TableProjectionReadWriter[chargers.MeterReading]) Server {
	r := mux.NewRouter()
	srv := http.Server{
		Addr:              addr,
		Handler:           r,
		ReadTimeout:       time.Second * 20,
		ReadHeaderTimeout: 0,
		WriteTimeout:      time.Second * 20,
		IdleTimeout:       time.Second * 20,
		MaxHeaderBytes:    1 << 20,
	}

	s := Server{l, &srv, db}
	s.registerRoutes(r)

	return s
}

// TODO: improve loggerMW with better logging
func (s Server) loggerMW(h http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		s.logger.Infof("calling %s w method %s", r.URL, r.Method)
		h(rw, r)
	}
}

func (s Server) registerRoutes(fh *mux.Router) {
	routes := map[string]struct {
		fn     http.HandlerFunc
		method string
	}{
		"/ping":             {s.pong(), http.MethodGet},
		"/v1/chargers/{id}": {s.getchargersV1(), http.MethodGet},
		"/v1/chargers":      {s.postchargersV1(), http.MethodPost},
		// "/v2/chargers/{id}": {s.getchargersV2(), http.MethodGet},
		// "/v2/chargers":      {s.postchargersV2(), http.MethodPost},
	}
	for k, v := range routes {
		v.fn = s.loggerMW(v.fn)
		fh.HandleFunc(k, v.fn).Methods(v.method)
	}
}

func (s Server) ListenAndServe() error {
	return s.srv.ListenAndServe()
}

func (s Server) ShutDown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func (s Server) pong() http.HandlerFunc {
	return func(rw http.ResponseWriter, _ *http.Request) {
		rw.Write([]byte("PONG"))
	}
}

func (s Server) getAuth() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

	}
}

func (s Server) getchargersV1() http.HandlerFunc {
	type response struct {
		charger chargers.MeterReading
	}
	return func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		chargersID, ok := params["id"]
		if !ok {
			http.Error(rw, "id not provided", http.StatusBadRequest)
			return
		}

		rec, err := s.memDB.Read(r.Context(), chargersID)
		if err != nil {
			if errors.Is(err, chargers.ErrRecordNotFound) {
				http.Error(rw, err.Error(), http.StatusNotFound)
				return
			}
			http.Error(rw, "id not found", http.StatusNotFound)
			return
		}

		resp := response{
			charger: rec,
		}

		if err := json.NewEncoder(rw).Encode(&resp); err != nil {
			panic(err)
		}
	}
}

func (s Server) postchargersV1() http.HandlerFunc {
	type request struct {
		ID    string `json:"id"`
		Value string `json:"value"`
	}
	validateBodyFn := func(r request) error {
		if r.ID == "" || r.Value == "" {
			return errors.New("bad body values")
		}
		return nil
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var body request
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := validateBodyFn(body); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := s.memDB.Write(r.Context(), chargers.MeterReading{
			// ID:    body.ID,
			// Value: body.Value,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusAccepted)

	}
}
