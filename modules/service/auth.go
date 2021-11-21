package service

import (
	"errors"
	"time"
	"tubespbbo/hashing"
	"tubespbbo/modules/dto"
	"tubespbbo/modules/repository"

	"github.com/dgrijalva/jwt-go"
)

var (
	ACCESS_SECRET = "test-secret"
)

func Login(loginRequestDTO *dto.LoginDto) (string, error) {
	username := loginRequestDTO.Username
	password := loginRequestDTO.Password

	user, err := repository.FindUserByUsername(username)
	if !hashing.ComparePasswords(user.Password, []byte(password)) || err != nil || user == nil {
		return "", errors.New("invalid password")
	}

	return createToken(user.Id)
}

func createToken(userid int64) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(ACCESS_SECRET))
	if err != nil {
		return "", err
	}

	return token, nil
}
