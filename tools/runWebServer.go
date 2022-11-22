package tools

// This a simple gin server that exposes the output of checkRedisFunctionality
// on the /metrics endpoint
import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RunWebServer(port int) {
	str_port := ":" + strconv.Itoa(port)
	log.Println("Starting the server on port", port)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/metrics", func(c *gin.Context) {
		c.String(200, CheckRedisFunctionality())
	})
	log.Fatal(r.Run(str_port))
}
