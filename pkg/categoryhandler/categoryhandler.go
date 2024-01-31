package categoryhandler

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

type Category struct {
	Id                int    `form:"id"`
	CategoryName      string `form:"name" binding:"required"`
	CategoryCreatedAt string `form:"created"`
	CategoryUpdatedAt string `form:"updated"`
}

func createCategory(c *gin.Context) {
	var request Category
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var dbData Category
	event := db.Db.QueryRow("SELECT * FROM `categories` WHERE categoryName = ?;", request.CategoryName).Scan(&dbData.CategoryName)
	switch {
	case event == sql.ErrNoRows:
		// not found
		// insert into db
		_, err := db.Db.Exec("INSERT INTO `categories`(`eventName`) VALUES (?);", request.CategoryName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"success": "Category created Sucessfully."})
		return
	case event != nil:
		// query failed
		fmt.Print(colors.ColorRed, "Query failed\n"+event.Error(), colors.ColorEnd, "\n")
		c.JSON(http.StatusInternalServerError, gin.H{"error": event.Error()})
		// log.Fatal(err)
		return
	default:
		// found
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category already in database with name: " + request.CategoryName})
		return
	}
}

func deleteCategoryById(c *gin.Context) {
	category, err := db.Db.Exec("DELETE FROM `categories` WHERE id = ?;", c.Param("id"))
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
}

func updateCategoryById(c *gin.Context) {
	var request Category
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Db.Exec("UPDATE `categories` SET `categoryName`= ? WHERE id = ?;", request.CategoryName, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, "Category updated successfully.")
}

func getAllCategories(c *gin.Context) {
	categories := make([]*Category, 0)
	rows, err := db.Db.Query("SELECT * FROM `categories`;")

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
}
