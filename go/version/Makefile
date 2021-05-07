all: build_restart

build_restart:
	go build -o bin/restart cmd/restart.go

restart:build_restart
	bin/restart

clean:
	rm -f bin/version

.PHONY:clean build_restart restart all
