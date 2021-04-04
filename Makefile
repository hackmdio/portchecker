
all: prebuild buildLinux386 buildLinuxAmd64 buildLinuxArm64 buildDarwin386 buildDarwinAmd64 buildWindows386 buildWindowsAmd64

prebuild:
	mkdir -p dist

buildLinux386:
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o dist/portchecker-linux-386 ./portchecker.go

buildLinuxAmd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/portchecker-linux-amd64 ./portchecker.go

buildLinuxArm64:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o dist/portchecker-linux-arm64 ./portchecker.go

buildDarwin386:
	CGO_ENABLED=0 GOOS=darwin GOARCH=386 go build -o dist/portchecker-darwin-386 ./portchecker.go

buildDarwinAmd64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/portchecker-darwin-amd64 ./portchecker.go

buildWindows386:
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o dist/portchecker-windows-386.exe ./portchecker.go

buildWindowsAmd64:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/portchecker-windows-amd64.exe ./portchecker.go
