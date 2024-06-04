package main

import (
	"fmt"
	"net/http"

	"github.com/dsypasit/social-clone/server/config"
	"github.com/dsypasit/social-clone/server/internal/share/db"
	"github.com/dsypasit/social-clone/server/internal/user"
	"github.com/dsypasit/social-clone/server/pkg/logger"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func main() {
	cfg := config.New().All()
	db.Init(cfg.DBConnection)
	defer db.Close()

	usrRepo := user.NewUserRepository(db.DB)

	usrSrv := user.NewUserService(usrRepo)

	usrHandler := user.NewUserHandler(usrSrv)

	router := mux.NewRouter()
	router = router.PathPrefix("/api/v1").Subrouter()

	user.RegisterUserRouter(router, usrHandler)
	router.HandleFunc("/healtcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// slogger := logger.NewLogger()
	zlog, _ := zap.NewProduction()
	defer zlog.Sync()
	router.Use(logger.Middleware(zlog))

	fmt.Printf("Running server with port %d\n", 8000)
	http.ListenAndServe(":8000", router)
}
