package auth

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func CreateToken(id string, exp int, secret string) (string, error) {
	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Minute * time.Duration(exp)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err

}

func AccessToken(id string) string {
	godotenv.Load()
	expInt, _ := strconv.Atoi(os.Getenv("EXP_ACCESS_TOKEN"))
	token, _ := CreateToken(id, expInt, os.Getenv("ACCESS_TOKEN_SECRET"))
	return token
}

func RefreshToken(id string) string {
	godotenv.Load()
	expInt, _ := strconv.Atoi(os.Getenv("EXP_REFRESH_TOKEN"))
	token, _ := CreateToken(id, expInt, os.Getenv("ACCESS_TOKEN_SECRET"))
	return token
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type ReqRefToken struct {
	RefreshToken string `json:"refresh_token"`
}

func Restricted(c echo.Context) string {
	post := c.Get("post").(*jwt.Token)
	claims := post.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	return id
}

func Middleware() echo.MiddlewareFunc {
	godotenv.Load()
	config := echojwt.Config{
		SigningKey: []byte(os.Getenv("ACCESS_TOKEN_SECRET")),
	}
	return echojwt.WithConfig(config)
}

func Parse(refreshToken, secretkey string) (interface{}, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("error :")
		}

		return []byte(secretkey), nil
	})

	if err != nil {
		return nil, err
	}

	t := token.Claims.(jwt.MapClaims)
	id := t["id"].(string)
	return id, nil
}
