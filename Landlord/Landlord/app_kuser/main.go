package main

import (
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// healthy
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// action
	r.GET("/:action", func(c *gin.Context) {
		kuser, err := exec.Command("/bin/bash", "ckuser", c.ClientIP()).Output()
		if err == nil {
			rp, _ := exec.Command("/bin/bash", c.Params.ByName("action"), string(kuser)).Output()
			c.String(http.StatusOK, "%s", string(rp))
		} else {
			c.String(http.StatusOK, "%s", err)
		}
	})

	// 綁定 port number
	r.Run(":80") // listen and serve on 0.0.0.0:80 (for windows "localhost:8080")
}
