package server

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/geraldofigueiredo/mypage/config"
	"github.com/geraldofigueiredo/mypage/handler"
	"github.com/geraldofigueiredo/mypage/service"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() error {
	// load api configs
	config := config.LoadConfig()
	// services
	services := service.NewServices(config)

	ctx := context.Background()
	ctx = context.WithValue(ctx, "config", config)
	ctx = context.WithValue(ctx, "services", services)

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.Handle("/", &handler.Template{Filename: "index.html"})

	services.Log.Infof("server running on port %s", config.AppPort)
	return http.ListenAndServe(config.AppPort, mux)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("template", "index.html")

	tmpl, err := template.ParseFiles(lp)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	}
}
