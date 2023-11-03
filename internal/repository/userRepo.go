package repository

import (
	"finalAssing/internal/models"
)

func (r *ReposStruct) SaveUser(userData models.NewUser, hashedPass []byte) (models.User, error) {
	u := models.User{
		Name:     userData.Name,
		Email:    userData.Email,
		PassHash: string(hashedPass),
	}
	err := r.db.Create(&u).Error
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

func (r *ReposStruct) CheckEmail(email, password string) (models.User, error) {
	var u models.User
	tx := r.db.Where("email = ?", email).First(&u)
	if tx.Error != nil {
		return models.User{}, tx.Error
	}
	return u, nil
}
