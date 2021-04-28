package main

import (
	"flag"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
	"runtime"
	"log"
)

var host string //ip
var port int 	//端口
var conn redis.Conn 

func init() {
	flag.Usage = usage
	flag.StringVar(&host, "host", "127.0.0.1", "set the redis connect host")
	flag.IntVar(&port, "port", 6379, "set the redis connect port")
}

//命令提示帮助信息
func usage() {
	fmt.Println("Usage: go run demo.go [-host] [-port]")
	flag.PrintDefaults()
}

//获取调用函数文件名及行号加函数名
func getFuncInfo() string {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return ""
	}

	funcName := runtime.FuncForPC(pc).Name()
	return fmt.Sprintf("%s:%d:%s", file, line, funcName)
}

//统一错误校验及输出
func checkErr(err error, msg string) {
	if err != nil {
		fmt.Println(getFuncInfo(), msg, err)
		os.Exit(1)
	}
}

//打印输出结果
func show(arg interface{}, msg string) {
	fmt.Printf("%s -- %s: %#v\n", getFuncInfo(), msg, arg)
}

func main() {
	flag.Parse()
	addr := fmt.Sprintf("%s:%v", host, port)
	fmt.Println("going to dial",addr)
	/*
	var err error
	conn, err = redis.Dial("tcp", addr)
	checkErr(err, "redis.Dial")
	//*/
	
	c, err := redis.Dial("tcp", addr)
	checkErr(err, "redis.Dial")
	logger := log.Default()
	conn = redis.NewLoggingConn(c, logger, "demo")
	
	defer conn.Close()

	_string()
	_bool()
}

//输出结果为字符串类型
func _string() {
	conn.Do("set", "hello", "world")
	s, err := redis.String(conn.Do("get", "hello"))
	checkErr(err, "redis.String")
	show(s, "get hello")
}

//输出结果为bool 类型
func _bool() {
	conn.Do("set", "foo", 1)
	exists, err := redis.Bool(conn.Do("exists", "foo"))
	checkErr(err, "redis.Bool")
	show(exists, "exists foo")
}
