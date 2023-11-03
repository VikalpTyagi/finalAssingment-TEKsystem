// * User service
package services

import (
	"context"
	"finalAssing/internal/models"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser is a method that creates a new user record in the database.
func (s *Store) CreateUser(ctx context.Context, nu models.NewUser) (models.User, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, fmt.Errorf("generating password hash: %w", err)
	}
	userData,err:= s.Repo.SaveUser(nu,hashedPass)
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

	// Successful authentication! Generate JWT claims.
	c := jwt.RegisteredClaims{
		Issuer:    "service project",
		Subject:   strconv.FormatUint(uint64(userData.ID), 10),
		Audience:  jwt.ClaimStrings{"students"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	// And return those claims.
	return c, nil
}
