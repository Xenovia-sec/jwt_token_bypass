package controllers

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type Jwt struct {
	Jwt string `json:"jwt"`
}

func getSecret() string {
	return "sfgiants25"
}

func Index(c *fiber.Ctx) error {
	return c.Render("index", nil)
}

func GetJWT(c *fiber.Ctx) error {
	//jwt actions
	payload := jwt.StandardClaims{
		Subject:   "admin",
		ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, payload).SignedString([]byte(getSecret()))
	if err != nil {
		fmt.Println(err)
	}

	return c.Render("index", fiber.Map{
		"Jwt": token,
		"Msg": "success",
	})
}

func CheckJwt(c *fiber.Ctx) error {
	register := new(Jwt)

	if err := c.BodyParser(register); err != nil {
		log.Fatalln("error = ", err)
		return c.SendStatus(404)
	}
	//jwt actions
	token, err := jwt.ParseWithClaims(register.Jwt, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(getSecret()), nil
	})
	if err != nil || !token.Valid {
		return c.Render("index", fiber.Map{
			"Jwt": "",
			"Msg": "Your JWT is not valid.",
		})
	}
	return c.Render("index", fiber.Map{
		"Jwt": "",
		"Msg": "Your JWT is valid.",
	})
}
