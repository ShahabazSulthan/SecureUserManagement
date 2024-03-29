package middleware

import (
	"admin/helpers"

	"github.com/gin-gonic/gin"
)

//middleware in the Gin web framework.

func AuthMiddleWare() gin.HandlerFunc {
	//This function represents the actual middleware logic.
	return func(c *gin.Context) {
		var stat bool //Declares a boolean variable stat to store the validation status of the token.
		cookie, err := c.Cookie("user") /*Retrieves the value of the "user" 
		                                cookie from the request context and stores
										it in the cookie variable.*/

		if err != nil {
			c.Set("stat", "false")
		}

		stat = helpers.ValidateTokens(cookie)
		c.Set("stat", stat) /*Sets the value of the "stat" key in the request 
		                    context to the value of the stat variable,
							 indicating whether the token is valid.*/



		c.Next() /*Calls the Next method of the gin.
		           Context object to pass control
				   to the next middleware or route handler in the chain.*/
	}
}