package main

import (
	"log"
	"net/http"
	"net/rpc"
)

//LOGIN
type User struct {
	Username string
	Password string
}
type Rect struct{}


//RPC方法

//LOGIN
func (r *Rect)Login(inf User,res *int)error{
	result,err :=LOGIN(inf.Username)
	if result.Password == inf.Password {
		*res = 200
	}else {
		*res=300
	}
	return err
}

//SIGNUP
func (r *Rect)Signup(inf User,res *int)error{
	err :=SIGNUP(inf)
	if err != nil {
		*res=300
	}
	*res = 200
	return nil
}

//修改密码
func (r *Rect)Updatepsw(inf User,res *int)error{
	err :=UPDATEPSW(inf)
	if err != nil {
		*res=300
	}
	*res =200
	return nil
}

func main ()  {
	//1.注册服务
	rect :=new(Rect)
	//2.注册一个rect的服务
	rpc.Register(rect)
	//3.服务处理绑定到http协议上
	rpc.HandleHTTP()
	//4.监听服务
	err :=http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Println(err)
	}

}
