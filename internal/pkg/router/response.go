package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/slipneff/gogger/log"
)

func newErrorResponse(c *gin.Context, statusCode int, err error) {
	log.Error(err, fmt.Sprintf("|%d|", statusCode), "HTTP Error")
	c.AbortWithStatusJSON(statusCode, err.Error())
}

func responseOK(c *gin.Context, obj any) {
	c.JSON(http.StatusOK, obj)
}
