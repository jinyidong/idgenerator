package generator

import (
	"fmt"
	"github.com/jinyidong/IdGenerator/common"
	"github.com/jinyidong/IdGenerator/zkmanager"
	"github.com/samuel/go-zookeeper/zk"
	"strconv"
	"time"
)

var bizChanMap common.BizChanMap

func init()  {
	bizChanMap=common.BizChanMap{Map:make(map[string]chan int64)}
	zkConn:= zkmanager.GetZkConn()
	children,_,err:=zkConn.Children(common.IgRootPath)
	if err!=nil {
		fmt.Println("get idgenerator children failed")
	}

	for _,child:=range children{
		fmt.Println(child)
		bizChanMap.Map[child]=make(chan int64,common.DefaultChanCap)
	}

	go func() {//单独起协程，向相关channel生产数据，是否需要加锁
		t := time.NewTicker(5*time.Second)
		for t:=range t.C {
			for biz,ch:=range bizChanMap.Map{
				go generate(biz,ch,t.String())
			}
		}
	}()
}

func generate(biz string,ch chan<- int64,timeNow string)  {
	if len(ch)> cap(ch)/3{
		return
	}
	bizPath:=common.IgRootPath+"/"+biz
	zkConn:= zkmanager.GetZkConn()
	lock := zk.NewLock(zkConn, bizPath, zk.WorldACL(zk.PermAll))
	if err := lock.Lock(); err != nil {
		fmt.Println(err)
		return
	}

	flagBytes,_,err := zkConn.Get(bizPath)
	if err!=nil {
		fmt.Println(err)
		return
	}

	if len(flagBytes)==0 {
		fmt.Println("zk路径不存在"+bizPath)
		return
	}

	flag, err := strconv.ParseInt(string(flagBytes), 10, 64)
	if err!=nil {
		fmt.Println(err)
		return
	}

	currentValue:=flag+common.DefaultStep

	_,err=zkConn.Set(bizPath,[]byte(strconv.FormatInt(currentValue,10)),-1)

	if err!=nil {
		fmt.Println(err)
		return
	}

	go func() {
		for i:=1;i<=common.DefaultStep;i++{
			ch<-flag+int64(i)
		}
	}()

	if err := lock.Unlock(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("generate to "+biz+" chan at:"+timeNow)
}

func GetId(biz string) (int64,error) {
	bizChanMap.RLock()
	ch:=bizChanMap.Map[biz]
	bizChanMap.RUnlock()
	if nil!=ch {
		return <-ch,nil
	}

	//Step1：判断节点是否存在
	bizPath:=common.IgRootPath+"/"+biz
	zkConn:= zkmanager.GetZkConn()
	flagBytes,_,err := zkConn.Get(bizPath)
	if err!=nil {
		fmt.Println(err)
		return 0,err
	}

	if len(flagBytes)==0 {
		fmt.Println("zk路径不存在"+bizPath)
		return 0,fmt.Errorf("zk路径不存在"+bizPath)
	}

	bizChanMap.Lock()
	ch=make(chan int64,common.DefaultChanCap)
	bizChanMap.Map[biz]=ch
	go generate(biz,ch,time.Now().String())
	bizChanMap.Unlock()

	return <-ch,nil
}
