package server

import (
	"firebase.google.com/go/auth"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/rs/cors"
	db2 "github.com/voyagegroup/treasure-app/db"
	"github.com/voyagegroup/treasure-app/handler"
	"github.com/voyagegroup/treasure-app/middleware"
	"github.com/voyagegroup/treasure-app/util"
	"net/http"
)

type Server struct {
	dbx    *sqlx.DB
	router *mux.Router
}

func NewServer(datasource string) (*Server, error) {
	authClient, err := util.InitAuthClient()
	if err != nil {
		return nil, err
	}

	db := db2.NewDb(datasource)
	dbx, err := db.Open()
	if err != nil {
		return nil, err
	}

	return &Server{
		router: Route(authClient, dbx),
	}, nil
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

func Route(client *auth.Client, dbx *sqlx.DB) *mux.Router {

	authMiddleware := middleware.NewAuthMiddleware(client)
	r := mux.NewRouter()
	r.Handle("/public", handler.NewPublicHandler())
	r.Handle("/private", authMiddleware.Handler(handler.NewPrivateHandler(dbx)))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/dist/index.html")
	})
	r.PathPrefix("/").Handler(
		http.StripPrefix("/", http.FileServer(http.Dir("../frontend/dist"))))

	return r
}
