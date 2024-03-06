BINARY_NAME=themes

make:
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME} main.go

run:
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME} main.go
	./themes list  
# ./themes help  
# ./themes install package_name flag3 
	

clean:
	go clean
	rm themes