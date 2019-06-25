package router

import (
	"log"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/iceyang/gin-exam/pkg"
)

func responseHandler(c *gin.Context) {
	c.Next()
	if c.Writer.Status() == http.StatusNotFound && c.Writer.Size() <= 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "找不到资源",
		})
		return
	}
	if c.Writer.Status() == http.StatusOK && c.Writer.Size() <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
		return
	}
}

func recovery(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			buf := make([]byte, 65536)
			buf = buf[:runtime.Stack(buf, false)]
			e, ok := err.(*pkg.Error)
			if ok {
				if e.Code >= 500 {
					// 错误发送到kibana
					log.Printf("%s\n%s", err, buf)
				}
				c.AbortWithStatusJSON(e.Code, gin.H{
					"message": e.Msg,
				})
				return
			}
			// 错误发送到kibana
			log.Printf("%s\n%s", err, buf)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "服务出错，请稍后尝试",
			})
		}
	}()
	c.Next()
}
