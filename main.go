package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"image/color"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	qrcode "github.com/skip2/go-qrcode"
)

var colorEnd string = "\033[0m"
var colorGreen string = "\033[32m"
var colorYellow string = "\033[33m"
var colorRed string = "\033[31m"
var jwtKey = []byte(os.Getenv("JWT_SECRET"))

// colorBlue := "\033[34m"
// colorPurple := "\033[35m"
// colorCyan := "\033[36m"
// colorWhite := "\033[37m"

type User struct {
	Id       int    `form:"id"`
	Role     int    `form:"role"`
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Email    string `form:"email" binding:"required"`
}

type PasswordReset struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type Login struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// type GenerateTicket struct {
// 	Result string `form:"content" binding:"required"`
// }

type Category struct {
	Id                int    `form:"id"`
	CategoryName      string `form:"categoryName" binding:"required"`
	CategoryCreatedAt string `form:"created"`
	CategoryUpdatedAt string `form:"updated"`
}

type Event struct {
	Id               int    `form:"id"`
	EventName        string `form:"eventName" json:"eventName" binding:"required"`
	EventDescription string `form:"eventDescription" json:"eventDescription" binding:"required"`
	EventCategory    int    `form:"eventCategory" json:"eventCategory" binding:"required"`
	EventDate        string `form:"eventDate" json:"eventDate" binding:"required"`
	EventLocation    string `form:"eventLocation" json:"eventLocation" binding:"required"`
	EventImage       string `form:"eventImage" json:"eventImage" binding:"required"`
	EventBanner      string `form:"eventBanner" json:"eventBanner" binding:"required"`
	EventCapacity    int    `form:"eventCapacity" json:"eventCapacity" binding:"required"`
	EventTicketTypes string `form:"ticketTypes" json:"ticketTypes" binding:"required"`
	SeatsRequired    string `form:"seatsRequired"  json:"seatsRequired"`
	// SeatsRequired    int    `form:"seatsRequired" binding:"required"`
	EventCreatedAt string `form:"created"`
	EventUpdatedAt string `form:"updated"`
}

type GetEvent struct {
	Id               int    `form:"id"`
	EventName        string `form:"eventName" binding:"required"`
	EventDescription string `form:"eventDescription" binding:"required"`
	EventCategory    int    `form:"eventCategory" binding:"required"`
	EventDate        string `form:"eventDate" binding:"required"`
	EventLocation    string `form:"eventLocation" binding:"required"`
	EventImage       string `form:"eventImage" binding:"required"`
	EventBanner      string `form:"eventBanner" binding:"required"`
	EventCapacity    int    `form:"eventCapacity" binding:"required"`
	SeatsRequired    int    `form:"seatsRequired"`
	EventCreatedAt   string `form:"created"`
	EventUpdatedAt   string `form:"updated"`
}

type TicketTypes struct {
	Id                int     `form:"id"`
	TicketTypeName    string  `form:"typeName" json:"typeName" binding:"required"`
	TicketTypePrice   float64 `form:"typePrice" json:"typePrice" binding:"required"`
	TicketTypeCreated string  `form:"created"`
	TicketTypeUpdated string  `form:"updated"`
}

type Ticket struct {
	Id               int    `form:"id"`
	Title            string `form:"ticketTitle" binding:"required"`
	EventId          int    `form:"eventId"`
	EventName        string `form:"eventName"`
	EventDescription string `form:"eventDescription"`
	TicketTypeId     int    `form:"ticketTypeId"`
	TicketUID        string `form:"ticketUID"`
	TicketSeat       string `form:"ticketSeat" binding:"required"`
	TicketUsed       int    `form:"ticketUsed"`
	TickedQRPath     string `form:"qrPath"`
	TicketDate       string `form:"ticketDate"`
	TicketLocation   string `form:"ticketLocation"`
	TicketCreatedAt  string `form:"created"`
	TicketUpdatedAt  string `form:"updated"`
}

type Review struct {
	Id              int    `form:"id"`
	EventId         int    `form:"eventId" binding:"required"`
	EventName       string `form:"eventName"`
	UserId          int    `form:"userId"`
	UserName        string `form:"userName"`
	Rating          int    `form:"rating" binding:"required"`
	Comment         string `form:"comment" binding:"required"`
	ReviewCreatedAt string `form:"reviewCreated"`
	ReviewUpdatedAt string `form:"reviewUpdated"`
}

type BoughtTicket struct {
	Id                   int    `form:"id"`
	TicketId             int    `form:"ticketId"`
	UserId               int    `form:"userId" binding:"required"`
	UserName             string `form:"userName"`
	EventName            string `form:"eventName"`
	TicketBoughtAt       string `form:"purchaseCreated"`
	TicketBoughtAtUpdate string `form:"purchaseUpdated"`
	QrPath               string `form:"qrPath"`
	TicketDate           string `form:"ticketDate"`
	TicketLocation       string `form:"ticketLocation"`
	TicketSeat           string `form:"ticketSeat"`
}

