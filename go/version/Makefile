all:build

build:
	go build -o bin/version cmd/version.go

version:
	bin/version

build_restart:
	go build -o bin/restart cmd/restart.go

restart:build_restart
	bin/restart

clean:
	rm -rf bin

.PHONY:clean build_restart build version restart all
