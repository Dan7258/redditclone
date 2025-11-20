package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"os"
	"redditclone/internal/models"
	"strconv"
	"strings"
	"time"
)

var (
	NoAuthHeaderError       = errors.New("no authorization header")
	BadSignMethodError      = errors.New("bad sign method")
	NoSecretKeyError        = errors.New("no secret key")
	InvalidTokenClaimsError = errors.New("invalid token claims")
)

var secretKey []byte

type Claims struct {
	Username string `json:"username"`
	ID       uint   `json:"id"`
	jwt.RegisteredClaims
}

type JwtResponse struct {
	Token string `json:"token"`
}

func Init() error {
	secretKey = []byte(os.Getenv("SECRET_KEY"))
	if string(secretKey) == "" {
		return NoSecretKeyError
	}
	return nil
}

func ParseToken(r *http.Request) (*jwt.Token, error) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		return nil, NoAuthHeaderError
	}
	inToken := strings.TrimPrefix(auth, "Bearer ")
	return jwt.Parse(inToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, BadSignMethodError
		}
		return secretKey, nil
	})
}

func ParseClaims(token *jwt.Token) (Claims, error) {
	claims := new(Claims)
	mapClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return *claims, InvalidTokenClaimsError
	}
	user := mapClaims["user"].(map[string]interface{})
	claims.Username = user["username"].(string)
	userID := user["id"].(string)
	id, err := strconv.Atoi(userID)
	if err != nil {
		return *claims, InvalidTokenClaimsError
	}
	claims.ID = uint(id)
	return *claims, nil
}

func GenerateJWT(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user": map[string]string{
			"id":       strconv.Itoa(int(user.ID)),
			"username": user.Username,
		},
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}
	log.Println(user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