type ReviewValid struct {
	TicketId   int `form:"ticketId"`
	UserId     int `form:"userId" binding:"required"`
	TicketUsed int `form:"ticketUsed"`
}

func main() {

	// load enviornment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// stripe stuff don't touch
	// config := &stripe.BackendConfig{
	// 	MaxNetworkRetries: stripe.Int64(0), // Zero retries
	// }
	// sc := &client.API{}
	// sc.Init("sk_key", &stripe.Backends{
	// 	API:     stripe.GetBackendWithConfig(stripe.APIBackend, config),
	// 	Uploads: stripe.GetBackendWithConfig(stripe.UploadsBackend, config),
	// })
	// stripe.Init(os.Getenv("STRIPE_API_KEY"), nil)

	// connect to the database
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/biletes")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	// check the connection
	err = db.Ping()
	if err != nil {
		fmt.Print(colorRed, "Not Connected to DB!", colorEnd, "\n")
		log.Fatal(colorRed, err.Error(), colorEnd, "\n")
	}
	fmt.Print(colorGreen, "Connected to DB!", colorEnd, "\n")

	// _, err = db.Exec("CREATE DATABASE IF NOT EXISTS biletes")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("DB Created Susseccfully!")

	// Set the router as the default one provided by Gin
	r := gin.Default()
	r.ForwardedByClientIP = true
	r.SetTrustedProxies([]string{"http://localhost:5173", "http://127.0.0.1:5173", "http://localhost:8000", "http://127.0.0.1:8000"})
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Content-Type", "Authorization"}

	// r.Use(cors.Default())
	r.Use(cors.New(config))

	// serve static folder for qr codes
	r.Static("/codes", "./codes")

	// register
	r.POST("/api/register", func(c *gin.Context) {
		var request User
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var dbData User
		username := db.QueryRow("SELECT password FROM `users` WHERE username = ? OR email = ? ;", request.Username, request.Email).Scan(&dbData.Password)
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
			_, err = db.Exec("INSERT INTO users (username, password, email) VALUES (?, ?, ?)", request.Username, hash, request.Email)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			c.JSON(http.StatusOK, "User Registered Sucessfully.")
			return
		case username != nil:
			// query failed
			fmt.Println(colorRed, "Query failed\n"+username.Error(), colorEnd)
			c.JSON(http.StatusInternalServerError, gin.H{"error": username.Error()})
			// log.Fatal(err)
			return
		default:
			// found
			c.JSON(http.StatusBadRequest, gin.H{"error": "User already in database with Username or Email"})
			return
		}

	})

	// reset users password
	r.PUT("/api/password/reset", func(c *gin.Context) {
		var request PasswordReset
		var dbData PasswordReset
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		found := db.QueryRow("SELECT password FROM `users` WHERE email = ? ;", request.Email).Scan(&dbData.Password)
		switch {
		case found == sql.ErrNoRows:
			// not found

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "User with the specified email doesn't exist in database.",
			})
		case found != nil:
			// query failed
			fmt.Println(colorRed, "Query failed\n"+found.Error(), colorEnd)
			c.JSON(http.StatusInternalServerError, gin.H{"error": found.Error()})
			// log.Fatal(err)
			return
		default:
			// found
			// update password into db
			hash, err := argon2id.CreateHash(request.Password, argon2id.DefaultParams)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Couldn't hash password: " + err.Error(),
				})
				log.Fatal(err)
			}
			_, err = db.Exec("UPDATE `users` SET `password`= ? WHERE `email` = ?;", hash, request.Email)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			c.JSON(http.StatusOK, "Password Reset Sucessfully.")
		}
	})

	// login
	r.POST("/api/login", func(c *gin.Context) {
		var request Login
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var dbData User
		username := db.QueryRow("SELECT password, roleId, id FROM `users` WHERE username = ? ;", request.Username).Scan(&dbData.Password, &dbData.Role, &dbData.Id)
		switch {
		case username == sql.ErrNoRows:
			// not found
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		case username != nil:
			// query failed
			fmt.Println(colorRed, "Query failed\n"+username.Error(), colorEnd)
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
			_, err = db.Exec("INSERT INTO tokens (userId, jwtToken) VALUES (?, ?)", dbData.Id, signed)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store token in database." + err.Error()})
			}

			c.JSON(http.StatusOK, signed)
			return
		}

	})

	// logout
	r.POST("/api/logout", func(c *gin.Context) {
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
				_, err = db.Exec("DELETE FROM tokens WHERE jwtToken = ?", bearerToken[1])
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete token from database." + err.Error()})
				}
				c.JSON(http.StatusOK, "Successfully logged out")
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		}
	})

	// get all events
	r.GET("/api/events", func(c *gin.Context) {
		events := make([]*GetEvent, 0)
		rows, err := db.Query("SELECT * FROM `events`;")

		if err != nil {
			log.Fatal(err)
		}
		// defer rows.Close()

		for rows.Next() {
			event := new(GetEvent)
			if err := rows.Scan(&event.Id, &event.EventName, &event.EventDescription, &event.EventCategory, &event.EventDate, &event.EventLocation, &event.EventImage, &event.EventBanner, &event.EventCapacity, &event.SeatsRequired, &event.EventCreatedAt, &event.EventUpdatedAt); err != nil {
				panic(err)
			}
			events = append(events, event)
		}
		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			panic(err)
		}
		c.JSON(http.StatusOK, events)
	})

	// create event & auto generate tickets
	r.POST("/api/event/create", func(c *gin.Context) {
		// fmt.Println(c.GetPostForm("ticketTypes"))
		// fmt.Println(c.GetPostForm("a"))
		var request Event
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bind error" + err.Error()})
			return
		}
		var dbData Event
		exists := db.QueryRow("SELECT eventName FROM `events` WHERE eventName = ?;", request.EventName).Scan(&dbData.EventName)
		switch {
		case exists == sql.ErrNoRows:
			// event not found
			// insert new event into db
			var ticketTypeMap map[int]interface{}

			// check if the JSON string is empty
			if request.EventTicketTypes == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Ticket Types can't be empty"})
				return
			}

			// check if seats required is false
			var ticketSeat string = "A1"
			if request.SeatsRequired == "" {
				ticketSeat = "None"
				request.SeatsRequired = "0"
			}

			// Convert the JSON string to a map
			err := json.Unmarshal([]byte(request.EventTicketTypes), &ticketTypeMap)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			// c.JSON(http.StatusOK, gin.H{"message": len(ticketTypeMap)})

			// ---
			eventId, err := db.Exec("INSERT INTO `events` (`eventName`, `eventDescription`, `eventCategory`, `eventDate`, `eventLocation`, `eventImage`, `eventBanner`, `eventCapacity`, `seatsRequired`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);", request.EventName, request.EventDescription, request.EventCategory, request.EventDate, request.EventLocation, request.EventImage, request.EventBanner, request.EventCapacity, request.SeatsRequired)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error: " + err.Error()})
			}
			lasteventId, err := eventId.LastInsertId()
			if err != nil {
				// failed to get last insert id
				panic(err)
			}

			// loop for len of ticketTypeMap instead then another for count of tycket type
			for i := 1; i < len(ticketTypeMap)+1; i++ {
				// fmt.Println("i = ", i)
				for x := 1; x < int(ticketTypeMap[i].(float64))+1; x++ {
					// fmt.Println("x = ", x)
					generateTickets(db, err, request, lasteventId, 0, nil, "", int64(i), ticketSeat, c, nil)
				}
				// generateTickets(db, err, request, lasteventId, 0, nil, "", int64(ticketTypes[1]), "A1", c, nil)
			}

			c.JSON(http.StatusOK, "Added event "+request.EventName+" and created tickets.")
			// ---
		case exists != nil:
			// query failed
			fmt.Println(colorRed, "Query failed\n"+exists.Error(), colorEnd)
			c.JSON(http.StatusInternalServerError, gin.H{"error": exists.Error()})
			// log.Fatal(err)
			return
		default:
			// found
			c.JSON(http.StatusBadRequest, gin.H{"error": "Event with name: " + request.EventName + " already in database!"})
			return
		}
	})

	// get event by id
	r.GET("/api/event/:id", func(c *gin.Context) {
		var dbData Event
		event := db.QueryRow("SELECT * FROM `events` WHERE id = ?;", c.Param("id")).Scan(&dbData.Id, &dbData.EventName, &dbData.EventDescription, &dbData.EventCategory, &dbData.EventDate, &dbData.EventLocation, &dbData.EventImage, &dbData.EventBanner, &dbData.EventCapacity, &dbData.SeatsRequired, &dbData.EventCreatedAt, &dbData.EventUpdatedAt)
		switch {
		case event == sql.ErrNoRows:
			// not found
			c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
			return
		case event != nil:
			// query failed
			fmt.Println(colorRed, "Query failed\n"+event.Error(), colorEnd)
			c.JSON(http.StatusInternalServerError, gin.H{"error": event.Error()})
			// log.Fatal(err)
			return
		default:
			// found
			c.JSON(http.StatusOK, dbData)
			return
		}
	})

	// update event
	r.PUT("/api/event/:id", func(c *gin.Context) {
		var request GetEvent
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, err := db.Exec("UPDATE `events` SET `eventName`= ?,`eventDescription`= ?,`eventCategory`= ?,`eventDate`= ?,`eventLocation`= ?,`eventImage`= ?,`eventCapacity`= ?, `seatsRequired`= ? WHERE id = ?;", &request.EventName, &request.EventDescription, &request.EventCategory, &request.EventDate, &request.EventLocation, &request.EventImage, &request.EventCapacity, &request.SeatsRequired, c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, "Event updated Sucessfully.")
	})

	// delete event by id
	r.DELETE("/api/event/:id", func(c *gin.Context) {
		event, err := db.Exec("DELETE FROM `events` WHERE id = ?;", c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		affected, err := event.RowsAffected()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		if affected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
			return
		}
		c.JSON(http.StatusOK, "Event deleted successfully. "+strconv.FormatInt(affected, 10)+" rows affected")
	})

	// insert category
	r.POST("/api/category/create", func(c *gin.Context) {
		var request Category
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var dbData Category
		event := db.QueryRow("SELECT * FROM `categories` WHERE categoryName = ?;", request.CategoryName).Scan(&dbData.CategoryName)
		switch {
		case event == sql.ErrNoRows:
			// not found
			// insert into db
			_, err = db.Exec("INSERT INTO `categories`(`categoryName`) VALUES (?);", request.CategoryName)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			c.JSON(http.StatusOK, "Category created Sucessfully.")
			return
		case event != nil:
			// query failed
			fmt.Println(colorRed, "Query failed\n"+event.Error(), colorEnd)
			c.JSON(http.StatusInternalServerError, gin.H{"error": event.Error()})
			// log.Fatal(err)
			return
		default:
			// found
			c.JSON(http.StatusBadRequest, gin.H{"error": "Category already in database with name: " + request.CategoryName})
			return
		}
	})

	// delete category by id
	r.DELETE("/api/category/:id", func(c *gin.Context) {
		category, err := db.Exec("DELETE FROM `categories` WHERE id = ?;", c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		affected, err := category.RowsAffected()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		if affected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
			return
		}
		c.JSON(http.StatusOK, "Category deleted successfully. "+strconv.FormatInt(affected, 10)+" rows affected")
	})

	// update category by id
	r.PUT("/api/category/:id", func(c *gin.Context) {
		var request Category
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, err := db.Exec("UPDATE `categories` SET `categoryName`= ? WHERE id = ?;", request.CategoryName, c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, "Category updated successfully.")
	})

	// get all categories
	r.GET("/api/categories", func(c *gin.Context) {
		categories := make([]*Category, 0)
		rows, err := db.Query("SELECT * FROM `categories`;")

		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			category := new(Category)
			if err := rows.Scan(&category.Id, &category.CategoryName, &category.CategoryCreatedAt, &category.CategoryUpdatedAt); err != nil {
				panic(err)
			}
			categories = append(categories, category)
		}
		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			panic(err)
		}
		c.JSON(http.StatusOK, categories)
	})

	// get category by id
	r.GET("/api/category/:id", func(c *gin.Context) {
		var dbData Category
		category := db.QueryRow("SELECT * FROM `categories` WHERE id = ?;", c.Param("id")).Scan(&dbData.Id, &dbData.CategoryName, &dbData.CategoryCreatedAt, &dbData.CategoryUpdatedAt)
		switch {
		case category == sql.ErrNoRows:
			// not found
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
			return
		case category != nil:
			// query failed
			fmt.Println(colorRed, "Query failed\n"+category.Error(), colorEnd)
			c.JSON(http.StatusInternalServerError, gin.H{"error": category.Error()})
			// log.Fatal(err)
			return
		default:
			// found
			c.JSON(http.StatusOK, dbData)
			return
		}
	})

	// get ticket type by id
	r.GET("/api/ticket/type/:id", func(c *gin.Context) {
		var dbData TicketTypes
		category := db.QueryRow("SELECT * FROM `tickettypes` WHERE id = ?;", c.Param("id")).Scan(&dbData.Id, &dbData.TicketTypeName, &dbData.TicketTypePrice, &dbData.TicketTypeCreated, &dbData.TicketTypeUpdated)
		switch {
		case category == sql.ErrNoRows:
			// not found
			c.JSON(http.StatusNotFound, gin.H{"error": "Ticket type not found"})
			return
		case category != nil:
			// query failed
			fmt.Println(colorRed, "Query failed\n"+category.Error(), colorEnd)
			c.JSON(http.StatusInternalServerError, gin.H{"error": category.Error()})
			// log.Fatal(err)
			return
		default:
			// found
			c.JSON(http.StatusOK, dbData)
			return
		}
	})

	// insert ticket type
	r.POST("/api/ticket/type/create", func(c *gin.Context) {
		var request TicketTypes
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var dbData TicketTypes
		ticketType := db.QueryRow("SELECT * FROM `tickettypes` WHERE ticketName = ?;", request.TicketTypeName).Scan(&dbData.TicketTypeName)
		switch {
		case ticketType == sql.ErrNoRows:
			// not found
			// insert into db
			_, err = db.Exec("INSERT INTO `tickettypes`(`ticketName`, `typePrice`) VALUES ( ?, ?);", request.TicketTypeName, request.TicketTypePrice)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			c.JSON(http.StatusOK, "Category created Sucessfully.")
			return
		case ticketType != nil:
			// query failed
			fmt.Println(colorRed, "Query failed\n"+ticketType.Error(), colorEnd)
			c.JSON(http.StatusInternalServerError, gin.H{"error": ticketType.Error()})
			// log.Fatal(err)
			return
		default:
			// found
			c.JSON(http.StatusBadRequest, gin.H{"error": "Category already in database with name: " + request.TicketTypeName})
			return
		}
	})

	// get all ticket types
	r.GET("/api/ticket/types", func(c *gin.Context) {
		types := make([]*TicketTypes, 0)
		rows, err := db.Query("SELECT * FROM `tickettypes`;")

		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			ticketType := new(TicketTypes)
			if err := rows.Scan(&ticketType.Id, &ticketType.TicketTypeName, &ticketType.TicketTypePrice, &ticketType.TicketTypeCreated, &ticketType.TicketTypeUpdated); err != nil {
				panic(err)
			}
			types = append(types, ticketType)
		}
		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			panic(err)
		}
		c.JSON(http.StatusOK, types)
	})

	// delete ticket type by id
	r.DELETE("/api/ticket/type/:id", func(c *gin.Context) {
		category, err := db.Exec("DELETE FROM `tickettypes` WHERE id = ?;", c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		affected, err := category.RowsAffected()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		if affected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Ticket type not found"})
			return
		}
		c.JSON(http.StatusOK, "Ticket type deleted successfully. "+strconv.FormatInt(affected, 10)+" rows affected")
	})

	// update ticket type by id
	r.PUT("/api/ticket/type/:id", func(c *gin.Context) {
		var request TicketTypes
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		category, err := db.Exec("UPDATE `tickettypes` SET `ticketName`= ?, `typePrice` = ? WHERE `id` = ?;", request.TicketTypeName, request.TicketTypePrice, c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		affected, err := category.RowsAffected()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		if affected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Ticket type not found"})
			return
		}
		c.JSON(http.StatusOK, "Ticket type updated successfully. "+strconv.FormatInt(affected, 10)+" rows affected")
	})

	// get all bought tickets
	r.GET("/api/bought/tickets", func(c *gin.Context) {
		tickets := make([]*BoughtTicket, 0)
		rows, err := db.Query("SELECT boughttickets.id, boughttickets.ticketId, boughttickets.userId, boughttickets.createdAt, boughttickets.updatedAt, users.username, tickets.eventName FROM `boughttickets` INNER JOIN users ON boughttickets.userId = users.id INNER JOIN tickets ON boughttickets.ticketId = tickets.id;")

		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			ticket := new(BoughtTicket)
			if err := rows.Scan(&ticket.Id, &ticket.TicketId, &ticket.UserId, &ticket.TicketBoughtAt, &ticket.TicketBoughtAtUpdate, &ticket.UserName, &ticket.EventName); err != nil {
				panic(err)
			}
			tickets = append(tickets, ticket)
		}
		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			panic(err)
		}
		c.JSON(http.StatusOK, tickets)
	})

	// get all users bought tickets
	r.GET("/api/:userId/tickets", func(c *gin.Context) {
		tickets := make([]*BoughtTicket, 0)
		rows, err := db.Query("SELECT boughttickets.id, boughttickets.ticketId, boughttickets.userId, boughttickets.createdAt, boughttickets.updatedAt, users.username, tickets.eventName FROM `boughttickets` INNER JOIN users ON boughttickets.userId = users.id INNER JOIN tickets ON boughttickets.ticketId = tickets.id WHERE boughttickets.userId = ?;", c.Param("userId"))

		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			ticket := new(BoughtTicket)
			if err := rows.Scan(&ticket.Id, &ticket.TicketId, &ticket.UserId, &ticket.TicketBoughtAt, &ticket.TicketBoughtAtUpdate, &ticket.UserName, &ticket.EventName); err != nil {
				panic(err)
			}
			tickets = append(tickets, ticket)
		}
		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			panic(err)
		}
		c.JSON(http.StatusOK, tickets)
	})

	// get bought ticket by id
	r.GET("/api/bought/ticket/:id", func(c *gin.Context) {
		var bought BoughtTicket
		event := db.QueryRow("SELECT boughttickets.id, boughttickets.ticketId, boughttickets.userId, boughttickets.createdAt, boughttickets.updatedAt, users.username, tickets.eventName, tickets.qrPath, tickets.ticketDate, tickets.ticketLocation, tickets.ticketSeat FROM `boughttickets` INNER JOIN users ON boughttickets.userId = users.id INNER JOIN tickets ON boughttickets.ticketId = tickets.id WHERE boughttickets.id = ?;", c.Param("id")).Scan(&bought.Id, &bought.TicketId, &bought.UserId, &bought.TicketBoughtAt, &bought.TicketBoughtAtUpdate, &bought.UserName, &bought.EventName, &bought.QrPath, &bought.TicketDate, &bought.TicketLocation, &bought.TicketSeat)
		switch {
		case event == sql.ErrNoRows:
			// not found
			c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
			return
		case event != nil:
			// query failed
			fmt.Println(colorRed, "Query failed\n"+event.Error(), colorEnd)
			c.JSON(http.StatusInternalServerError, gin.H{"error": event.Error()})
			// log.Fatal(err)
			return
		default:
			// found
			c.JSON(http.StatusOK, bought)
			return
		}
	})

	// delete bought ticket by id
	r.DELETE("/api/bought/ticket/:id", func(c *gin.Context) {
		deletedTicket, err := db.Exec("DELETE FROM `boughttickets` WHERE id = ? AND userId = ?;", c.Param("id"), c.Param("userId"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		affected, err := deletedTicket.RowsAffected()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		if affected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "User's ticket not found"})
			return
		}
		c.JSON(http.StatusOK, "User's ticket deleted successfully. "+strconv.FormatInt(affected, 10)+" rows affected")
	})

	// update users bought ticket by id
	r.PUT("/api/bought/ticket/:id", func(c *gin.Context) {
		var request BoughtTicket
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		category, err := db.Exec("UPDATE `boughttickets` SET `ticketId`= ?, `userId` = ? WHERE `id` = ?;", request.TicketId, request.UserId, c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		affected, err := category.RowsAffected()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		if affected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "User hasn't bought ticket"})
			return
		}
		c.JSON(http.StatusOK, "User's bought ticket data updated successfully. "+strconv.FormatInt(affected, 10)+" rows affected")
	})

	// delete ticket by id
	r.DELETE("/api/ticket/:id", func(c *gin.Context) {
		deletedTicket, err := db.Exec("DELETE FROM `tickets` WHERE id = ?;", c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		affected, err := deletedTicket.RowsAffected()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		if affected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
			return
		}
		c.JSON(http.StatusOK, "Ticket deleted successfully. "+strconv.FormatInt(affected, 10)+" rows affected")
	})

	// update ticket by id
	r.PUT("/api/ticket/:id", func(c *gin.Context) {
		var request Ticket
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		category, err := db.Exec("UPDATE `tickets` SET `eventId`= ?,`eventName`= ?,`eventDescription`= ?,`ticketType`= ?,`ticketSeat`= ?,`used`= ?,`qrPath`= ?,`ticketDate`= ?,`ticketLocation`= ? WHERE id = ?;", request.EventId, request.EventName, request.EventDescription, request.TicketTypeId, request.TicketSeat, request.TicketUsed, request.TickedQRPath, request.TicketDate, request.TicketLocation, c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		affected, err := category.RowsAffected()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		if affected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Event ticket not found."})
			return
		}
		c.JSON(http.StatusOK, "Ticket data updated successfully. "+strconv.FormatInt(affected, 10)+" rows affected")
	})

	// get all tickets
	r.GET("/api/tickets", func(c *gin.Context) {
		tickets := make([]*Ticket, 0)
		rows, err := db.Query("SELECT * FROM `tickets`;")

		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			ticket := new(Ticket)
			if err := rows.Scan(&ticket.Id, &ticket.EventId, &ticket.EventName, &ticket.EventDescription, &ticket.TicketTypeId, &ticket.TicketUID, &ticket.TicketSeat, &ticket.TicketUsed, &ticket.TickedQRPath, &ticket.TicketDate, &ticket.TicketLocation, &ticket.TicketCreatedAt, &ticket.TicketUpdatedAt); err != nil {
				panic(err)
			}
			tickets = append(tickets, ticket)
		}
		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			panic(err)
		}
		c.JSON(http.StatusOK, tickets)
	})

	// get ticket by id
	r.GET("/api/ticket/:id", func(c *gin.Context) {
		var dbData Ticket
		event := db.QueryRow("SELECT * FROM `tickets` WHERE id = ?;", c.Param("id")).Scan(&dbData.Id, &dbData.EventId, &dbData.EventName, &dbData.EventDescription, &dbData.TicketTypeId, &dbData.TicketUID, &dbData.TicketSeat, &dbData.TicketUsed, &dbData.TickedQRPath, &dbData.TicketDate, &dbData.TicketLocation, &dbData.TicketCreatedAt, &dbData.TicketUpdatedAt)
		switch {
		case event == sql.ErrNoRows:
			// not found
			c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
			return
		case event != nil:
			// query failed
			fmt.Println(colorRed, "Query failed\n"+event.Error(), colorEnd)
			c.JSON(http.StatusInternalServerError, gin.H{"error": event.Error()})
			// log.Fatal(err)
			return
		default:
			// found
			c.JSON(http.StatusOK, dbData)
			return
		}
	})

	// get all reviews
	r.GET("/event/:eventId/reviews", func(c *gin.Context) {
		reviews := make([]*Review, 0)
		rows, err := db.Query("SELECT reviews.id, eventId, userId, rating, comment, reviews.createdAt, reviews.updatedAt, users.username FROM `reviews` INNER JOIN users ON reviews.userId = users.id WHERE eventId = ?;", c.Param("eventId"))

		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			review := new(Review)
			if err := rows.Scan(&review.Id, &review.EventId, &review.UserId, &review.Rating, &review.Comment, &review.ReviewCreatedAt, &review.ReviewUpdatedAt, &review.UserName); err != nil {
				panic(err)
			}
			reviews = append(reviews, review)
		}
		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			panic(err)
		}
		c.JSON(http.StatusOK, reviews)
	})

	// get review by id
	r.GET("/event/:eventId/review/:id", func(c *gin.Context) {
		var review Review
		event := db.QueryRow("SELECT reviews.id, eventId, userId, rating, comment, reviews.createdAt, reviews.updatedAt, users.username FROM `reviews` INNER JOIN users ON reviews.userId = users.id WHERE reviews.eventId = ? AND reviews.id = ?;", c.Param("eventId"), c.Param("id")).Scan(&review.Id, &review.EventId, &review.UserId, &review.Rating, &review.Comment, &review.ReviewCreatedAt, &review.ReviewUpdatedAt, &review.UserName)
		switch {
		case event == sql.ErrNoRows:
			// not found
			c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
			return
		case event != nil:
			// query failed
			fmt.Println(colorRed, "Query failed\n"+event.Error(), colorEnd)
			c.JSON(http.StatusInternalServerError, gin.H{"error": event.Error()})
			// log.Fatal(err)
			return
		default:
			// found
			c.JSON(http.StatusOK, review)
			return
		}
	})

	// delete review by id
	r.DELETE("/api/reviews/:id", func(c *gin.Context) {
		deletedReview, err := db.Exec("DELETE FROM `reviews` WHERE id = ?;", c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		affected, err := deletedReview.RowsAffected()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		if affected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
			return
		}
		c.JSON(http.StatusOK, "Review deleted successfully. "+strconv.FormatInt(affected, 10)+" rows affected")
	})

	// update review by id
	r.PUT("/api/reviews/:id", func(c *gin.Context) {
		var request Review
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		category, err := db.Exec("UPDATE `reviews` SET `eventId`= ?,`userId`= ?,`rating`= ?,`comment`= ? WHERE id = ?;", request.EventId, request.UserId, request.Rating, request.Comment, c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		affected, err := category.RowsAffected()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		if affected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Event review not found."})
			return
		}
		c.JSON(http.StatusOK, "Event review data updated successfully. "+strconv.FormatInt(affected, 10)+" rows affected")
	})

	// create review
	r.POST("/api/review", func(c *gin.Context) {
		// get user id from jwt
		var found ReviewValid
		authHeader := c.GetHeader("Authorization")
		bearerToken := strings.Split(authHeader, " ")
		jwtClaims := jwt.MapClaims{}

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
			// check token expiry
			expireTime := int64(claims["ExpiresAt"].(float64))
			now := time.Now().Unix()
			if now > expireTime {
				fmt.Println("Token is expired")
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Expired token"})
				return
			}
			jwtClaims = claims
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		}

		ticketUsed := db.QueryRow("SELECT boughttickets.ticketId, boughttickets.userId, tickets.used FROM `boughttickets` INNER JOIN tickets ON boughttickets.ticketId = tickets.id WHERE boughttickets.userId = ? AND tickets.used = 1;", jwtClaims["userId"]).Scan(&found.TicketId, &found.UserId, &found.TicketUsed)
		switch {
		case ticketUsed == sql.ErrNoRows:
			// not found
			c.JSON(http.StatusNotFound, gin.H{"error": "You must visit the event to post review"})
			return
		case ticketUsed != nil:
			// query failed
			fmt.Println(colorRed, "Query failed\n"+ticketUsed.Error(), colorEnd)
			c.JSON(http.StatusInternalServerError, gin.H{"error": ticketUsed.Error()})
			// log.Fatal(err)
			return
		default:
			// found
			var request Review
			if err := c.ShouldBind(&request); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			_, err = db.Exec("INSERT INTO `reviews`(`eventId`, `userId`, `rating`, `comment`) VALUES (?, ?, ?, ?);", request.EventId, jwtClaims["userId"], request.Rating, request.Comment)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			c.JSON(http.StatusOK, "Event review created successfully.")
			return
		}
	})

	fmt.Print(colorGreen, "Satarting server on http://localhost:", os.Getenv("PORT"), colorEnd, "\n")
	fmt.Print(colorRed, "Fix ticket auto SEAT generate", colorEnd, "\n")
	fmt.Print(colorRed, "Biletes izveidot PDF ar QR kodu", colorEnd, "\n")
	fmt.Print(colorRed, "Implement ticket VERIFICATION", colorEnd, "\n")
	fmt.Print(colorRed, "Atstat review tikai tas kurs nopircis bileti", colorEnd, "\n")
	fmt.Print(colorRed, "Izveidot ReviewHandler", colorEnd, "\n")
	fmt.Print(colorRed, "Izveidot BoughtTicketsHanlder", colorEnd, "\n")
	fmt.Print(colorRed, "Split code", colorEnd, "\n")

	// execute database migrations
	// execFile("roles.sql", db, nil, nil)
	// execFile("users.sql", db, nil, nil)
	// execFile("tokens.sql", db, nil, nil)
	// execFile("categories.sql", db, nil, nil)
	// execFile("events.sql", db, nil, nil)
	// execFile("ticket_types.sql", db, nil, nil)
	// execFile("tickets.sql", db, nil, nil)
	// execFile("bought_tickets.sql", db, nil, nil)
	// execFile("reviews.sql", db, nil, nil)
	// execSeeder("seeder.sql", db)
	// db.Close()
	r.Run(":" + os.Getenv("PORT")) // listen and serve on 0.0.0.0:8000
}

