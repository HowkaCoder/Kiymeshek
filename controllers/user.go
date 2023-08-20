package controllers

import (
	"go-auth/config"
	"go-auth/database"
	"go-auth/models"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetUser(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "invalid user ID"})
	}

	var user models.User
	if err := database.DB.Preload("Ava").First(&user, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "user not found"})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "server error"})
	}

	return c.JSON(user)
}
func UpdateUser(c *fiber.Ctx) error {
  userID , err := strconv.Atoi(c.Params("id"))
  if err != nil {
    return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message":"invalid user id"})
  }
  var user models.User
  if err := database.DB.First(&user , userID).Error; err != nil {
    return c.Status(http.StatusNotFound).JSON(fiber.Map{"message":"User not found"})
  }
  var updatedUser models.User 
  if err := c.BodyParser(&updatedUser); err != nil {
    return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message":"invalid request body"})
  }
  user.Name = updatedUser.Name
  user.Email = updatedUser.Email
  user.Password = updatedUser.Password
  if err := database.DB.Save(&user).Error; err != nil {
    return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message":"server error"})
  }

  return c.JSON(fiber.Map{"message":"user updated"})
}
func DeleteUser(c *fiber.Ctx) error {
    userID  , err := strconv.Atoi(c.Params("id"))
    if err != nil {
      return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message":"invalid request body"})
    }
    var user models.User

    if err := database.DB.First(&user, userID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
    }

    if err := database.DB.Delete(&user).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
    }

    return c.JSON(fiber.Map{"message": "user deleted"})
}

func Register(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		log.Fatal(err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	user.Password = string(hashedPassword)
	r := database.DB.Create(&user)
	if r.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "server error"})
	}

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		log.Fatal(err)
	}

	var foundUser models.User
	r := database.DB.Where("name = ?", user.Name).First(&foundUser)
	if r.Error != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "not found"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password)); err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "invalid password"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.JWTClaims{
		ID: foundUser.ID,
    Role:foundUser.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	})

	tokenString, err := token.SignedString([]byte(config.Config("SECRET_KEY")))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "error password",
		})
	}

	return c.JSON(fiber.Map{
		"token": tokenString,
	})
}

/* func Auth(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "error with Authorization header",
		})
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := jwt.ParseWithClaims(tokenString, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config("SECRET_KEY")), nil
	})
	if err != nil || !token.Valid {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid token",
		})
	}

	claims, ok := token.Claims.(*models.JWTClaims)
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid token",
		})
	}

	c.Locals("ID", claims.ID)
	return c.Next()
}

// ... (другие функции и импорты)



*/ 
func Auth(role string) fiber.Handler {
    return func(c *fiber.Ctx) error {
        authHeader := c.Get("Authorization")

        if authHeader == "" {
            return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
                "message": "Missing Authorization header",
            })
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        token, err := jwt.ParseWithClaims(tokenString, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
            return []byte(config.Config("SECRET_KEY")), nil
        })

        if err != nil || !token.Valid {
            return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
                "message": "Invalid token",
            })
        }

        claims, ok := token.Claims.(*models.JWTClaims)
        if !ok {
            return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
                "message": "Invalid token",
            })
        }

        c.Locals("ID", claims.ID)
        c.Locals("Role", claims.Role)

        // Проверка роли
        if claims.Role != role {
            return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
                "message": "Access denied",
                "claims":claims,
            })
        }

        return c.Next()
    }
}

