package main

import (
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

func SetupRouter() *gin.Engine {

	router := gin.New()
	//router.Use(gin.Recovery())
	//router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
	testGroup := router.Group("/test")
	{
		testGroup.GET("/address", address)
	}
	return router
}

func address(context *gin.Context) {
	context.JSON(http.StatusOK, GetLocalIP())
}

// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func main() {
	router := SetupRouter()
	router.Run(":8888")
}
