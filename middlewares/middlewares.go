package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seefnasrul/go-gin-gorm/auth"
)

// func SetMiddlewareJSON() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		w.Header().Set("Content-Type", "application/json")
// 		c.Next()
// 	}
// }

func SetMiddlewareAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}


// func TokenAuthMiddleware() gin.HandlerFunc {
// 	requiredToken := os.Getenv("API_TOKEN")
  
// 	// We want to make sure the token is set, bail if not
// 	if requiredToken == "" {
// 	  log.Fatal("Please set API_TOKEN environment variable")
// 	}
  
// 	return func(c *gin.Context) {
// 	  token := c.Request.FormValue("api_token")
  
// 		auth := c.Request.Header.Get("Authorization")
// 		if auth == "" {
// 			c.String(http.StatusForbidden, "No Authorization header provided")
// 			c.Abort()
// 			return
// 		}
// 		token := strings.TrimPrefix(auth, "Bearer ")
// 		if token == auth {
// 			c.String(http.StatusForbidden, "Could not find bearer token in Authorization header")
// 			c.Abort()
// 			return
// 		}
  
// 	  c.Next()
// 	}
// }


// func OAuthIntrospectionHandler(endpoint, iss, aud string, pk jose.JsonWebKey) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		auth := c.Request.Header.Get("Authorization")
// 		if auth == "" {
// 			c.String(http.StatusForbidden, "No Authorization header provided")
// 			c.Abort()
// 			return
// 		}
// 		token := strings.TrimPrefix(auth, "Bearer ")
// 		if token == auth {
// 			c.String(http.StatusForbidden, "Could not find bearer token in Authorization header")
// 			c.Abort()
// 			return
// 		}
// 		jwt := NewClientJWT(iss, aud)
// 		clientAssertion, err := SignJWT(jwt, pk)
// 		if err != nil {
// 			c.AbortWithError(http.StatusInternalServerError, err)
// 			return
// 		}
// 		ir, err := IntrospectToken(endpoint, token, iss, clientAssertion)
// 		if err != nil {
// 			c.AbortWithError(http.StatusInternalServerError, err)
// 			return
// 		}
// 		if !ir.Active {
// 			c.String(http.StatusForbidden, "Provided token is no longer active")
// 			c.Abort()
// 			return
// 		}
// 		c.Set("scopes", ir.SplitScope())
// 		c.Set("subject", ir.SUB)
// 		c.Set("clientID", ir.ClientID)
// 	}
// }