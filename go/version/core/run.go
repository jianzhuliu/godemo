//入口文件
package core

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"version/logger"
)

//全局配置
var (
	server     *http.Server
	ctx        context.Context
	cancelFunc context.CancelFunc
)

//启动入口
func Run() {
	logger.Info("current pid is", os.Getpid())

	//异常退出信号
	onExit := make(chan error, 1)

	go func() {
		//业务逻辑
		if err := run(); err != nil {
			onExit <- err
		}
		close(onExit)
	}()

	//监听系统信号
	onSignal := make(chan os.Signal)
	signal.Notify(onSignal, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-onSignal:
		logger.Info("exit by signal ", sig)
	case err := <-onExit:
		logger.Info("exit by err ", err)
	}

	//收尾清理工作
	clean()
}

//收尾清理，释放连接
func clean() {
	logger.Info("going to shutdown")
	server.Shutdown(context.Background())
}

//真实业务逻辑
func run() error {
	ctx, cancelFunc = context.WithCancel(context.Background())
	addr := GetConfAddr()
	listener, err := Listen(ctx, "tcp", addr)
	if err != nil {
		logger.Exit("create listener fail", err)
	}
	server = &http.Server{
		Addr: addr,
	}

	logger.Info("server is going to listen at ", addr)
	go updateProc()

	return server.Serve(listener)
}
