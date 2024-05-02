package controller

import (
	"net/http"
	"fmt"
	"log"
	"math/rand"
	"unicode"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/yermanovberik/golang-final-project/internal/controller/tokenutil"
	"github.com/yermanovberik/golang-final-project/internal/models"
	"github.com/yermanovberik/golang-final-project/pkg"
	"golang.org/x/crypto/bcrypt"
)

type SignupController struct {
	UserRepository models.UserRepository
	Env            *pkg.Env
}

var (
	verifier = emailverifier.NewVerifier()
)

func (sc *SignupController) Login(c *gin.Context) {
	var loginRequest models.LoginRequest

	err := c.ShouldBind(&loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_BIND_JSON",
					Message: "Datas dont match with struct of signin",
				},
			},
		})
		return
	}

	if loginRequest.Email == "" || loginRequest.Password == "" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "EMPTY_VALUES",
					Message: "Empty values are written in the form",
				},
			},
		})
		return
	}

	user, err := sc.UserRepository.GetUserByEmail(c, loginRequest.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_GET_USER",
					Message: "User with this email doesn't found",
				},
			},
		})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)) != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "PASSWORD_INCORRECT",
					Message: "Password doesn't match",
				},
			},
		})
		return
	}
	accessToken, err := tokenutil.CreateAccessToken(&user, sc.Env.AccessTokenSecret, sc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "TOKEN_ERROR",
					Message: "Error to create access token",
				},
			},
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{Result: accessToken})
}

func (uc *SignupController) Signup(c *gin.Context) {
	var request models.UserRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_BIND_JSON",
					Message: "Datas dont match with struct of signup",
				},
			},
		})
		return
	}

	verifier = verifier.EnableSMTPCheck()
	verifier = verifier.EnableDomainSuggest()

	if request.Email == "" || request.Password == "" || request.FullName == "" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "EMPTY_FIELDS",
					Message: "Not all fields provided",
				},
			},
		})
		return
	}

	ret, _ := verifier.Verify(request.Email)
	
	log.Println(ret.Reachable)
	
	if !ret.Syntax.Valid {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "Неверный адрес электронной почты",
					Message: "email address syntax is invalid",
				},
			},
		})
		return
	}

	user, _ := uc.UserRepository.GetUserByEmail(c, request.Email)
	if user.ID > 0 {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "USER_EXISTS",
					Message: "User with this email already exists",
				},
			},
		})
		return
	}
	err := validatePassword(request.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_PASSWORD_FORMAT",
					Message: err.Error(),
				},
			},
		})
		return
	}
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_ENCRYPTE_PASSWORD",
					Message: "Couldn't encrypte password",
				},
			},
		})
		return
	}
	request.Password = string(encryptedPassword)

	_, err = uc.UserRepository.CreateUser(c, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_CREATE_USERS",
					Message: "Couldn't create user",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
		return
	}
	user, err = uc.UserRepository.GetUserByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_GET_USER",
					Message: "User with this email doesn't found",
				},
			},
		})
		return
	}
	accessToken, err := tokenutil.CreateAccessToken(&user, uc.Env.AccessTokenSecret, uc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "TOKEN_ERROR",
					Message: "Error to create access token",
				},
			},
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{Result: accessToken})
}

func (lc *SignupController) Logout(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	
	splitToken := strings.Split(tokenString, " ")
	if len(splitToken) != 2 {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_TOKEN",
					Message: "token doesn't provided",
				},
			},
		})
		return
	}

	token, err := jwt.Parse(splitToken[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(lc.Env.AccessTokenSecret), nil
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "INVALID_TOKEN",
					Message: "token isn't correct",
				},
			},
		})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "INVALID_TOKEN_CLAIMS",
					Message: "token isn't correct",
				},
			},
		})
		return
	}

	claims["exp"] = time.Now().Unix()

	c.JSON(http.StatusOK, models.SuccessResponse{Result: "Successfully logout"})
}

func GenerateRandomPassword(size int) string {
	var alpha = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	password := make([]rune, size)
	for i := 0; i < size; i++ {
		password[i] = alpha[rand.Intn(len(alpha)-1)]
	}
	hashPassword := string(password)
	return hashPassword
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}

	var (
		hasUpper, hasLower, hasDigit bool
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		// case unicode.IsPunct(char) || unicode.IsSymbol(char):
		// 	hasSpecial = true
		}
	}
	if !hasUpper || !hasLower || !hasDigit {
		return fmt.Errorf("password must contain at least one uppercase letter, one lowercase letter and one digit")
	}
	return nil
}

