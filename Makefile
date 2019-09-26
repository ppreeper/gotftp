all: clean build

build:
	go build gotftp.go

test:
	go test -cover -race -v

win32: clean
	env GOOS=windows GOARCH=386 go build -a -o gotftp32.exe gotftp.go

win64: clean
	env GOOS=windows GOARCH=amd64 go build -a -o gotftp.exe gotftp.go

clean:
	go clean -i

