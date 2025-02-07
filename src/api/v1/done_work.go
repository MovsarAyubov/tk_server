package v1

import (
	"fmt"
	"net/http"
	"strconv"
	"tk_app_v1/src/models"

	"tk_app_v1/src/database"

	"github.com/gin-gonic/gin"
)

func AddDoneWork(c *gin.Context) {

	// cellid := c.Query("cellId")
	// date := c.Query("date")

	var exists bool
	// checkQuery := `SELECT EXISTS(SELECT 1 FROM don_work WHERE date=$1 AND cell_id=$2)`
	// err := database.Connection.Db.QueryRow(checkQuery, date, cellid).Scan(&exists)
	// if err != nil {
	// 	c.JSON(500, err)
	// 	return
	// }

	if exists {
		c.JSON(409, gin.H{"error": "already exists"})
		return
	}

	var doneWork models.DoneWorkModel

	if err := c.ShouldBindJSON(&doneWork); err != nil {
		c.JSON(400, err)
		return
	}

	query := `INSERT INTO done_work (date, worker_id, type_of_work_id, cell_id, row_id, count, income) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	err := database.Connection.Db.QueryRow(query, doneWork.Date, doneWork.Worker_id, doneWork.Type_of_work_id, doneWork.Cell_id, doneWork.Row_id, doneWork.Count, doneWork.Income).Err()

	if err != nil {
		fmt.Println("Database error:", err)
		c.JSON(400, gin.H{
			"error":   "Database error",
			"details": err.Error(),
		})
		return
	}
	c.JSON(200, nil)
}

func FetchDonWorksByWorkerId(c *gin.Context) {
	workerIDStr := c.Query("workerId")
	if workerIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "workerId parameter is required"})
		return
	}
	fmt.Println("workerIDStr:", workerIDStr)

	// Convert worker_id to an integer
	workerID, err := strconv.Atoi(workerIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "workerId  must be a valid integer"})
		return
	}
	fmt.Println("workerID:", workerID)

	doneWorks := []models.DoneWorkModelResponse{}
	err = database.Connection.Db.Select(&doneWorks, `
		SELECT 
    		done_work.*, 
    		type_of_work.name AS type_of_work_name
		FROM 
   			 done_work
		JOIN 
    		type_of_work 
		ON 
    		done_work.type_of_work_id = type_of_work.id
		WHERE 
			worker_id = $1`,
		workerID,
	)
	if err != nil {
		fmt.Println("Database error:", err)
		c.JSON(400, gin.H{"error": "Database error", "details": err.Error()})
		return
	}
	c.JSON(200, models.DoneWorkServerResponse{Items: doneWorks})
}

func GetDoneWorkByDateAndCellID(c *gin.Context) {
	cellId := c.Query("cellId")
	date := c.Query("date")

	items := []models.CountAndIncome{}

	query := `
		SELECT count, income
		FROM done_work
		WHERE date = $1 AND cell_id = $2
	`
	err := database.Connection.Db.Select(&items, query, cellId, date)

	if err != nil {
		fmt.Println("Database error:", err)
		c.JSON(400, gin.H{"error": "Database error", "details": err.Error()})
		return
	}
	c.JSON(200, models.CountAndIncomeResponse{Items: items})
}
