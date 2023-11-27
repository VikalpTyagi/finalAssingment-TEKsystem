package main

import (
	"context"
	"finalAssing/internal/auth"
	"finalAssing/internal/config"
	"finalAssing/internal/database"
	"finalAssing/internal/handlers"
	"finalAssing/internal/repository"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog/log"
)

func main() {
	err := startApp()
	if err != nil {
		log.Panic().Err(err).Send()
	}
	log.Info().Msg("hello this is our app")
}
func startApp() error {
	config.Init()
	cfg := config.GetConfig() //@ this will give us config and initialize it
	log.Info().Msg("Config intialize sucessfully")
	// Initialize authentication support
	log.Info().Msg("main : Started : Initializing authentication support")
	// privatePEM, err := os.ReadFile(`private.pem`)
	// if err != nil {
	// 	return fmt.Errorf("reading auth private key %w", err)
	// }
	privatePEM := []byte(cfg.AuthKeys.PrivateKey)
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privatePEM)
	if err != nil {
		return fmt.Errorf("parsing auth private key %w", err)
	}

	// publicPEM, err := os.ReadFile(`pubkey.pem`)
	// if err != nil {
	// 	return fmt.Errorf("reading auth public key %w", err)
	// }
	publicPEM := []byte(cfg.AuthKeys.PublicKey)

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicPEM)
	if err != nil {
		return fmt.Errorf("parsing auth public key %w", err)
	}

	a, err := auth.NewAuth(privateKey, publicKey)
	if err != nil {
		return fmt.Errorf("constructing auth %w", err)
	}

	// Start Database
	log.Info().Msg("main : Started : Initializing db support")
	db, err := database.Open(cfg)
	if err != nil {
		return fmt.Errorf("connecting to db %w", err)
	}
	pg, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w ", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = pg.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("database is not connected: %w ", err)
	}

	//Initialize Conn layer support
	ms, err := database.NewConn(db)
	if err != nil {
		return err
	}
	err = database.AutoMigrate(ms)
	if err != nil {
		return err
	}
	// Intializing Repository layer
	repoStruct, err := repository.NewRepo(db)
	if err != nil {
		return err
	}
	// Intializing redis connection
	redis := database.NewRedis(cfg)
	_, err = redis.Ping(ctx).Result()
	if err != nil {
		log.Panic().Err(err).Msg("Connection with Redis not establish")
		return err
	}

	// Initialize http service
	api := http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.AppConfig.Host, cfg.AppConfig.Port),
		ReadTimeout:  time.Duration(cfg.AppConfig.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.AppConfig.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.AppConfig.IdleTimeout) * time.Second,
		Handler:      handlers.API(a, repoStruct, redis),
	}

	serverErrors := make(chan error, 1)
	go func() {
		log.Info().Str("port", api.Addr).Msg("main: API listening")
		serverErrors <- api.ListenAndServe()
	}()
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error %w", err)
	case sig := <-shutdown:
		log.Info().Msgf("main: Start shutdown %s", sig)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err := api.Shutdown(ctx)
		if err != nil {
			log.Error().Err(err).Msg("Server not working")
			err = api.Close()
			return fmt.Errorf("could not stop server gracefully %w", err)
		}

	}
	return nil

}
