# Linux

GOOS=linux GOARCH=amd64 go build -o dist/slrun-installer-linux-x64 -ldflags="-s -w" src/*.go
upx --brute dist/slrun-installer-linux-x64

GOOS=linux GOARCH=386 go build -o dist/slrun-installer-linux-x86 -ldflags="-s -w" src/*.go
upx --brute dist/slrun-installer-linux-x86

# MacOS

GOOS=darwin GOARCH=amd64 go build -o dist/slrun-installer-macos-x64 -ldflags="-s -w" src/*.go
upx --brute dist/slrun-installer-macos-x64

GOOS=darwin GOARCH=386 go build -o dist/slrun-installer-macos-x86 -ldflags="-s -w" src/*.go
upx --brute dist/slrun-installer-macos-x86

# FreeBSD

#GOOS=freebsd GOARCH=amd64 go build -o dist/slrun-installer-freebsd-x64 -ldflags="-s -w" src/*.go
#upx --brute dist/slrun-installer-freebsd-x64

#GOOS=freebsd GOARCH=386 go build -o dist/slrun-installer-freebsd-x86 -ldflags="-s -w" src/*.go
#upx --brute dist/slrun-installer-freebsd-x86

# Windows

GOOS=windows GOARCH=amd64 go build -o dist/slrun-installer-win-x64.exe -ldflags="-s -w" src/*.go
upx --brute dist/slrun-installer-win-x64.exe

GOOS=windows GOARCH=386 go build -o dist/slrun-installer-win-x86.exe -ldflags="-s -w" src/*.go
upx --brute dist/slrun-installer-win-x86.exe
