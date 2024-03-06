BINARY_NAME=themes

make:
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME} main.go utils.go

run:
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME} main.go utils.go
	./themes
	./themes help
	./themes list  
	./themes list icons 
	./themes list themes 
	./themes list create 
	./themes list build 
	./themes list install 
	./themes list install icons https://github.com/sudo-adduser-jordan/mint-y-winx/raw/main/mint-y-winx.tar.xz
	./themes list install themes https://github.com/sudo-adduser-jordan/mint-y-winx/raw/main/mint-y-winx.tar.xz
	# ,/themes set icons package_1
	# ,/themes set themes package_1
	./themes remove icons package_1
	./themes remove themes package_1


clean:
	go clean
	rm themes