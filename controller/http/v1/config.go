package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func newConfigurationRouter(handler *gin.RouterGroup, m map[string]interface{}) {
	h := handler.Group("/configuration")
	//r := &Rout{m: m}
	{
		for k, v := range m {
			h.GET(fmt.Sprintf("/%s", k), middleware(v))
		}
		//h.GET("/config", middlewareParam(m))
		//h.GET("/configg", r.getConfig)
		//h.GET("/history", r.history)
		//h.POST("/do-translate", r.doTranslate)
	}
	handler.GET("/config", middlewareParam(m))
}

func middlewareParam(v map[string]interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := c.Query("prm")
		c.JSON(http.StatusOK, v[t])
	}
}

func middleware(v interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, v)
	}
}
