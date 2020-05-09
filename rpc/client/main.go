package main

import (
	"fmt"
	"log"
	"net/rpc"
)

//传的参数
type Params struct {
	Username string
	Password string
}

func main(){
	//1.连接rpc服务
	conn,err :=rpc.DialHTTP("tcp",":8080")
	if err != nil {
		log.Fatal(err)
	}
	//2.调用方法
	var code1,code2,code3 int
	//注册
	err1:=conn.Call("Rect.Login",Params{"abc","123"},&code2)
	if err1 != nil {
		log.Fatal(err1)
	}
	//登录
	err2:=conn.Call("Rect.Login",Params{"123","123456"},&code2)
	if err2 != nil {
		log.Fatal(err2)
	}
	//修改密码
	err3:=conn.Call("Rect.Updatepsw",Params{"123","1234567"},&code3)
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Println(code1,code2,code3)

}

