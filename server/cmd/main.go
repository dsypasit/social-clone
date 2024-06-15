package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dsypasit/social-clone/server/config"
	"github.com/dsypasit/social-clone/server/internal/auth"
	"github.com/dsypasit/social-clone/server/internal/post"
	"github.com/dsypasit/social-clone/server/internal/share/db"
	"github.com/dsypasit/social-clone/server/internal/user"
	"github.com/dsypasit/social-clone/server/pkg"
	"github.com/dsypasit/social-clone/server/pkg/logger"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	cfg := config.New().All()
	log.Println(cfg.DBConnection)
	if err := db.Init(cfg.DBConnection); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	usrRepo := user.NewUserRepository(db.DB)
	postRepo := post.NewPostRepository(db.DB)

	usrSrv := user.NewUserService(usrRepo)
	jwtSrv := auth.NewJwtService("test")
	authSrv := auth.NewAuthService(usrSrv, jwtSrv)
	postSrv := post.NewPostService(postRepo, usrSrv)

	usrHandler := user.NewUserHandler(usrSrv)
	authHandler := auth.NewAuthHandler(authSrv)
	postHandler := post.NewPostHandler(postSrv)

	router := mux.NewRouter()
	router = router.PathPrefix("/api/v1").Subrouter()

	user.RegisterUserRouter(router, usrHandler)
	auth.RegisterAuthRouter(router, authHandler)
	post.RegisterPostRouter(router, postHandler)

	router.HandleFunc("/healtcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello world"))
	}).Methods(http.MethodGet)

	// slogger := logger.NewLogger()
	zlogConfig := zap.NewProductionConfig()
	zlogConfig.EncoderConfig.TimeKey = "timestamp"
	zlogConfig.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	zlog, _ := zlogConfig.Build()
	defer zlog.Sync()
	logMiddleware := logger.Middleware(zlog)

	fmt.Printf("Running server with port %d\n", cfg.Server.Port)
	http.ListenAndServe(fmt.Sprintf(":%v", cfg.Server.Port), logMiddleware(pkg.CorsMiddleware(router)))
}
