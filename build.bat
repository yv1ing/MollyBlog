SET TARGET=dist
SET CGO_ENABLED=0

SET EXECUTE_NAME=molly_blog

SET GOOS=linux
SET GOARCH=amd64
go build -v -a -ldflags "-s -w" -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -o %TARGET%/%EXECUTE_NAME%_linux_amd64


SET GOOS=windows
SET GOARCH=amd64
go build -v -a -ldflags "-s -w" -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -o %TARGET%/%EXECUTE_NAME%_windows_x64.exe
