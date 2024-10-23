package controllers

import (
	"fmt"
	"net/http"
	"time"

	db "HR_management_system/database"
	"HR_management_system/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your_secret_key") // Change this in production!

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// -----------------------------------------
func CheckUserExists(username string) bool {
	var count int
	fmt.Println("CheckUserExists_username : ", username)
	err := db.DB.Get(&count, "SELECT COUNT(*) FROM users WHERE username = $1", username)
	if err != nil {
		fmt.Printf("CHeckUserError")
	}
	//fmt.Println("CheckUserExists_count : ", count)
	//fmt.Println("CheckUserExists : ", err)

	//fmt.Println()
	if count == 0 {
		return false
	} else {
		return true
	}
}

// ------------------------------------------
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	//fmt.Println("Registration started : ", user)
	exists := CheckUserExists(user.Username) // check if user exists
	//fmt.Print("Username exists :", exists)
	if exists {
		fmt.Printf("Username already exists !")
		return
	}
	_, err := db.DB.NamedExec("INSERT INTO users (username, password) VALUES (:username, :password)", &user)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "username already exists"})
		return
	}
	fmt.Println("Registration Completed : ", user)
		c.JSON(http.StatusCreated, gin.H{"message": "user registered"})
}

// ------------------------------------------
func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//log.Println("Recieved username : ", user)
	var dbUser models.User
	//fmt.Println("db,DB: ", db.DB)
	err := db.DB.Get(&dbUser, "SELECT username, password FROM users WHERE username=$1", user.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User does not exist"})
		return
	}
	fmt.Println("Line 62 - Junaid - Username Found !")

	// if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(dbUser.Password)); err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "password error@"})
	// 	return
	// }

	claims := &Claims{
		Username: dbUser.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

//------------------------------------------
