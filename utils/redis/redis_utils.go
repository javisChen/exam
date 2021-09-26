package redis

import (
	"context"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

var ctx = context.Background()
var client *redis.Client

func init() {
	host := beego.AppConfig.DefaultString("redis.host", "localhost")
	port := beego.AppConfig.DefaultInt("redis.port", 6379)
	password := beego.AppConfig.DefaultString("redis.password", "")
	db := beego.AppConfig.DefaultInt("redis.db", 0)
	client = redis.NewClient(&redis.Options{
		Addr:        host + ":" + strconv.Itoa(port),
		Password:    password, // no password set
		DB:          db,       // use default DB
		DialTimeout: 10 * time.Second,
		PoolTimeout: 2 * time.Second,
	})
}

func Set(key string, value interface{}) error {
	return SetWithExpire(key, value, 0, 0)
}

// 过期时间单位默认为秒
func SetWithExpire(key string, value interface{}, ttl int64, duration time.Duration) error {
	fmt.Println("start", time.Now())
	err := client.Set(ctx, key, value, time.Duration(ttl)*duration).Err()
	fmt.Println("end", time.Now())
	return err
}

func Get(key string) (string, error) {
	return client.Get(ctx, key).Result()
}

func Remove(key string) error {
	return client.Del(ctx, key).Err()
}
