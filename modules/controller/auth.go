package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	e "tubespbbo/err"
	"tubespbbo/modules/dto"
	"tubespbbo/modules/service"
	"tubespbbo/response"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var (
	ACCESS_SECRET = "test-secret"
)

type AccessDetails struct {
	AccessUuid string
	UserId     uint64
}

func Login(c *fiber.Ctx) error {
	loginDto := new(dto.LoginDto)
	err := c.BodyParser(loginDto)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var token string
	token, err = service.Login(loginDto)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: token,
	})
	return nil
}

func ExtractToken(c *fiber.Ctx) string {
	bearerToken := string(c.Request().Header.Peek("Authorization"))
	strArr := strings.Split(bearerToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(ACCESS_SECRET), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValid(c *fiber.Ctx) error {
	token, err := VerifyToken(c)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func ExtractTokenMetadata(c *fiber.Ctx) (*AccessDetails, error) {
	token, err := VerifyToken(c)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &AccessDetails{
			AccessUuid: accessUuid,
			UserId:     userId,
		}, nil
	}
	return nil, err
}

func AuthCheck(c *fiber.Ctx) {
	_, authErr := ExtractTokenMetadata(c)
	if authErr != nil {
		e.HandleErr(c, authErr)
	}
}
