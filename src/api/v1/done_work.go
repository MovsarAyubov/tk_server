package v1

import (
	"tk_app_v1/src/models"

	"tk_app_v1/src/database"

	"github.com/gin-gonic/gin"
)

func AddDoneWork(c *gin.Context) {

	var doneWork models.DoneWorkModel

	if err := c.ShouldBindJSON(&doneWork); err != nil {
		c.JSON(400, err)
		return
	}

	query := `INSERT INTO done_work (date, worker_id, type_of_work_id, cell_id, row_id, count, income) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	err := database.Connection.Db.QueryRow(query, doneWork.Date, doneWork.WorkerId, doneWork.TypeOfWorkId, doneWork.CellId, doneWork.RowId, doneWork.Count, doneWork.Income).Err()

	if err != nil {
		c.JSON(400, err)
		return
	}
	c.JSON(200, nil)
}
