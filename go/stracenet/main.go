package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
)

//监听端口号
var port int

func init() {
	flag.IntVar(&port, "port", 8001, "set the port")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

//检验错误
//支持配置是否程序退出
func checkErr(err error, msg string, exit bool) {
	if err != nil {
		log.Println(msg, err)
		if exit {
			os.Exit(1)
		}
	}
}

func main() {
	flag.Parse()
	addr := fmt.Sprintf(":%d", port)
	log.Println("going to listen", addr)
	log.Println("system info cpu", runtime.NumCPU(), "goroutine", runtime.GOMAXPROCS(-1))
	//1、开始监听
	l, err := net.Listen("tcp", addr)
	checkErr(err, "Listen", true)
	defer l.Close()
	log.Println("pid", os.Getpid())

	for {
		//2、获取连接
		conn, err := l.Accept()
		if err != nil {
			log.Println("Accept", err)
			continue
		}
		checkErr(err, "Accept", false)
		log.Println("RemoteAddr", conn.RemoteAddr().String())
		//3、处理连接
		go handleConn(conn)
	}
}

//针对连接读写操作
func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		rb := make([]byte, 1024)
		//4、读取数据
		n, err := conn.Read(rb)
		if n == 0 {
			log.Println("client is closed")
			return
		}
		if err != nil {
			log.Println("Read", err)
			return
		}
		log.Println("Read", n, string(rb[0:n-1]))

		//5、回写数据
		_, err = conn.Write(append([]byte("Reply:"), rb[0:n]...))
		checkErr(err, "Write", false)
	}
}
