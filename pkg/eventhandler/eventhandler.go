package eventhandler

import (
	"biletes/pkg/colors"
	"biletes/pkg/db"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"image/color"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

type Event struct {
	Id               int    `form:"id"`
	EventName        string `form:"eventName" binding:"required"`
	EventDescription string `form:"eventDescription" binding:"required"`
	EventCategory    int    `form:"eventCategory" binding:"required"`
	EventDate        string `form:"eventDate" binding:"required"`
	EventLocation    string `form:"eventLocation" binding:"required"`
	EventImage       string `form:"eventImage" binding:"required"`
	EventBanner      string `form:"eventBanner" binding:"required"`
	EventCapacity    int    `form:"eventCapacity" binding:"required"`
	EventTicketTypes string `form:"ticketTypes" binding:"required"`
	SeatsRequired    int    `form:"seatsRequired" binding:"required"`
	EventCreatedAt   string `form:"created"`
	EventUpdatedAt   string `form:"updated"`
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
	SeatsRequired    int    `form:"seatsRequired" binding:"required"`
	EventCreatedAt   string `form:"created"`
	EventUpdatedAt   string `form:"updated"`
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

func createEvent(c *gin.Context) {
	var request Event
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bind error" + err.Error()})
		return
	}
	var dbData Event
	exists := db.Db.QueryRow("SELECT eventName FROM `events` WHERE eventName = ?;", request.EventName).Scan(&dbData.EventName)
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

		// Convert the JSON string to a map
		err := json.Unmarshal([]byte(request.EventTicketTypes), &ticketTypeMap)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// c.JSON(http.StatusOK, gin.H{"message": len(ticketTypeMap)})

		// ---
		eventId, err := db.Db.Exec("INSERT INTO `events` (`eventName`, `eventDescription`, `eventCategory`, `eventDate`, `eventLocation`, `eventImage`, `eventBanner`, `eventCapacity`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);", request.EventName, request.EventDescription, request.EventCategory, request.EventDate, request.EventLocation, request.EventImage, request.EventBanner, request.EventCapacity, request.SeatsRequired)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error: " + err.Error()})
		}
		lasteventId, err := eventId.LastInsertId()
		if err != nil {
			// failed to get last insert id
			panic(err)
		}

		var ticketSeat string = ""
		if request.SeatsRequired == 0 {
			ticketSeat = "None"
		}

		// loop for len of ticketTypeMap instead then another for count of tycket type
		for i := 1; i < len(ticketTypeMap)+1; i++ {
			// fmt.Println("i = ", i)
			for x := 1; x < int(ticketTypeMap[i].(float64))+1; x++ {
				// fmt.Println("x = ", x)
				generateTickets(db.Db, err, request, lasteventId, 0, nil, "", int64(i), ticketSeat, c, nil)
			}
			// generateTickets(db, err, request, lasteventId, 0, nil, "", int64(ticketTypes[1]), "A1", c, nil)
		}

		c.JSON(http.StatusOK, gin.H{"message": "Added event and tickets"})
		// ---

	case exists != nil:
		// query failed
		fmt.Print(colors.ColorRed, "Query failed\n"+exists.Error(), colors.ColorEnd, "\n")
		c.JSON(http.StatusInternalServerError, gin.H{"error": exists.Error()})
		// log.Fatal(err)
		return
	default:
		// found
		c.JSON(http.StatusBadRequest, gin.H{"error": "Event with name: " + request.EventName + " already in database!"})
		return
	}
}

func getEventById(c *gin.Context) {
	var dbData Event
	event := db.Db.QueryRow("SELECT * FROM `events` WHERE id = ?;", c.Param("id")).Scan(&dbData.Id, &dbData.EventName, &dbData.EventDescription, &dbData.EventCategory, &dbData.EventDate, &dbData.EventLocation, &dbData.EventImage, &dbData.EventBanner, &dbData.EventCapacity, &dbData.SeatsRequired, &dbData.EventCreatedAt, &dbData.EventUpdatedAt)
	switch {
	case event == sql.ErrNoRows:
		// not found
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	case event != nil:
		// query failed
		fmt.Print(colors.ColorRed, "Query failed\n"+event.Error(), colors.ColorEnd, "\n")
		c.JSON(http.StatusInternalServerError, gin.H{"error": event.Error()})
		// log.Fatal(err)
		return
	default:
		// found
		c.JSON(http.StatusOK, dbData)
		return
	}
}

func getAllEvents(c *gin.Context) {
	events := make([]*GetEvent, 0)
	rows, err := db.Db.Query("SELECT * FROM `events`;")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

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
}

func updateEventById(c *gin.Context) {
	var request Event
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Db.Exec("UPDATE `events` SET `eventName`= ?,`eventDescription`= ?,`eventCategory`= ?,`eventDate`= ?,`eventLocation`= ?,`eventImage`= ?,`eventCapacity`= ?, `seatsRequired`= ? WHERE id = ?;", &request.EventName, &request.EventDescription, &request.EventCategory, &request.EventDate, &request.EventLocation, &request.EventImage, &request.EventCapacity, &request.SeatsRequired, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"success": "Event updated Sucessfully."})
}

func deleteEventById(c *gin.Context) {
	event, err := db.Db.Exec("DELETE FROM `events` WHERE id = ?;", c.Param("id"))
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
		fmt.Print(colors.ColorRed, "Query failed\n", colors.ColorEnd, "\n")
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
