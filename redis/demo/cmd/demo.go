package main

import (
	"flag"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
	"runtime"
)

var port int
var conn redis.Conn

func init() {
	flag.Usage = usage
	flag.IntVar(&port, "port", 6379, "set the redis connect port")
}

func usage() {
	fmt.Println("Usage: go run demo.go [-port]")
	flag.PrintDefaults()
}

func getFuncInfo() string {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return ""
	}

	funcName := runtime.FuncForPC(pc).Name()
	return fmt.Sprintf("%s:%d:%s", file, line, funcName)
}

func checkErr(err error, msg string) {
	if err != nil {
		fmt.Println(getFuncInfo(), msg, err)
		os.Exit(1)
	}
}

func show(arg interface{}, msg string) {
	fmt.Printf("%s -- %s: %#v\n", getFuncInfo(), msg, arg)
}

func main() {
	flag.Parse()
	addr := fmt.Sprintf(":%v", port)
	var err error

	conn, err = redis.Dial("tcp", addr)
	checkErr(err, "redis.Dial")
	defer conn.Close()

	_string()
	_bool()
}

func _string() {
	conn.Do("set", "hello", "world")
	s, err := redis.String(conn.Do("get", "hello"))
	checkErr(err, "redis.String")
	show(s, "get hello")
}

func _bool() {
	conn.Do("set", "foo", 1)
	exists, err := redis.Bool(conn.Do("exists", "foo"))
	checkErr(err, "redis.Bool")
	show(exists, "exists foo")
}
