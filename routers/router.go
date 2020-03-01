package routers

import (
	"gin-blog/pkg/upload"
	"github.com/gin-gonic/gin"
	"gin-blog/middleware/jwt"
	"gin-blog/routers/api"
	v1 "gin-blog/routers/api/v1"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"

	"gin-blog/pkg/setting"

	_ "gin-blog/docs"
)

func InitRouter() *gin.Engine{
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	//设置静态文件路由
	r.StaticFS("/upload/images",http.Dir(upload.GetImageFullPath()))

	//swagger 自动生成api接口文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//获取token
	r.GET("/auth", api.GetAuth)

	//上传图片
	r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("api/v1")
	apiv1.Use(jwt.JWT())
	{
		/***************标签相关接口*****************/
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//修改标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		/***************文章相关接口**************/
		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取制定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
		/**************************/

	}

	/*
	r.GET("/test",func(c *gin.Context){
		c.JSON(200,gin.H{
			"message":	"test",
		})
	})
	*/
	return r
}