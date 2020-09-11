clean:
	rm -Rfv bin
	mkdir bin

build: clean
	go build -o bin/restic-exporter src/*.go

build-all: clean
	GOOS="linux"   GOARCH="amd64"       go build -o bin/restic-exporter__linux-amd64 src/*.go
	GOOS="linux"   GOARCH="arm" GOARM=6 go build -o bin/restic-exporter__linux-armv6 src/*.go
	GOOS="linux"   GOARCH="arm" GOARM=7 go build -o bin/restic-exporter__linux-armv7 src/*.go
	GOOS="linux"   GOARCH="arm"         go build -o bin/restic-exporter__linux-arm   src/*.go
	GOOS="darwin"  GOARCH="amd64"       go build -o bin/restic-exporter__macos-amd64 src/*.go
	GOOS="windows" GOARCH="amd64" go build -o bin/restic-exporter__win-amd64 src/*.go
