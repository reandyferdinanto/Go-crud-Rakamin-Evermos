package handler

import (
	"log"
	"time"
	"tryFiberGo/database"
	"tryFiberGo/model/entity"
	"tryFiberGo/model/request"
	"tryFiberGo/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func LoginHandler(c *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
	log.Println(loginRequest)}

	// VALIDASI REQUEST
	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if  errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"Message" : "failed",
			"error" : errValidate.Error(),
		}) 
	}


// FIND USER BY EMAIl
var user entity.User
err := database.DB.First(&user, "email = ?", loginRequest.Email).Error
if err != nil {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"Message": "Wrong Credential",
	})
}

// CHECK VALIDATION PASSWORD
isValid := utils.CheckPasswordHash(loginRequest.Password, user.Password)
if !isValid{
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"Message" : "Wrong credential",
	})
}

//GENERATE JWT
claims := jwt.MapClaims{}
claims["name"] = user.Name
claims["email"] = user.Email
// claims["address"] = user.Address
claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

if user.Email == "aqua@gmail.com" {
	claims["role"] = "admin"
}else {
	claims ["role"] = "user"
}


token, errGenerateToken := utils.GenerateToken(&claims)
if errGenerateToken != nil {
	log.Println(errGenerateToken)
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "wrong credential",
	})
	
}


return c.JSON(fiber.Map{
	"token": token,
})
}