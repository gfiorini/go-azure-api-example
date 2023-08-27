del out /q

set GOOS=windows
set GOARCH=amd64

go build -o ./out/handler handler.go
