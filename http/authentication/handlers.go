package authentication

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gocondor/core/auth"
	"github.com/gocondor/core/database"
	"github.com/gocondor/core/jwt"
	"github.com/gocondor/gocondor/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginCreds struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	DB := database.Resolve()
	Auth := auth.Resolve()
	JWT := jwt.Resolve()

	// validate and bind user input
	var loginData LoginCreds
	if err := c.ShouldBind(&loginData); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}

	// get the user record by email from db
	var user models.User
	result := DB.Where("email = ?", loginData.Email).First(&user)
	// check if the record not found
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "wrong credentials",
		})
		return
	}

	// handle database error incase there is any
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "someting went wrong",
		})
		return
	}

	//compare the password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		// wrong password
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "wrong credentials",
		})
		return
	}

	// prepare jwt token payload
	tokenPayload := map[string]interface{}{
		"userId": user.ID,
	}
	// generate the jwt token
	token, err := JWT.CreateToken(tokenPayload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
		})
	}

	// prepare the refresh token payload
	refreshTokenPayload := map[string]interface{}{
		"userId": user.ID,
	}
	// generate the token
	refreshToken, err := JWT.CreateRefreshToken(refreshTokenPayload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
		})
	}

	// login the user
	Auth.Login(user.ID, c)

	// render response
	c.JSON(http.StatusOK, gin.H{
		"data": map[string]string{
			"token":        token,
			"refreshToken": refreshToken,
		},
	})
}

func Logout(c *gin.Context) {
	Auth := auth.Resolve()
	JWT := jwt.Resolve()

	// extract the jwt token
	token, err := JWT.ExtractToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
		})
		return
	}
	// decode the jwt token
	payload, err := JWT.DecodeToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
		})
		return
	}
	// get the user id of the token and convert it to uint64
	userId, err := strconv.ParseUint(fmt.Sprintf("%v", payload["userId"]), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
		})
		return
	}

	fmt.Println(Auth.Check(userId, c))

	// log the user out only if he is authenticated
	if Auth.Check(userId, c) {
		// log the user out
		err := Auth.Logout(userId, c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "logged out successfuly",
	})
}

func Register(c *gin.Context) {
	DB := database.Resolve()
	// bind the input to the user's model
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}

	// check if there is a record with the given email
	res := DB.Where("email = ?", user.Email).First(&models.User{})
	if res.Error == nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "user already signed up",
		})
		return
	}

	//hash the passowrd
	hahsedPWD, err := hashPassword(user.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	//use the hashed password
	user.Password = hahsedPWD
	// create the db record
	res = DB.Create(&user)
	if res.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": res.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "signup successfully",
	})
}

// hashPassword hashs passwords
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
