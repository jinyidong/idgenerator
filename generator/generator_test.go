package generator

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	GetId("test")
}

func TestGetId(t *testing.T) {
	id,err:=GetId("test")
	if err!=nil {
		t.Error(err)
	}
	fmt.Println(id)
}
