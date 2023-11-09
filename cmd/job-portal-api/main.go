package main

import (
	"context"
	"finalAssing/internal/auth"
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

	// Initialize authentication support
	log.Info().Msg("main : Started : Initializing authentication support")
	privatePEM, err := os.ReadFile(`C:\Users\ORR Training 17\Documents\final assingment\job portal api\private.pem`)
	// privatePEM, err := os.ReadFile(`F:\go-TEKsystem\finalAssingment-TEKsystem\private.pem`)
	if err != nil {
		return fmt.Errorf("reading auth private key %w", err)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privatePEM)
	if err != nil {
		return fmt.Errorf("parsing auth private key %w", err)
	}

	publicPEM, err := os.ReadFile(`C:\Users\ORR Training 17\Documents\final assingment\job portal api\pubkey.pem`)
	// publicPEM, err := os.ReadFile(`F:\go-TEKsystem\finalAssingment-TEKsystem\pubkey.pem`)
	if err != nil {
		return fmt.Errorf("reading auth public key %w", err)
	}

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
	db, err := database.Open()
	if err != nil {
		return fmt.Errorf("connecting to db %w", err)
	}
	pg, err := db.DB()
	if err != nil {
		return fmt.Errorf("Failed to get database instance: %w ", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = pg.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("Database is not connected: %w ", err)
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

	// Initialize http service
	api := http.Server{
		Addr:         ":8080",
		ReadTimeout:  8000 * time.Second,
		WriteTimeout: 800 * time.Second,
		IdleTimeout:  800 * time.Second,
		Handler:      handlers.API(a, repoStruct),
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
			err = api.Close()
			return fmt.Errorf("could not stop server gracefully %w", err)
		}

	}
	return nil

}
