package server

import (
	"fmt"
	"net/http"
	"time"

	_ "net/http/pprof"

	"github.com/ashwinath/simple/framework"
	"github.com/gorilla/mux"
)

type Server struct {
	fw           framework.FW
	router       *mux.Router
	port         int
	readTimeout  time.Duration
	writeTimeout time.Duration
}

func NewServer(fw framework.FW, options ...func(*Server)) *Server {
	s := &Server{
		fw:           fw,
		router:       mux.NewRouter(),
		port:         6000,
		readTimeout:  5 * time.Second,
		writeTimeout: 5 * time.Second,
	}

	for _, o := range options {
		o(s)
	}

	s.RegisterRoute("/ping", http.MethodGet, Ping)
	s.RegisterHandler("/debug/pprof", http.DefaultServeMux)

	return s
}

func WithPort(port int) func(*Server) {
	return func(s *Server) {
		s.port = port
	}
}

func WithReadTimeout(timeout time.Duration) func(*Server) {
	return func(s *Server) {
		s.readTimeout = timeout
	}
}

func WithWriteTimeout(timeout time.Duration) func(*Server) {
	return func(s *Server) {
		s.writeTimeout = timeout
	}
}

func (m *Server) RegisterRoute(path, method string, handler CustomHTTPHandler) *Server {
	m.router.HandleFunc(path, convertHandler(m.fw, handler)).Methods(method)
	return m
}

func (m *Server) RegisterHandler(path string, handler http.Handler) *Server {
	m.router.PathPrefix(path).Handler(handler)
	return m
}

func (m *Server) Serve() {
	srv := &http.Server{
		Handler:      m.router,
		Addr:         fmt.Sprintf("0.0.0.0:%d", m.port),
		WriteTimeout: m.writeTimeout,
		ReadTimeout:  m.readTimeout,
	}

	m.fw.GetLogger().Infof("running server on port: %d", m.port)
	m.fw.GetLogger().Fatal(srv.ListenAndServe())
}
