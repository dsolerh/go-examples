package handlers

import (
	"gin_service/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const (
	ExpirationStartDuration    = 10 * time.Minute
	ExpirationIncresedDuration = 5 * time.Minute
	ExpirationRefreshPeriod    = 30 * time.Second
)

type AuthHandler struct {
	useAPIKey bool
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type JWTOutput struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}

func (handler AuthHandler) SignInHandler(c *gin.Context) {
	if handler.useAPIKey {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot signin, request an API key"})
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Username != "admin" || user.Password != "password" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	expirationTime := time.Now().Add(ExpirationStartDuration)
	claims := &Claims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(os.Getenv("JWT_SECRET"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	jwtOutput := JWTOutput{
		Token:   tokenStr,
		Expires: expirationTime,
	}
	c.JSON(http.StatusOK, jwtOutput)
}

func (handler AuthHandler) RefreshHandler(c *gin.Context) {
	if handler.useAPIKey {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot refresh token, request an API key"})
		return
	}

	tokenValue := c.GetHeader("Authorization")
	var claims Claims

	tkn, err := jwt.ParseWithClaims(
		tokenValue,
		&claims,
		func(_ *jwt.Token) (interface{}, error) { return []byte(os.Getenv("JWT_SECRET")), nil },
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if tkn == nil || !tkn.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	if claims.ExpiresAt.Sub(time.Now()) > ExpirationRefreshPeriod {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is not expired yet"})
		return
	}

	expirationTime := time.Now().Add(ExpirationIncresedDuration)
	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(os.Getenv("JWT_SECRET"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	jwtOutput := JWTOutput{
		Token:   tokenString,
		Expires: expirationTime,
	}
	c.JSON(http.StatusOK, jwtOutput)
}

func (handler AuthHandler) AuthMiddleware() gin.HandlerFunc {
	if handler.useAPIKey {
		return func(c *gin.Context) {
			if c.GetHeader("X-API-KEY") != os.Getenv("X_API_KEY") {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
			c.Next()
		}
	}

	// jwt verification
	return func(c *gin.Context) {
		tokenValue := c.GetHeader("Authorization")
		var claims Claims

		tkn, err := jwt.ParseWithClaims(
			tokenValue,
			&claims,
			func(_ *jwt.Token) (interface{}, error) { return []byte(os.Getenv("JWT_SECRET")), nil },
		)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		if tkn == nil || !tkn.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Next()
	}
}
