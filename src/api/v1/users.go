package v1

import (
	"fmt"
	"net/http"
	"tk_app_v1/src/models"

	"tk_app_v1/src/database"

	"github.com/gin-gonic/gin"
)

func GetWorkers(c *gin.Context) {
	workers := []models.Worker{}
	err := database.Connection.Db.Select(&workers, `
		SELECT id, name, firstname, patronomic 
		FROM workers 
		ORDER BY firstname
	`)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, err)
		return
	}
	c.JSON(200, models.ServerResponse{Items: workers})
}

func CreateWorker(c *gin.Context) {
	var worker models.Worker

	if err := c.ShouldBindJSON(&worker); err != nil {
		c.JSON(400, err)
		return
	}

	query := `INSERT INTO workers (name, firstname, patronomic) VALUES ($1, $2, $3) RETURNING id`

	var id int
	err := database.Connection.Db.QueryRow(query, worker.Name, worker.Firstname, worker.Patronomic).Scan(&id)
	if err != nil {
		c.JSON(400, err)
		return
	}

	worker.ID = id
	c.JSON(http.StatusCreated, worker)
}
