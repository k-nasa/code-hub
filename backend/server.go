package server

import (
	"fmt"

	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/justinas/alice"

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
	log.Printf("Listening on port %s", addr)
	err := http.ListenAndServe(
		fmt.Sprintf(":%s", addr),
		handlers.LoggingHandler(os.Stdout, s.router),
	)
	if err != nil {
		panic(err)
	}
}

func (s *Server) Route() *mux.Router {
	authMiddleware := middleware.NewAuthMiddleware(s.authClient, s.dbx)
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Authorization"},
	})

	commonChain := alice.New(
		middleware.RecoverMiddleware,
		corsMiddleware.Handler,
	)

	authChain := commonChain.Append(
		authMiddleware.Handler,
	)

	r := mux.NewRouter()

	r.Methods(http.MethodGet).Path("/public").Handler(commonChain.Then(handler.NewPublicHandler()))
	r.Methods(http.MethodGet).Path("/private").Handler(authChain.Then(handler.NewPrivateHandler(s.dbx)))

	articleController := controller.NewArticle(s.dbx)
	r.Methods(http.MethodPost).Path("/articles").Handler(authChain.Then(AppHandler{articleController.Create}))
	r.Methods(http.MethodPut).Path("/articles/{id}").Handler(authChain.Then(AppHandler{articleController.Update}))
	r.Methods(http.MethodDelete).Path("/articles/{id}").Handler(authChain.Then(AppHandler{articleController.Destroy}))
	r.Methods(http.MethodGet).Path("/articles").Handler(commonChain.Then(AppHandler{articleController.Index}))
	r.Methods(http.MethodGet).Path("/articles/{id}").Handler(commonChain.Then(AppHandler{articleController.Show}))

	fileServeRouter := r.PathPrefix("/").Subrouter()
	fileServeRouter.Use(commonChain.Then)
	fileServeRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/dist/index.html")
	})
	fileServeRouter.PathPrefix("/").Handler(
		http.StripPrefix("/", http.FileServer(http.Dir("../frontend/dist"))))

	return r
}
