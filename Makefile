BINARY_NAME=themes

make:
	GOARCH=amd64 GOOS=darwin go build -o bin/${BINARY_NAME}-darwin.exe main.go
	GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux.exe main.go
	GOARCH=amd64 GOOS=windows go build -o bin/${BINARY_NAME}-windows.exe main.go

clean:
	go clean
	rm -rf bin