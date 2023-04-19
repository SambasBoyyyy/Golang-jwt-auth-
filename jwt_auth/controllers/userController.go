package controllers

import (
	"go/jwt_auth/intializers"
	"go/jwt_auth/models"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	//"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {

	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// hash password

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash",
		})
		return
	}

	user := models.User{Email: body.Email, Password: string(hash)}

	result := intializers.GetDB().Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	// respond

	c.JSON(http.StatusOK, gin.H{})

}

func Login(c *gin.Context) {

	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	//looking for user

	var user models.User

	intializers.GetDB().First(&user, "email=?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user or password",
		})
		return
	}

	// now we get the user and going to match the hashed password

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user or password 2",
		})
		return
	}

	// if  pass matched then genearte token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{

		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// sign in and get encoded token

	tokenstring, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization",tokenstring,3600 * 24 * 30,"","",false,true)


	c.JSON(http.StatusOK, gin.H{
         "ID" : user.ID,
		"token": tokenstring,
	})

}


func Validate(c *gin.Context){
  
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message" : user,
	})
}
