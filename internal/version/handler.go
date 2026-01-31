package version

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func VersionHandler(c *gin.Context) {

	c.JSON(http.StatusOK, VersionResponse{
		Version:   Version,
		Commit:    Commit,
		BuildTime: BuildTime,
	})
}
