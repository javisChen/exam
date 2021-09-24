package redis

import (
	"fmt"
	"testing"
	"time"
)

var key = "test:key2"

func TestSet(t *testing.T) {
	err := Set(key, "hello redis")
	if err != nil {
		t.Fail()
	}
}

func TestSetWithExpr(t *testing.T) {
	err := SetWithExpire(key, "hello redis", 15, time.Second)
	if err != nil {
		t.Fail()
	}
}

func TestGet(t *testing.T) {
	result, err := Get(key)
	if err != nil {
		t.Fail()
	}
	fmt.Println(result)
}
