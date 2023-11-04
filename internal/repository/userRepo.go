package repository

import (
	"context"
	"finalAssing/internal/models"
	"fmt"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (r *ReposStruct) SaveUser(ctx context.Context, nu models.NewUser) (models.User, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Err(err).Msg("Error in hashing of Password")
		return models.User{}, fmt.Errorf("generating password hash: %w", err)
	}
	userData := models.User{
		Name:     nu.Name,
		Email:    nu.Email,
		PassHash: string(hashedPass),
	}
	err = r.db.Create(&userData).Error
	if err != nil {
		return models.User{}, err
	}
	return userData, nil
}

func (r *ReposStruct) CheckEmail(email, password string) (models.User, error) {
	var u models.User
	tx := r.db.Where("email = ?", email).First(&u)
	if tx.Error != nil {
		return models.User{}, tx.Error
	}
	return u, nil
}
