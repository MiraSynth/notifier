build-all:
	GOOS=windows GOARCH=amd64 go build -o ./build/windows/notifier.exe
	GOOS=linux GOARCH=amd64 go build -o ./build/linux/notifier
	GOOS=darwin GOARCH=arm64 go build -o ./build/darwin/notifier