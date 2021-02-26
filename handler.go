package log2file

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// Starts a router with paths to logfiles
func router() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	lg := r.Group("/logs")
	lg.GET("/log_m", getLog1)
	lg.GET("/log_b", getLog2)

	log.Println(r.Run(":" + strconv.Itoa(standOptions.router.port)))
}

// Handler for main log file
func getLog1(c *gin.Context) {
	c.File(fmt.Sprintf("./%s.%s", standOptions.fiNames.logMain, standOptions.fiNames.logExtension))
}

// Handler for backup log file
func getLog2(c *gin.Context) {
	c.File(fmt.Sprintf("./%s.%s", standOptions.fiNames.logBackup, standOptions.fiNames.logExtension))
}
