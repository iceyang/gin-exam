package ctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExampleController struct{}

var Example = &ExampleController{}

func (ec *ExampleController) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (ec *ExampleController) NotFound(c *gin.Context) {
	c.Status(http.StatusNotFound)
}

func (ec *ExampleController) OK(c *gin.Context) {
	c.Status(http.StatusOK)
}
