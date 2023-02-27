#64bit
#GOOS=windows GOARCH=amd64 go build -o bin/wechatbot-amd64.exe main.go

#32-bit
#GOOS=windows GOARCH=386 go build -o bin/wechatbot-386.exe main.go

# 64-bit
#GOOS=linux GOARCH=amd64 go build -o bin/wechatbot-amd64-linux main.go

# 32-bit
#GOOS=linux GOARCH=386 go build -o bin/wechatbot-386-linux main.go

# 64-bit
#GOOS=darwin GOARCH=amd64 go build -o bin/wechatbot-amd64-darwin main.go

# 32-bit
#GOOS=darwin GOARCH=386 go build -o bin/wechatbot-386-darwin main.go

# telegram
GOOS=linux GOARCH=amd64 go build -o bin/telegram-amd64-linux telegram.go

# poetry
GOOS=linux GOARCH=amd64 go build -o bin/peotry-amd64-linux peotry.go