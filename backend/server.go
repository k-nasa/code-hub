package server

import (
	"fmt"

	"log"
	"net/http"
	"os"

	"firebase.google.com/go/auth"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/justinas/alice"
	"github.com/k-nasa/code-hub/controller"
	db2 "github.com/k-nasa/code-hub/db"
	"github.com/k-nasa/code-hub/firebase"
	"github.com/k-nasa/code-hub/middleware"
	"github.com/k-nasa/code-hub/sample"
	"github.com/rs/cors"
)

type Server struct {
	db         *sqlx.DB
	router     *mux.Router
	authClient *auth.Client
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init(datasource string) {
	authClient, err := firebase.InitAuthClient()
	if err != nil {
		log.Fatalf("failed init auth client. %s", err)
	}
	s.authClient = authClient

	db := db2.NewDB(datasource)
	dbcon, err := db.Open()
	if err != nil {
		log.Fatalf("failed db init. %s", err)
	}
	s.db = dbcon
	s.router = s.Route()
}

func (s *Server) Run(addr string) {
	log.Printf("Listening on port %s", addr)
	err := http.ListenAndServe(
		fmt.Sprintf(":%s", addr),
		handlers.CombinedLoggingHandler(os.Stdout, s.router),
	)
	if err != nil {
		panic(err)
	}
}

func (s *Server) Route() *mux.Router {
	authMiddleware := middleware.NewAuth(s.authClient, s.db)
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
	r.Methods(http.MethodGet).Path("/public").Handler(commonChain.Then(sample.NewPublicHandler()))
	r.Methods(http.MethodGet).Path("/private").Handler(authChain.Then(sample.NewPrivateHandler(s.db)))

	codeController := controller.NewCode(s.db)
	r.Methods(http.MethodPost).Path("/codes").Handler(authChain.Then(AppHandler{codeController.Create}))
	r.Methods(http.MethodGet).Path("/codes").Handler(commonChain.Then(AppHandler{codeController.Index}))
	r.Methods(http.MethodGet).Path("/codes/{id}").Handler(commonChain.Then(AppHandler{codeController.Show}))

	// FIXME urlとIndexWithUserっていうのが微妙、、、
	r.Methods(http.MethodGet).Path("/users/codes").Handler(commonChain.Then(AppHandler{codeController.IndexWithUser}))
	r.Methods(http.MethodGet).Path("/users/{id}/codes").Handler(commonChain.Then(AppHandler{codeController.ShowUserCode}))

	r.PathPrefix("").Handler(commonChain.Then(http.StripPrefix("/img", http.FileServer(http.Dir("./img")))))
	return r
}
