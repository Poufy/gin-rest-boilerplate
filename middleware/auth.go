package middleware

import (
	"gin-boilerplate/config"
	"gin-boilerplate/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		token, err := validateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Extract the user ID and other claims from the token if needed
		userID := claims["id"].(float64)
		email := claims["email"].(string)
		expirationTime := claims["exp"].(float64)

		if time.Now().Unix() > int64(expirationTime) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has expired"})
			c.Abort()
			return
		}

		// Set the claims in the context for future use
		c.Set("userID", userID)
		c.Set("email", email)

		c.Next()
	}
}

func GenerateAuthToken(user *models.User) (string, error) {
	// Set the expiration time for the token
	expirationTime := time.Now().Add(15 * time.Minute) // Token expires in 15 minutes

	// Create a new claim with user information and additional claims
	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"role":  "standard",            // Set the role claim to "standard"
		"exp":   expirationTime.Unix(), // Set the token expiration time
	}

	// Create a new token with the claims and signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	signedToken, err := token.SignedString([]byte(config.Cfg.JWT_SECRET.Key))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func validateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Provide the secret key used for signing the token
		return []byte(config.Cfg.JWT_SECRET.Key), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
