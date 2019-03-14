package generator

import (
	"fmt"
	"github.com/jinyidong/IdGenerator/common"
	"github.com/jinyidong/IdGenerator/zkmanager"
	"github.com/samuel/go-zookeeper/zk"
	"strconv"
)


func Generate(ch chan<- int) {
	zkConn:= zkmanager.GetZkConn()
	lock := zk.NewLock(zkConn, common.DefaultIgPath, zk.WorldACL(zk.PermAll))
	if err := lock.Lock(); err != nil {
		fmt.Println(err)
		return
	}

	flagBytes,stat,err := zkConn.Get(common.DefaultIgPath)
	if err!=nil {
		fmt.Println(err)
		return
	}

	if len(flagBytes)==0 {
		panic(common.DefaultIgPath+"节点不存在")
	}

	flag,err:=strconv.Atoi(string(flagBytes))

	if err!=nil {
		fmt.Println(err)
		return
	}

	currentValue:=flag+common.DefaultStep

	_,err=zkConn.Set(common.DefaultIgPath,[]byte(strconv.Itoa(currentValue)),stat.Version+1)

	if err!=nil {
		fmt.Println(err)
		return
	}

	go func() {
		for i:=1;i<=common.DefaultStep;i++{
			ch<-flag+i
		}
	}()

	if err := lock.Unlock(); err != nil {
		fmt.Println(err)
		return
	}
}
