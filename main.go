package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinyidong/IdGenerator/common"
	"github.com/jinyidong/IdGenerator/generator"
	"net/http"
	"time"
)

func main()  {

	defaultChan:=make(chan int,common.DefaultChanLength)

	generator.Generate(defaultChan)

	go func() {
		t := time.NewTicker(time.Second)
		for t:=range t.C {
			if len(defaultChan)<= cap(defaultChan)/3{
				generator.Generate(defaultChan)
				fmt.Println("generate to chan at:"+t.String())
			}
		}
	}()

	router:=gin.Default()

	router.POST("/", func(context *gin.Context) {
		context.JSON(http.StatusOK,gin.H{"Id":<-defaultChan})
	})

	router.Run(":8081")
}
