package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinyidong/IdGenerator/common"
	"github.com/jinyidong/IdGenerator/generator"
	"net/http"
)

func main()  {

	router:=gin.Default()

	router.POST("/", func(context *gin.Context) {
		var req common.IdRequest
		err:=context.BindJSON(&req)
		if err!=nil {
			context.JSON(http.StatusOK,gin.H{"ErrorCode":-1,"Id":0})
		}else{
			id,err:=generator.GetId(req.Biz)
			if err!=nil {
				context.JSON(http.StatusOK,gin.H{"ErrorCode":-1,"Id":0})
			}else {
				context.JSON(http.StatusOK,gin.H{"ErrorCode":0,"Id":id})
			}
		}
	})

	router.Run(":8080")
}
