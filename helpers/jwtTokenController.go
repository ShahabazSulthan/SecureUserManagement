package helpers

import (
	models "admin/Models"
	"fmt"
	"log"
	"time"
	"os"

	"github.com/golang-jwt/jwt" 
	//jwt package for working with JSON Web Tokens (JWTs).
)


var expiryTime = time.Now().Local().Add(5 * time.Minute).Unix()
var SECRET_KEY = []byte(os.Getenv("KEY"))

func GenerateTokens(username, usertype string) string {

	//Creates a new instance of models.
	// SignedDetails struct, which likely contains user details and JWT claims.
	claims := &models.SignedDetails{
		Username:  username,
		User_type: usertype,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiryTime,
		},
	}

	//Creates a new JWT token using the provided claims and signs it with the HS256 signing method and the SECRET_KEY.
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("KEY")))
	if err != nil {
		log.Println(err)
	}
	//erro.CheckError(err)

	return token
}

func ValidateTokens(signedtoken string) bool {

	//Parses the signed token with the provided claims and returns a token object and an error
	token, err := jwt.ParseWithClaims(
		signedtoken,
		&models.SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("KEY")), nil
		})

	if err != nil {
		fmt.Println(err)
	}

	//Attempts to extract the claims from the token and asserts them to be of type models.SignedDetails.
	claims, ok := token.Claims.(*models.SignedDetails)

	//Checks if the token is not valid or expired and returns false.

	if !ok && claims.ExpiresAt < time.Now().Local().Unix() && !token.Valid {
		return false
	}

	return true

}
