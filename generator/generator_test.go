package generator

import (
	"fmt"
	"testing"
	"time"
)


func TestGetId(t *testing.T) {
	id,err:=GetId("test")
	if err!=nil {
		t.Error(err)
	}
	fmt.Println(id)
}

//generate with unconfig tag
func TestGetId3(t *testing.T) {
	ch:=make(chan int64,100)
	generate("unconfig",ch,time.Now().String())
	tt := time.NewTicker(time.Second)
Loop:
	for{
		select {
		case id:=<-ch:
			fmt.Println(id)
		case <-tt.C:
			fmt.Println("timeout...break loop")
			break Loop
		}
	}
	fmt.Println("end")
}


//generate with test tag
func TestGetId2(t *testing.T) {
	ch:=make(chan int64,100)
	generate("test",ch,time.Now().String())
	tt := time.NewTicker(time.Second)
	Loop:
	for{
		select {
		case id:=<-ch:
			fmt.Println(id)
		case <-tt.C:
			fmt.Println("timeout...break loop")
			break Loop
		}
	}
	fmt.Println("end")
}
