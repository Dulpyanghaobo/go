package routers

import (
	"github.com/gin-gonic/gin"
	"gin-blog/pkg/setting"
	"gin-blog/routers/api/v1"
)
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	/*
		r.GET("/test", func(c *gin.Context) {
			c.JSON(200,gin.H{
				"message": "test",
			})
		})
	*/
	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}
	return r
}
