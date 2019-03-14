package common

import "sync"

type IdRequest struct {
	Biz string
}

//手动实现Map
type BizChanMap struct {
	Map map[string]chan int64
	sync.RWMutex
}
