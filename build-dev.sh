go build -o slrun-installer-dev -ldflags="-s -w" src/*.go
upx --brute slrun-installer-dev
