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
	readTimeout  int
	writeTimeout int
}

func NewServer(fw framework.FW, port, readTimeout, writeTimeout int) *Server {
	s := &Server{
		fw:           fw,
		router:       mux.NewRouter(),
		port:         port,
		readTimeout:  readTimeout,
		writeTimeout: writeTimeout,
	}

	s.RegisterRoute("/ping", http.MethodGet, Ping)
	s.RegisterHandler("/debug/pprof", http.DefaultServeMux)

	return s
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
		WriteTimeout: time.Duration(m.writeTimeout) * time.Second,
		ReadTimeout:  time.Duration(m.readTimeout) * time.Second,
	}

	m.fw.GetLogger().Infof("running server on port: %d", m.port)
	m.fw.GetLogger().Fatal(srv.ListenAndServe())
}
