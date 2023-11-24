package repository

import (
	"context"
	"errors"
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
		DOB:      nu.DOB,
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
		log.Error().Err(tx.Error).Msg("Email not found in database")
		return models.User{}, tx.Error
	}
	return u, nil
}

func (r *ReposStruct) CheckEmailDob(data *models.ForgetPass) error {
	userData, err := r.CheckEmail(data.Email, "not necessary")
	if err != nil {
		log.Error().Err(err).Msg("Email verification failed in checkEmailDob")
		return err
	}
	if userData.DOB != data.DOB {
		log.Error().Msg("Dob of passed doesn't match with what is data base")
		return errors.New("DOB is Invalid")
	}
	return nil
}

func (r *ReposStruct) UpdatePassword(userEmail string, password string) error {
	var u models.User
	tx := r.db.Where("email = ?", userEmail).First(&u)
	if tx.Error != nil {
		log.Error().Err(tx.Error).Msg("Email not found in database")
		return tx.Error
	}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Err(err).Msg("Error in hashing of Password")
		return fmt.Errorf("generating password hash: %w", err)
	}
	u.PassHash = string(hashedPass)
	updateTx := r.db.Model(&u).Updates(models.User{PassHash: u.PassHash})
	if updateTx.Error != nil {
		log.Error().Err(updateTx.Error).Msg("password Updation failed in repo")
		return err
	}
	log.Info().Interface("User email",userEmail).Msg("Password Updated")
	return nil
}
