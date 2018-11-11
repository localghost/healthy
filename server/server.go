package server

import (
	"github.com/gorilla/mux"
	"github.com/localghost/healthy/checker"
	"github.com/localghost/healthy/utils"
	"github.com/spf13/viper"
	"log"
	"net/http"
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
	v1.HandleFunc("/check/{name}", s.healthCheck)

	s.server = &http.Server{
		Handler: s.router,
		Addr: viper.Get("server.listen_on").(string),
	}
}

func (s *Server) Start() error {
	log.Println("Listening on:", s.server.Addr)
	return s.server.ListenAndServe()
}

func (s *Server) healthCheck(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	if checkError := s.checker.Check(vars["name"]); checkError != nil {
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
