package userhandler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"biletes/pkg/colors"
	"biletes/pkg/db"

	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Id       int    `form:"id"`
	Role     int    `form:"role"`
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Email    string `form:"email" binding:"required"`
}

type Login struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func login(c *gin.Context) {
	// jwtKey := []byte(os.Getenv("JWT_SECRET"))
	var request Login
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dbData User
	username := db.Db.QueryRow("SELECT password, roleId, id FROM `users` WHERE username = ? ;", request.Username).Scan(&dbData.Password, &dbData.Role, &dbData.Id)
	switch {
	case username == sql.ErrNoRows:
		// not found
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	case username != nil:
		// query failed
		fmt.Println(colors.ColorRed, "Query failed\n"+username.Error(), colors.ColorEnd)
		c.JSON(http.StatusInternalServerError, gin.H{"error": username.Error()})
		// log.Fatal(err)
		return
	default:
		// found
		match, err := argon2id.ComparePasswordAndHash(request.Password, dbData.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Couldn't verify password: " + err.Error(),
			})
			log.Fatal(err)
		}
		if !match {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
			return
		}

		// Create a new token
		t := jwt.NewWithClaims(jwt.SigningMethodHS256,
			jwt.MapClaims{
				// "iss":      "my-auth-server",
				"Iat":       time.Now().Unix(),
				"ExpiresAt": time.Now().Add(time.Minute * 5).Unix(),
				"username":  request.Username,
				"role":      dbData.Role,
				"userId":    dbData.Id,
			})
		// t := jwt.New(jwt.SigningMethodHS256)
		signed, err := t.SignedString(jwtKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't sign JWT token! " + err.Error()})
			// log.Fatal(err)
		}

		// insert token in db
		_, err = db.Db.Exec("INSERT INTO tokens (userId, jwtToken) VALUES (?, ?)", dbData.Id, signed)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store token in database." + err.Error()})
		}

		c.JSON(http.StatusOK, signed)
		return
	}
}

func register(c *gin.Context) {
	var request User
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dbData User
	username := db.Db.QueryRow("SELECT password FROM `users` WHERE username = ? OR email = ? ;", request.Username, request.Email).Scan(&dbData.Password)
	switch {
	case username == sql.ErrNoRows:
		// not found
		// hash password
		hash, err := argon2id.CreateHash(request.Password, argon2id.DefaultParams)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Couldn't hash password: " + err.Error(),
			})
			log.Fatal(err)
		}

		// insert into db
		_, err = db.Db.Exec("INSERT INTO users (username, password, email) VALUES (?, ?, ?)", request.Username, hash, request.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"success": "User Registered Sucessfully."})
		return
	case username != nil:
		// query failed
		fmt.Println(colors.ColorRed, "Query failed\n"+username.Error(), colors.ColorEnd)
		c.JSON(http.StatusInternalServerError, gin.H{"error": username.Error()})
		// log.Fatal(err)
		return
	default:
		// found
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already in database with Username or Email"})
		return
	}
}

func logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	bearerToken := strings.Split(authHeader, " ")

	if len(bearerToken) == 2 {
		token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtKey), nil
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JWT token." + err.Error()})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims)
			_, err = db.Db.Exec("DELETE FROM tokens WHERE jwtToken = ?", bearerToken[1])
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete token from database." + err.Error()})
			}
			c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
	}
}
