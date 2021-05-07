/*
程序优雅退出及平滑版本更新
*/
package main

import (
	"context"
	"flag"
	"fmt"
	"golang.org/x/sys/unix"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

var (
	port       int
	listener   net.Listener
	ctx        context.Context
	cancelFunc context.CancelFunc
	pidFile    string
)

func init() {
	flag.IntVar(&port, "port", 8001, "set the listen port")
	log.SetFlags(log.LstdFlags | log.Lshortfile)

}

//创建监听
func listen(ctx context.Context, network string, addr string) (net.Listener, error) {
	lisCfg := &net.ListenConfig{
		Control: func(network string, address string, c syscall.RawConn) error {
			var err error
			err1 := c.Control(func(fd uintptr) {
				err = syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, unix.SO_REUSEPORT, 1)
				if err != nil {
					log.Fatalln("set socket option failed ", err)
				}
			})
			if err1 != nil {
				log.Fatalln("control listener failed ", err1)
				err = err1
			}
			return err
		},
	}

	return lisCfg.Listen(ctx, network, addr)
}

//处理老程序退出及 pid 记录
func updateProc() {
	//如果pid 文件存在，读取pid
	if fd, err := os.Open(pidFile); err == nil {
		pidBytes, _ := io.ReadAll(fd)
		pid, _ := strconv.Atoi(string(pidBytes))
		if pid > 0 {
			log.Println("going to close last pid=", pid)
			//发送关闭信号
			signals := []syscall.Signal{syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL}
			if proc, err := os.FindProcess(pid); err == nil {
				for _, sig := range signals {
					if err = proc.Signal(sig); err != nil {
						log.Println("proc.Signal,", sig, err)
						break
					}
					var stat *os.ProcessState
					stat, err = proc.Wait()
					if err != nil || stat.Exited() {
						break
					}
				}
			} else {
				log.Println("os.FindProcess ", err)
			}

		}
		fd.Close()
	} else {
		log.Println("os.Open pidFile ", err)
	}

	//保存当前 pid
	if fd, err := os.Create(pidFile); err == nil {
		pid := os.Getpid()
		fd.Write([]byte(strconv.Itoa(pid)))
		fd.Close()
	} else {
		log.Println("os.Create pidFile err ", err)
	}
}

//业务逻辑入口
func run() error {
	//监听端口
	addr := fmt.Sprintf(":%d", port)
	log.Println("going to listen ", addr)
	log.Println("current pid=", os.Getpid())

	//上下文及取消函数
	ctx, cancelFunc = context.WithCancel(context.Background())

	var err error
	listener, err = listen(ctx, "tcp", addr)
	if err != nil {
		log.Fatalln("fail to listen tcp ", addr)
		return err
	}

	//关闭老程序
	go updateProc()

	for {
		//需要判断是否已经退出
		select {
		case <-ctx.Done():
			log.Println("listener isClosed")
			return nil
		default:
		}

		log.Println("going to accept,pid=", os.Getpid())
		conn, err := listener.Accept()
		if err == nil {
			go handleConn(conn)
		}

	}

}

//处理连接
func handleConn(conn net.Conn) {
	defer conn.Close()

	remoteAddr := conn.RemoteAddr().String()
	prefix := fmt.Sprintf("client[%s]", remoteAddr)
	log.Printf("%s is connected \n", prefix)

	for {
		select {
		case <-ctx.Done():
			log.Println("going to close conn, addr=", remoteAddr)
			return
		default:
		}

		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if n == 0 {
			log.Printf("%s is closed \n", prefix)
			return
		}
		if err != nil {
			log.Printf("%s read err=%v \n", prefix, err)
			return

		}
		log.Printf("%s read n=%d, content=%s ", prefix, n, string(buf[0:n-1]))

		conn.Write(append([]byte("Reply:"), buf[0:n]...))

	}
}

//程序退出清理工作
func stop() {
	log.Println("going to close")
	//取消上下文，关闭历史连接
	cancelFunc()

	//关闭监听
	listener.Close()
}

func main() {
	//命令行参数解析
	flag.Parse()

	//定义 pid 存储文件名
	pidFile = fmt.Sprintf("pid.%d", port)
	log.Println("pidFile ", pidFile)

	//异常退出信号
	onExit := make(chan error, 1)

	//执行业务逻辑
	go func() {
		if err := run(); err != nil {
			onExit <- err
		}
		close(onExit)
	}()

	//监听系统退出信号
	onSignal := make(chan os.Signal)
	signal.Notify(onSignal, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-onSignal:
		log.Println("exit by signal ", sig)
	case err := <-onExit:
		log.Println("exit by err ", err)
	}

	//退出收尾工作
	stop()

}
