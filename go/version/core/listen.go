package core

import (
	"context"
	"golang.org/x/sys/unix"
	"net"
	"syscall"
	"version/logger"
)

func Listen(ctx context.Context, network, address string) (net.Listener, error) {
	lisCfg := &net.ListenConfig{
		Control: func(network, address string, c syscall.RawConn) error {
			var err error
			err1 := c.Control(func(fd uintptr) {
				err = syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, unix.SO_REUSEPORT, 1)
				if err != nil {
					logger.Error("syscall.SetsockoptInt failed ", err)
				}
			})
			if err1 != nil {
				logger.Error("syscall.RawConn.Control failed ", err1)
				err = err1
			}

			return err
		},
	}

	return lisCfg.Listen(ctx, network, address)

}
