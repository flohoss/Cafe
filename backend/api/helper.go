package api

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func getQueryAsUint64(c *gin.Context, query string) uint64 {
	str := c.Query(query)
	converted, _ := strconv.ParseUint(str, 10, 64)
	return converted
}

func getQueryAsFloat64(c *gin.Context, query string) float64 {
	str := c.Query(query)
	converted, _ := strconv.ParseFloat(str, 64)
	return converted
}
