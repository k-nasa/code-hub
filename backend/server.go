package server

import (
	"fmt"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"

	"firebase.google.com/go/auth"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/rs/cors"
	"github.com/voyagegroup/treasure-app/controller"
	db2 "github.com/voyagegroup/treasure-app/db"
	"github.com/voyagegroup/treasure-app/handler"
	"github.com/voyagegroup/treasure-app/middleware"
	"github.com/voyagegroup/treasure-app/util"
)

type Server struct {
	dbx        *sqlx.DB
	router     *mux.Router
	authClient *auth.Client
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init(datasource string) {
	authClient, err := util.InitAuthClient()
	if err != nil {
		log.Fatalf("failed init auth client. %s", err)
	}
	s.authClient = authClient

	db := db2.NewDb(datasource)
	dbx, err := db.Open()
	if err != nil {
		log.Fatalf("failed db init. %s", err)
	}
	s.dbx = dbx
	s.router = s.Route()
}

func (s *Server) Run(addr string) {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Authorization"},
	})
	h := c.Handler(s.router)
	log.Printf("Listening on port %s", addr)
	err := http.ListenAndServe(fmt.Sprintf(":%s", addr), handlers.LoggingHandler(os.Stdout, h))
	if err != nil {
		panic(err)
	}
}

func (s *Server) Route() *mux.Router {
	authMiddleware := middleware.NewAuthMiddleware(s.authClient, s.dbx)
	r := mux.NewRouter()

	publicHandler := r.PathPrefix("/public").Subrouter()
	publicHandler.Methods("GET").Path("").Handler(handler.NewPublicHandler())

	privateHandler := r.PathPrefix("/private").Subrouter()
	privateHandler.Use(authMiddleware.Handler())
	privateHandler.Methods("GET").Path("").Handler(handler.NewPrivateHandler(s.dbx))

	articleController := controller.NewArticle(s.dbx)
	articleHandler := r.PathPrefix("/articles").Subrouter()
	articleHandler.Use(authMiddleware.Handler())
	articleHandler.Methods("POST").Path("").Handler(AppHandler{articleController.Create})
	articleHandler.Methods("PUT").Path("/{id}").Handler(AppHandler{articleController.Update})
	articleHandler.Methods("DELETE").Path("/{id}").Handler(AppHandler{articleController.Destroy})
	articleHandler.Methods("GET").Path("").Handler(AppHandler{articleController.Index})
	articleHandler.Methods("GET").Path("/{id}").Handler(AppHandler{articleController.Show})

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/dist/index.html")
	})
	r.PathPrefix("/").Handler(
		http.StripPrefix("/", http.FileServer(http.Dir("../frontend/dist"))))

	return r
}
