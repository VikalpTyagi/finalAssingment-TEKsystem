// * User service
package services

import (
	"context"
	"finalAssing/internal/models"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *Store) CreateUser(ctx context.Context, nu models.NewUser) (models.User, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Err(err).Msg("Error in hasing of Password")
		return models.User{}, fmt.Errorf("generating password hash: %w", err)
	}
	userData,err:= s.Repo.SaveUser(nu,hashedPass)
	if err!= nil{
		return models.User{},err
	}
	return userData, nil
}

func (s *Store) Authenticate(ctx context.Context, email, password string) (jwt.RegisteredClaims,
	error) {
	userData,err := s.Repo.CheckEmail(email,password)
	if err != nil {
		return jwt.RegisteredClaims{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userData.PassHash), []byte(password))
	if err != nil {
		return jwt.RegisteredClaims{}, err
	}

	c := jwt.RegisteredClaims{
		Issuer:    "service project",
		Subject:   strconv.FormatUint(uint64(userData.ID), 10),
		Audience:  jwt.ClaimStrings{"users"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	return c, nil
}
