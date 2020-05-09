package main

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"log"
)

//连接数据库
var Db *sqlx.DB

func init(){
	db,err := sqlx.Open("mysql","root:root@tcp(127.0.0.1:3306)/LOGIN?charset=utf8")
	if err != nil{
		log.Fatal(err.Error())
	}
	if err = db.Ping() ;err != nil{
		log.Fatal(err.Error())
	}
	Db=db

}


func LOGIN(Username string)(User,error){
	mod := User{}
	err := Db.Get(&mod,"select * from user where username = ?",Username)
	if err != nil {
		log.Println(err)
	}
	return mod,err
}

func SIGNUP(inf User)error{
	//开启事务
	sp,err :=Db.Begin()
	if err != nil{
		return err
	}
	result,err := sp.Exec("insert into user(username,password)values (?,?)",inf.Username,inf.Password)
	if err != nil{
		//回滚
		sp.Rollback()
		return err
	}
	rows,_ := result.RowsAffected()
	if rows <1{
		//回滚
		sp.Rollback()
		return errors.New("rows affected < 1")
	}
	//提交
	sp.Commit()
	return nil
}

func UPDATEPSW(inf User)error{
	//开启事务
	sp,err :=Db.Begin()
	if err != nil{
		return err
	}
	result,err := sp.Exec("update user set username = ?,password = ?",inf.Username,inf.Password)
	if err != nil{
		//回滚
		sp.Rollback()
		return err
	}
	rows,_ := result.RowsAffected()
	if rows <1{
		//回滚
		sp.Rollback()
		return errors.New("rows affected < 1")
	}
	//提交
	sp.Commit()
	return nil
}