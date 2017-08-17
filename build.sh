# x64
GOOS=linux GOARCH=amd64 go build -o dist/slrun-installer-linux-x64 -ldflags="-w" src/*.go
GOOS=darwin GOARCH=amd64 go build -o dist/slrun-installer-mac-x64 -ldflags="-w" src/*.go
GOOS=windows GOARCH=amd64 go build -o dist/slrun-installer-win-x64.exe -ldflags="-w" src/*.go
# x86
GOOS=linux GOARCH=386 go build -o dist/slrun-installer-linux-x86 -ldflags="-w" src/*.go
GOOS=darwin GOARCH=386 go build -o dist/slrun-installer-mac-x86 -ldflags="-w" src/*.go
GOOS=windows GOARCH=386 go build -o dist/slrun-installer-win-x86.exe -ldflags="-w" src/*.go