package util

import (
	"github.com/gin-gonic/gin"
	"gin-blog/pkg/setting"
	"github.com/unknwon/com"
)

// 返回上一页最后一个元素的序号
func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}

	return result
}
