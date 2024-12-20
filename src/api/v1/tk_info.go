package v1

import (
	"fmt"

	"tk_app_v1/src/models"

	"tk_app_v1/src/database"

	"github.com/gin-gonic/gin"
)

func GetTKInfo(c *gin.Context) {
	tkInfo := models.ResultServerResponse{}
	err := database.Connection.Db.Select(&tkInfo.TypesOfWork, `SELECT id, name, uom, price, period FROM type_of_work`)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, err)
		panic(err)
	}
	fmt.Println(err)
	err2 := database.Connection.Db.Select(&tkInfo.Cell, `SELECT id FROM cell`)

	if err2 != nil {
		fmt.Println(err2)
		c.JSON(400, err2)
		panic(err2)
	}
	err3 := database.Connection.Db.Select(&tkInfo.Row, `SELECT id FROM row`)
	fmt.Println(err3)
	if err3 != nil {
		fmt.Println(err3)
		c.JSON(400, err3)
		panic(err3)
	}
	c.JSON(200, models.ResultServerResponse{TypesOfWork: tkInfo.TypesOfWork, Cell: tkInfo.Cell, Row: tkInfo.Row})
}
