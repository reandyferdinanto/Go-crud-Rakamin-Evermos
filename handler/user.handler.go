package handler

import (
	"log"
	"tryFiberGo/database"
	"tryFiberGo/model/entity"
	"tryFiberGo/model/request"
	"tryFiberGo/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserHandlerGetAll(c *fiber.Ctx) error {
	userInfo := c.Locals("userInfo")
	log.Println("user info data :: ", userInfo)

	var users[]entity.User
	result:= database.DB.Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return c.JSON(users)
}

func UserHandlerCreate(c *fiber.Ctx) error {
	user := new(request.UserCreateRequest)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message" : "failed to store data",
			"error" : errValidate.Error(), 
		})
	}

	newUser := entity.User{
		Name: user.Name,
		No_telp: user.No_telp,
		Tanggal_lahir: user.Tanggal_lahir,
		Jenis_kelamin: user.Jenis_kelamin,
		Tentang: user.Tentang,
		Pekerjaan: user.Pekerjaan,
		Email: user.Email,
		Id_provinsi: user.Id_provinsi,
		Id_kota: user.Id_kota,
		IsAdmin: user.IsAdmin,
	}

	hashedPassword, err := utils.HashingPassword(user.Password)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"Message" : "Internal Server Error",
		})
	}

	newUser.Password = hashedPassword

	errCreateUser := database.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return c.Status(500).JSON(fiber.Map{
			"message" : "failed to store data",
		})
	}
	return c.JSON(fiber.Map{
		"message" : "data successfully stored",
	})
}

func UserHandlerGetById(c *fiber.Ctx) error {
	userId := c.Params ("id")

	var user entity.User

	err:= database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"Message" : "user not found",
		})
	}

	return c.JSON(fiber.Map{
		"Message" : "Success",
		"data" : user,
	})
}

func UserHandlerUpdate(c *fiber.Ctx) error {
	userRequest := new(request.UserUpdateRequest)

	if err := c.BodyParser(userRequest); err != nil {
	   return c.Status(400).JSON(fiber.Map{
		"Message" : "bad request",
	})
   }

	var user entity.User
	
	userId := c.Params ("id")
	
	err:= database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"Message" : "user not found",
		})
	}

	// UPDATE USER DATA
	if userRequest.Name != "" {
		user.Name = userRequest.Name
	}
	user.No_telp = userRequest.No_telp
	user.Tanggal_lahir = userRequest.Tanggal_lahir
	user.Jenis_kelamin = userRequest.Jenis_kelamin
	user.Tentang = userRequest.Tentang
	user.Pekerjaan = userRequest.Pekerjaan

	errUpdate := database.DB.Save(&user).Error
	if errUpdate != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message" : "Internal Server Error",
			"data" : user,
		})
		
	}
	return c.JSON(fiber.Map{
		"Message" : "Updated",
		"data" : user,
	})
}


func UserHandlerUpdateEmail(c *fiber.Ctx) error {
	userRequest := new(request.UserEmailRequest)
	if err := c.BodyParser(userRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"Message" : "bad request",
		})
	}
	var user entity.User
	var isEmailUserExist entity.User
	
	userId := c.Params("id")
	//CHECK AVAILABLE USER

	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"Message" : "user not found",
		})
	}

	errCheckEmail := database.DB.First(&isEmailUserExist, "email = ?",userRequest.Email).Error
	if errCheckEmail == nil {
		return c.Status(402).JSON(fiber.Map{
			"Message" : "email already used.",
		})
	}

		//UPDATE USER DATA - EMAIL
		user.Email = userRequest.Email
		errUpdate := database.DB.Save (&user).Error
		if errUpdate != nil {
			return c.Status(500).JSON(fiber.Map{
				"Message" : "Internal Server Error",
			})
		}
		return c.JSON(fiber.Map{
			"message" : "success",
			"data" : user,
		})
}

func UserHandlerDelete(c *fiber.Ctx) error {
	userId := c.Params("id")
	var user entity.User

	err := database.DB.Debug().First(&user, "id = ?", userId).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"Message": "user not found",
		})
	}

	errDelete := database.DB.Debug().Delete(&user).Error
	if errDelete != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": "Internal Server Error",
		})
	}

	response := fiber.Map{
		"Message": "User was deleted",
	}

	return c.JSON(response)
}
	