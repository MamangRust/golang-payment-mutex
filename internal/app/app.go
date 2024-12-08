package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"payment-mutex/internal/handler"
	recordmapper "payment-mutex/internal/mapper/record"
	"payment-mutex/internal/repository"
	"payment-mutex/internal/service"
	"payment-mutex/pkg/auth"
	"payment-mutex/pkg/dotenv"
	"payment-mutex/pkg/hash"
	"payment-mutex/pkg/logger"
	"runtime"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Run() {
	log, err := logger.NewLogger()

	if err != nil {
		log.Fatal("Error creating logger: ", zap.Error(err))
	}

	if runtime.NumCPU() > 2 {
		runtime.GOMAXPROCS(runtime.NumCPU() / 2)
	}

	err = dotenv.Viper()

	if err != nil {
		log.Fatal("Error loading .env file: ", zap.Error(err))
	}

	hashing := hash.NewHashingPassword()

	repository := repository.NewRepositorys(repository.Deps{
		MapperRecord: *recordmapper.NewRecordMapper(),
	})

	token, err := auth.NewManager(viper.GetString("JWT_SECRET"))

	if err != nil {
		log.Fatal("Error creating manager: ", zap.Error(err))

	}

	service := service.NewServices(service.Deps{
		Repository: repository,
		Logger:     *log,
		Hash:       hashing,
		Token:      token,
	})

	myhandler := handler.NewHandler(service)

	serve := &http.Server{
		Addr:         fmt.Sprintf(":%s", viper.GetString("PORT")),
		WriteTimeout: time.Duration(viper.GetInt("WRITE_TIME_OUT")) * time.Second * 10,
		ReadTimeout:  time.Duration(viper.GetInt("READ_TIME_OUT")) * time.Second * 10,

		IdleTimeout: time.Second * 60,
		Handler:     myhandler.Init(),
	}

	go func() {
		if err := serve.ListenAndServe(); err != nil {
			log.Fatal("Error starting server: ", zap.Error(err))
		}
	}()

	log.Info("Connected to port: " + viper.GetString("PORT"))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	serve.Shutdown(ctx)
	os.Exit(0)
}
