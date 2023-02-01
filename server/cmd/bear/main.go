package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PakSerVal/bear-masha/internal/application/http/auth"
	"github.com/PakSerVal/bear-masha/internal/application/http/auth/middleware"
	"github.com/PakSerVal/bear-masha/internal/application/http/cards"
	"github.com/PakSerVal/bear-masha/internal/application/http/file"
	"github.com/PakSerVal/bear-masha/internal/config"
	"github.com/PakSerVal/bear-masha/internal/domain/auth/usecase/sign_in"
	"github.com/PakSerVal/bear-masha/internal/domain/auth/usecase/validate"
	"github.com/PakSerVal/bear-masha/internal/domain/cards/usecase/create"
	"github.com/PakSerVal/bear-masha/internal/domain/cards/usecase/delete"
	"github.com/PakSerVal/bear-masha/internal/domain/cards/usecase/get_by_id"
	"github.com/PakSerVal/bear-masha/internal/domain/cards/usecase/get_list"
	"github.com/PakSerVal/bear-masha/internal/domain/cards/usecase/update"
	"github.com/PakSerVal/bear-masha/internal/domain/file/usecase/upload"
	"github.com/PakSerVal/bear-masha/internal/infrastucture/logger"
	"github.com/PakSerVal/bear-masha/internal/infrastucture/repository"
	"github.com/PakSerVal/bear-masha/pkg/psql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.uber.org/zap"
)

func main() {
	var err error

	defer func() {
		if panicErr := recover(); panicErr != nil {
			log.Fatal(panicErr)
		}

		if err != nil {
			log.Fatal(err)
		}
	}()

	c := config.New()

	logger.InitLogger(c.IsDev)

	db, err := psql.Connect(c.DbConn)
	if err != nil {
		logger.Fatal("connect to db failed", zap.Error(err))
	}

	cardRepo := repository.NewCardRepository(db)
	fileRepo := repository.NewFileRepository(db)
	cardFilesRepo := repository.NewCardFileRepository(db)

	getListUc := get_list.New(cardRepo, fileRepo)
	getByIDUc := get_by_id.New(cardRepo, fileRepo)
	createUc := create.New(cardRepo, cardFilesRepo)
	deleteUc := delete.New(cardRepo)
	updateUc := update.New(cardRepo)
	signInUC := sign_in.New(c.Login, c.Password, c.SecretKey)
	validateUC := validate.New(c.SecretKey)
	uploadFileUC := upload.New(fileRepo)

	cardHandler := cards.New(getListUc, getByIDUc, createUc, deleteUc, updateUc)
	authHandler := auth.New(signInUC)

	fileHandler := file.New(uploadFileUC)

	r := mux.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

	r.HandleFunc("/auth", authHandler.Auth).Methods("POST", "OPTIONS")

	r.HandleFunc("/cards", cardHandler.GetList).Methods("GET", "OPTIONS")
	r.HandleFunc("/cards", cardHandler.Create).Methods("POST", "OPTIONS")
	r.HandleFunc("/cards/{id:[0-9]+}", cardHandler.GetByID).Methods("GET", "OPTIONS")
	r.HandleFunc("/cards/{id:[0-9]+}", cardHandler.Delete).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/cards/{id:[0-9]+}", cardHandler.Update).Methods("PUT", "OPTIONS")

	r.HandleFunc("/file/upload", fileHandler.Upload)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("static/"))))

	authMiddleware := middleware.New(validateUC)
	r.Use(authMiddleware.Middleware)

	addr := fmt.Sprintf(":%d", c.HTTPPort)
	logger.Info("starting server", zap.String("addr", addr))
	_ = http.ListenAndServe(addr, cors.Handler(r))
}
