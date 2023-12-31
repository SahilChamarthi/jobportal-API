package main

import (
	"fmt"
	"net/http"
	"os"
	"project/internal/auth"
	"project/internal/database"
	"project/internal/handlers"
	redispack "project/internal/redisPack"
	"project/internal/repository"
	"project/internal/services"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func main() {
	err := startApp()
	if err != nil {
		log.Panic().Err(err).Send()
	}
}
func startApp() error {
	log.Info().Msg("started main")
	privatePEM, err := os.ReadFile(`C:\Users\ORR Training 7\Documents\job-portal-apii\private.pem`)
	if err != nil {
		return fmt.Errorf("cannot find file private.pem %w", err)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privatePEM)
	if err != nil {
		return fmt.Errorf("cannot convert byte to key %w", err)
	}

	publicPEM, err := os.ReadFile(`C:\Users\ORR Training 7\Documents\job-portal-apii\pubkey.pem`)
	if err != nil {
		return fmt.Errorf("cannot find file pubkey.pem %w", err)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicPEM)
	if err != nil {
		return fmt.Errorf("cannot convert byte to key %w", err)
	}
	a, err := auth.NewAuth(privateKey, publicKey)
	if err != nil {
		return fmt.Errorf("cannot create auth instance %w", err)
	}

	redisClient := redispack.NewRedisClient()
	rc := redispack.NewRedisConnection(redisClient)

	db, err := database.DataBaseConnect()
	if err != nil {
		return err
	}
	repo, err := repository.NewRepo(db)
	if err != nil {
		return err
	}

	se, err := services.NewServices(repo, rc)

	if err != nil {
		return err
	}

	api := http.Server{ //server configuration
		Addr:    ":8082",
		Handler: handlers.Api(a, se),
	}
	api.ListenAndServe()

	return nil

}
