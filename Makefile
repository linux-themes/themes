BINARY_NAME=themes

make:
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME} main.go

run:
	./themes flag1 flag2 flag3 

clean:
	go clean
	rm themes