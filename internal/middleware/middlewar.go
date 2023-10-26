package middleware

import (
	"errors"
	"finalAssing/internal/auth"
)

type Mid struct {
	auth *auth.Auth
}

func NewMid(a *auth.Auth ) (Mid,error){
if a == nil {
return Mid{}, errors.New("auth struct not provided")
}
return Mid{auth: a},nil
}
