package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/localghost/healthy/utils"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"github.com/localghost/healthy/checker"
)

type Server struct {
	router *mux.Router
	server *http.Server
	checker *checker.Checker
}

func New(checker *checker.Checker) *Server {
	result := &Server{
		checker: checker,
	}
	result.setup()
	return result
}

func (s *Server) setup() {
	s.router = mux.NewRouter()

	v1 := s.router.PathPrefix("/v1/").Subrouter()
	v1.HandleFunc("/healthy/check/{healthcheck}", s.healthCheck)

	s.server = &http.Server{
		Handler: s.router,
		Addr: fmt.Sprintf("%s:%d", viper.Get("server.address"), viper.Get("server.port")),
	}
}

func (s *Server) Start() error {
	log.Println("Listening on:", s.server.Addr)
	return s.server.ListenAndServe()
}

func (s *Server) healthCheck(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	if checkError := s.checker.Check(vars["healthcheck"]); checkError != nil {
		switch checkError.(type) {
		case utils.NoSuchCheckError:
			response.WriteHeader(http.StatusNotFound)
		default:
			response.WriteHeader(http.StatusExpectationFailed)
		}
		response.Write([]byte(checkError.Error()))
	} else {
		response.WriteHeader(http.StatusOK)
	}
}
