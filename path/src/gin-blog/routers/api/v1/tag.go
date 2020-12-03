package v1

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/astaxie/beego/core/validation"
	"github.com/Unknwon/com"
	"gin-blog/pkg/e"
	"gin-blog/models"
	"gin-blog/pkg/util"
	"gin-blog/pkg/setting"
)

//获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")
	maps :=make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(c),setting.PageSize,maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
		"data":data,
	})

}

//新增文章标签
func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state","0")).MustInt()
	createdBy := c.Query("created_by")

	vaild := validation.Validation{}
	vaild.Required(name, "name").Message("名称不能为空")
	vaild.MaxSize(name, 100, "name").Message("名称最长为100字符")
	vaild.Required(createdBy, "created_by").Message("创建人不能为空")
	vaild.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	vaild.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	if ! vaild.HasErrors() {
		if ! models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	} else {
		for _, err := range vaild.Errors {
			log.Fatalf("%v",err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]string),
	})
}
//修改文章标签
func EditTag(c *gin.Context) {

}
//删除文章标签
func DeleteTag(c *gin.Context) {

}