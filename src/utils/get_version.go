package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBuildVersion(c *gin.Context) int {
	version, err := strconv.Atoi(c.Request.Header.Get("Build-Version"))

	if err != nil {
		return 0
	}

	return version
}

func GetOs(c *gin.Context) string {
	return c.Request.Header.Get("Platform")

}
