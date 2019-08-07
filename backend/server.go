package server

import (
	"fmt"
	"log"
	"net/http"

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
	fmt.Printf("Listening on port %s", addr)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", addr), h); err != nil {
		panic(err)
	}
}

func (s *Server) Route() *mux.Router {
	authMiddleware := middleware.NewAuthMiddleware(s.authClient, s.dbx)
	r := mux.NewRouter()
	r.Handle("/public", handler.NewPublicHandler())
	r.Handle("/private", authMiddleware.Handler(handler.NewPrivateHandler(s.dbx)))

	articleController := controller.NewArticle(s.dbx)
	articleAuthRequired := r.PathPrefix("/articles").Subrouter()
	articleAuthRequired.Use(authMiddleware.Handler)
	articleAuthRequired.Handle("", AppHandler{articleController.New}).Methods("POST")
	articleAuthRequired.Handle("/{id}", AppHandler{articleController.Edit}).Methods("PUT")

	articleNonAuth := r.PathPrefix("/articles").Subrouter()
	articleNonAuth.Handle("", AppHandler{articleController.Root}).Methods("GET")
	articleNonAuth.Handle("/{id}", AppHandler{articleController.Get}).Methods("GET")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/dist/index.html")
	})
	r.PathPrefix("/").Handler(
		http.StripPrefix("/", http.FileServer(http.Dir("../frontend/dist"))))

	return r
}