func execFile(filename string, conn *sql.DB, err error, dat []byte) {
	dat, err = os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	_, err = conn.Exec(string(dat))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s File %s executed successfully!%s\n", colorGreen, filename, colorEnd)
}

func execSeeder(filename string, conn *sql.DB) {
	dat, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	// Split the file data into separate statements
	requests := strings.Split(string(dat), ";")
	// Execute each statement
	for _, request := range requests {
		_, err = conn.Exec(request)
		if err != nil {
			panic(err)
		}
	}
	fmt.Printf("%s File %s executed successfully!%s\n", colorGreen, filename, colorEnd)
}

func generateTickets(conn *sql.DB, err error, request Event, lastInsertedEventId int64, ticketId int64, ticketBinIdError error, ticketBinId string, ticketType int64, ticketSeat string, c *gin.Context, ticket sql.Result) {
	ticket, err = conn.Exec("INSERT INTO `tickets`(`eventId`, `eventName`, `eventDescription`, `ticketType`, `ticketSeat`, `ticketDate`, `ticketLocation`) VALUES (?, ?, ?, ?, ?, ?, ?);", lastInsertedEventId, request.EventName, request.EventDescription, ticketType, ticketSeat, request.EventDate, request.EventLocation)
	if err != nil {
		// failed to insert
		// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		panic(err)
	}

	ticketId, err = ticket.LastInsertId()
	if err != nil {
		// failed to get last insert id
		panic(err)
	}

	ticketBinIdError = conn.QueryRow("SELECT ticketUID FROM `tickets` WHERE id = ?;", ticketId).Scan(&ticketBinId)
	switch {
	case ticketBinIdError == sql.ErrNoRows:
		// not found
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found."})
		return
	case ticketBinIdError != nil:
		// query failed
		fmt.Println(colorRed, "Query failed\n", colorEnd)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Query failed"})
		return
	default:
		// found
		bgColor, err := ParseHexColorFast("#B4ADEA")
		fgColor, err := ParseHexColorFast("#4F0147")
		err = qrcode.WriteColorFile(ticketBinId, qrcode.Medium, 256, bgColor, fgColor, "./codes/"+strconv.Itoa(int(ticketId))+".png")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "QR code couldn't be generated",
			})
			return
		}

		ticket, err = conn.Exec("UPDATE `tickets` SET qrPath = ? WHERE id = ?;", "./codes/"+strconv.Itoa(int(ticketId))+".png", ticketId)
		if err != nil {
			// failed to insert
			// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			panic(err)
		}
	}
	c.JSON(http.StatusOK, "Ticket for "+request.EventName+" successfully.")
}

var errInvalidFormat = errors.New("invalid format")

func ParseHexColorFast(s string) (c color.RGBA, err error) {
	c.A = 0xff

	if s[0] != '#' {
		return c, errInvalidFormat
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = errInvalidFormat
		return 0
	}

	switch len(s) {
	case 7:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		c.R = hexToByte(s[1]) * 17
		c.G = hexToByte(s[2]) * 17
		c.B = hexToByte(s[3]) * 17
	default:
		err = errInvalidFormat
	}
	return
}

// func CORSMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "false")

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(204)
// 			return
// 		}

// 		c.Next()
// 	}
// }
