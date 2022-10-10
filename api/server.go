package api

import (
	"DATN/configs"
	"DATN/infrastructure"
	"log"
	"net/http"
)

type Server struct {
	cfg *configs.Server
}

func New(cfg *configs.Server) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Start() error {

	db, err := infrastructure.NewDBHandle(s.cfg.Database, "localhost")
	if err != nil {
		log.Println(err)
		return err
	}
	err = db.Exec()
	if err != nil {
		log.Println(err)
	}
	//err = http.ListenAndServe(model.PortConn, nil)
	//if err != nil {
	//	log.Println(err)
	//}
	//http.HandleFunc("/", HandleRequest)

	return nil
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {

	switch r.RequestURI {
	case "/":

	case "/home":

	}

}
