package token

import (
	"fmt"
	"time"

	"github.com/dostonshernazarov/e-commerce-go/internal/pkg/logger"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
)

type JwtHandler struct {
	Sub       string
	Iss       string
	Exp       string
	Iat       string
	Aud       []string
	Role      string
	Token     string
	SigninKey string
	Log       *zap.Logger
	Timeout   int
}

func (jwtHandler *JwtHandler) GenerateJwt() (access, refresh string, err error) {
	var (
		accessToken, refreshToken *jwt.Token
		claims                    jwt.MapClaims
	)

	accessToken = jwt.New(jwt.SigningMethodHS256)
	refreshToken = jwt.New(jwt.SigningMethodHS256)

	claims = accessToken.Claims.(jwt.MapClaims)
	claims["sub"] = jwtHandler.Sub
	claims["iss"] = jwtHandler.Iss
	claims["exp"] = time.Now().Add(time.Hour * 200).Unix()
	claims["iat"] = time.Now().Unix()
	claims["role"] = jwtHandler.Role

	access, err = accessToken.SignedString([]byte(jwtHandler.SigninKey))
	if err != nil {
		jwtHandler.Log.Error("error generating access token", logger.Error(err))
		// logger.Error(err)
		return
	}

	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = jwtHandler.Sub
	rtClaims["exp"] = time.Now().Add(time.Hour * 400).Unix()
	rtClaims["iat"] = time.Now().Unix()
	rtClaims["role"] = jwtHandler.Role

	refresh, err = refreshToken.SignedString([]byte(jwtHandler.SigninKey))
	if err != nil {
		// jwtHandler.Log.Error("error generating refresh token", logger.Error(err))
		logger.Error(err)
		return
	}

	return access, refresh, nil
}

func ExtractClaim(tokenStr string, signingKey []byte) (jwt.MapClaims, error) {
	var (
		token *jwt.Token
		err   error
	)

	token, err = jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		err = fmt.Errorf("invalid JWT Token")
		return nil, err
	}

	return claims, nil
}
