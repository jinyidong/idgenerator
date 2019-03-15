package zkmanager

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"strings"
	"sync"
	"time"
)

var once sync.Once

var zkConn *zk.Conn

func GetZkConn() *zk.Conn {
	once.Do(func() {
		servers := strings.Split(BasicConfig.Zookeeper, ",")
		var wg = &sync.WaitGroup{}
		wg.Add(1) //Add用来设置等待线程数量
		go func() {
			var ech <-chan zk.Event
			var err error
			var c *zk.Conn
			if zkConn == nil {
				c, ech, err = zk.Connect(servers, time.Second) //*10
				zkConn = c
			}
			if err != nil {
				panic(err)
			}
			go func() {
				for {
					select {
					case ch := <-ech:
						{
							switch ch.State {
							case zk.StateConnecting:
								{
									fmt.Println("StateConnecting")
								}
							case zk.StateConnected: //若链接后出现断连现象 重连时会报异常 因为此时同步信号量已为0
								{
									wg.Done()
									fmt.Println("StateConnected")
								}
							case zk.StateExpired:
								{
									fmt.Println("StateExpired")
								}
							case zk.StateDisconnected:
								{
									wg.Add(1)
									fmt.Println("StateDisconnected")
								}
							}
						}
					}
				}
			}()
		}()
		wg.Wait()
	})
	return zkConn
}

