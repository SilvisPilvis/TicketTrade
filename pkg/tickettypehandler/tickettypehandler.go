package tickettypehandler

import (
	"biletes/pkg/db"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TicketTypes struct {
	Id                int     `form:"id"`
	TicketTypeName    string  `form:"name" binding:"required"`
	TicketTypePrice   float64 `form:"price" binding:"required"`
	TicketTypeCreated string  `form:"created"`
	TicketTypeUpdated string  `form:"updated"`
}

func getAllTicketTypes(c *gin.Context) {
	types := make([]*TicketTypes, 0)
	rows, err := db.Db.Query("SELECT * FROM `tickettypes`;")

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
}

func deleteTicketTypeById(c *gin.Context) {
	category, err := db.Db.Exec("DELETE FROM `tickettypes` WHERE id = ?;", c.Param("id"))
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
}

func updateTicketTypeById(c *gin.Context) {
	var request TicketTypes
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := db.Db.Exec("UPDATE `tickettypes` SET `ticketName`= ?, `typePrice` = ? WHERE `id` = ?;", request.TicketTypeName, request.TicketTypePrice, c.Param("id"))
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
}
