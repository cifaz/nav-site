package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"time"
)

type JWT struct {
}

func (j *JWT) Make(name, secret string, expire time.Duration) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": name,
		"exp":  time.Now().Add(expire * time.Second).Unix(),
	})
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (j *JWT) Check(token, secret string) (string, error) {
	//secretBytes, err := base64.URLEncoding.DecodeString(secret)
	log.Info("token:{}, secret: {}", token, secret)
	secretBytes := []byte(secret)
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

		return secretBytes, nil
	})
	if err != nil {
		fmt.Println("jwt error:", err)
		return "", err
	}
	return claim.Claims.(jwt.MapClaims)["name"].(string), nil
}
