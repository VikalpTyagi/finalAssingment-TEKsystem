package middleware

import (
	"errors"
	"finalAssing/internal/auth"

	"github.com/rs/zerolog/log"
)

type Mid struct {
	auth *auth.Auth
}

func NewMid(a *auth.Auth ) (Mid,error){
if a == nil {
	log.Error().Err(errors.New("auth instance is not provided")).Send()
return Mid{}, errors.New("auth struct not provided")
}
return Mid{auth: a},nil
}
