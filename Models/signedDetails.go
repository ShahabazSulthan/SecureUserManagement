package models

import "github.com/golang-jwt/jwt"

type SignedDetails struct {
	Username  string
	User_type string
	jwt.StandardClaims 
	//is a struct that represents the standard claims defined by the JSON Web Token (JWT)
	//  specification, such as issuer, subject,
	//  audience, expiration time, not before, and issued at.
}