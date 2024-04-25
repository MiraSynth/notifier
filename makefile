build-all: build-windows build-linux build-darwin

build-windows:
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ./build/windows/notifier.exe

build-linux:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./build/linux/notifier

build-darwin:
	GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o ./build/darwin/notifier

build-docker:
	sudo docker build -t "miras-notifier:alpha" .

test:
	curl localhost:3038/api/v1/notify -X POST -H "Content-Type: application/json" -d '{"context":"Notifier App", "content":"This is a test, do not **sit** down"}'