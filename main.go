package main

import (
	"fmt"
	"log"
	"syscall"

	"github.com/fvbock/endless"

	"gin-blog/pkg/setting"
	"gin-blog/routers"
)

func main() {
	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err : %v", err)
	}
}

/*
//正常启动停止的main函数，没有热更新功能
import(
	"fmt"
	"net/http"

	"github.com/koral5408/gin-blog/pkg/setting"
	"github.com/koral5408/gin-blog/routers"
)
func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:		fmt.Sprintf(":%d",setting.HTTPPort),
		Handler:	router,
		ReadTimeout:setting.ReadTimeout,
		WriteTimeout:setting.WriteTimeout,
		MaxHeaderBytes:1<<20,
	}

	s.ListenAndServe()
}
*/

/*
func main() {
	router := gin.Default()

	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":"test",
		})
	})

	s := &http.Server{
		Addr:		fmt.Sprintf(":%d",setting.HTTPPort),
		Handler:	router,
		ReadTimeout:setting.ReadTimeout,
		WriteTimeout:setting.WriteTimeout,
		MaxHeaderBytes:1<<20,
	}

	s.ListenAndServe()
}
*/
