package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Student Student `json:"student"`
	Role    string  `json:"role"`
	jwt.StandardClaims
}
