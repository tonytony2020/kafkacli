all:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o kafkacli.bin main.go
