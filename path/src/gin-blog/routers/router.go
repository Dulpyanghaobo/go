package routers

import (
	"gin-blog/middleware/jwt"
	"gin-blog/routers/api"
	"github.com/gin-gonic/gin"
	"gin-blog/pkg/setting"
	"gin-blog/routers/api/v1"
)
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	r.GET("/auth",api.GetAuth)
	/*
		r.GET("/test", func(c *gin.Context) {
			c.JSON(200,gin.H{
				"message": "test",
			})
		})
	*/

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		/*
			注释
			标签页面路由
		*/
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//编辑文章
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		/*
			注释
			文章页面路由
		*/
		apiv1.GET("/articles",v1.GetArticles)
		apiv1.GET("/articles/:id",v1.GetArticle)
		apiv1.POST("/articles",v1.AddArticle)
		apiv1.PUT("/articles/:id",v1.EditArticle)
		apiv1.DELETE("/articles/:id",v1.DeleteArticle)


	}
	return r
}
