 rm -rf ./out/*

 GOOS=linux GOARCH=amd64 go build -o ./out/handler handler.go