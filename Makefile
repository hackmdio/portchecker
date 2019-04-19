
all: prebuild buildLinux386 buildLinuxAmd64 buildDarwin386 buildDarwinAmd64 buildWindows386 buildWindowsAmd64

prebuild:
	mkdir -p dist

buildLinux386:
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o dist/portchecker-linux-386 ./cmd/portchecker.go

buildLinuxAmd64:
	mkdir -p dist/linux/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/portchecker-linux-amd64 ./cmd/portchecker.go

buildDarwin386:
	mkdir -p dist/darwin/386
	CGO_ENABLED=0 GOOS=darwin GOARCH=386 go build -o dist/portchecker-darwin-386 ./cmd/portchecker.go

buildDarwinAmd64:
	mkdir -p dist/darwin/amd64
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/portchecker-darwin-amd64 ./cmd/portchecker.go

buildWindows386:
	mkdir -p dist/windows/386
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o dist/portchecker-windows-386.exe ./cmd/portchecker.go

buildWindowsAmd64:
	mkdir -p dist/windows/amd64
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/portchecker-windows-amd64.exe ./cmd/portchecker.go
