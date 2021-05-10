package core

import (
	"io"
	"os"
	"strconv"
	"syscall"
	"version/logger"
)

func updateProc() {
	pidFile := GetConfPidFile()
	//读取上次进程 pid，并尝试发送关闭信号
	if fd, err := os.Open(pidFile); err == nil {
		pidBytes, _ := io.ReadAll(fd)
		pid, _ := strconv.Atoi(string(pidBytes))
		if pid > 0 {
			logger.Info("going to close old server, pid=", pid)
			signals := []syscall.Signal{syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL}
			if proc, err := os.FindProcess(pid); err == nil {
				for _, sig := range signals {
					if err = proc.Signal(sig); err != nil {
						logger.Error("proc.Signal err", sig, err)
					}
					var stat *os.ProcessState
					stat, err = proc.Wait()
					if err != nil || stat.Exited() {
						break
					}
				}
			} else {
				logger.Error("os.FindProcess err= ", err)
			}
		}
		fd.Close()
	} else {
		logger.Info("open pidFile failed ", err)
	}

	//保存当前 pid
	if fd, err := os.Create(pidFile); err == nil {
		pid := os.Getpid()
		fd.Write([]byte(strconv.Itoa(pid)))
		fd.Close()
	}

}
