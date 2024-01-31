package tickethandler

import (
	"biletes/pkg/colors"
	"biletes/pkg/db"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

func deleteTicketById(c *gin.Context) {
	deletedTicket, err := db.Db.Exec("DELETE FROM `tickets` WHERE id = ?;", c.Param("id"))
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
}

func updateTicketById(c *gin.Context) {
	var request Ticket
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := db.Db.Exec("UPDATE `tickets` SET `eventId`= ?,`eventName`= ?,`eventDescription`= ?,`ticketType`= ?,`ticketSeat`= ?,`used`= ?,`qrPath`= ?,`ticketDate`= ?,`ticketLocation`= ? WHERE id = ?;", request.EventId, request.EventName, request.EventDescription, request.TicketTypeId, request.TicketSeat, request.TicketUsed, request.TickedQRPath, request.TicketDate, request.TicketLocation, c.Param("id"))
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
}

func getAllTickets(c *gin.Context) {
	tickets := make([]*Ticket, 0)
	rows, err := db.Db.Query("SELECT * FROM `tickets`;")

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
}

func getTicketById(c *gin.Context) {
	var dbData Ticket
	event := db.Db.QueryRow("SELECT * FROM `tickets` WHERE id = ?;", c.Param("id")).Scan(&dbData.Id, &dbData.EventId, &dbData.EventName, &dbData.EventDescription, &dbData.TicketTypeId, &dbData.TicketUID, &dbData.TicketSeat, &dbData.TicketUsed, &dbData.TickedQRPath, &dbData.TicketDate, &dbData.TicketLocation, &dbData.TicketCreatedAt, &dbData.TicketUpdatedAt)
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
