package main

import (
	_ "exam/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {

	//fmt.Println(123)
	beego.Run()
}
